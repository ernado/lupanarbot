package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/app"
	"github.com/go-faster/sdk/zctx"
	"github.com/gotd/contrib/middleware/floodwait"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/tg"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	entdb "github.com/ernado/lupanarbot/internal/db"
	"github.com/ernado/lupanarbot/internal/ent"
	"github.com/ernado/lupanarbot/internal/minust"
)

type Application struct {
	db     *ent.Client
	api    *tg.Client
	client *telegram.Client

	waiter *floodwait.Waiter
	trace  trace.Tracer
}

func (a *Application) Run(ctx context.Context) error {
	return a.waiter.Run(ctx, func(ctx context.Context) error {
		return a.client.Run(ctx, func(ctx context.Context) error {
			lg := zctx.From(ctx)
			if self, err := a.client.Self(ctx); err != nil || !self.Bot {
				if _, err := a.client.Auth().Bot(ctx, os.Getenv("BOT_TOKEN")); err != nil {
					return errors.Wrap(err, "auth bot")
				}
			} else {
				lg.Info("Already authenticated")
			}
			if self, err := a.client.Self(ctx); err == nil {
				lg.Info("Bot info",
					zap.Int64("id", self.ID),
					zap.String("username", self.Username),
					zap.String("first_name", self.FirstName),
					zap.String("last_name", self.LastName),
				)
			}
			if _, err := a.api.BotsSetBotCommands(ctx, &tg.BotsSetBotCommandsRequest{
				Scope:    &tg.BotCommandScopeDefault{},
				LangCode: "en",
				Commands: []tg.BotCommand{
					{
						Command:     "start",
						Description: "Start bot",
					},
					{
						Command:     "extremism",
						Description: "Какой экстремизм ты сегодня",
					},
				},
			}); err != nil {
				return errors.Wrap(err, "set commands")
			}
			<-ctx.Done()
			return ctx.Err()
		})
	})
}

func (a *Application) addChannel(ctx context.Context, channel *tg.Channel) error {
	return a.db.TelegramChannel.Create().
		SetID(channel.ID).
		SetAccessHash(channel.AccessHash).
		SetTitle(channel.Title).
		SetActive(true).
		Exec(ctx)
}

func (a *Application) removeChannel(ctx context.Context, channel *tg.Channel) error {
	if err := a.db.TelegramChannel.UpdateOneID(channel.ID).
		SetActive(false).
		Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return errors.Wrap(err, "update channel")
	}
	return nil
}

func (a *Application) onChannelParticipant(ctx context.Context, e tg.Entities, update *tg.UpdateChannelParticipant) error {
	switch update.NewParticipant.(type) {
	case *tg.ChannelParticipantBanned:
		// Bot was removed from channel.
		for _, c := range e.Channels {
			return a.removeChannel(ctx, c)
		}
	case *tg.ChannelParticipantAdmin:
		// Bot was added to channel.
		for _, c := range e.Channels {
			return a.addChannel(ctx, c)
		}
	default:
		if update.NewParticipant == nil {
			// Removed from channel.
			for _, c := range e.Channels {
				return a.removeChannel(ctx, c)
			}
		}
	}
	return nil
}

func extractUserID(m *tg.Message) (int64, bool) {
	if peerUser, ok := m.FromID.(*tg.PeerUser); ok {
		return peerUser.UserID, true
	}
	if peerUser, ok := m.PeerID.(*tg.PeerUser); ok {
		return peerUser.UserID, true
	}
	return 0, false
}

func (a *Application) onNewMessage(ctx context.Context, e tg.Entities, u *tg.UpdateNewMessage) error {
	ctx, span := a.trace.Start(ctx, "OnNewMessage")
	defer span.End()
	m, ok := u.Message.(*tg.Message)
	if !ok || m.Out {
		return nil
	}
	var (
		sender = message.NewSender(a.api)
		reply  = sender.Reply(e, u)
		lg     = zctx.From(ctx).With(zap.Int("msg.id", m.ID))
	)
	userID, ok := extractUserID(m)
	if !ok {
		if _, err := reply.Text(ctx, "Invalid"); err != nil {
			return err
		}
		return nil
	}
	user := e.Users[userID]
	if user == nil {
		return nil
	}
	lg.Info("New message",
		zap.String("text", m.Message),
		zap.String("user", user.Username),
		zap.String("first_name", user.FirstName),
		zap.String("last_name", user.LastName),
		zap.Int64("user_id", user.ID),
	)
	switch m.Message {
	case "/start", "/start@lupanar_chatbot":
		if _, err := reply.Text(ctx, "Hello, "+user.FirstName+"!"); err != nil {
			return errors.Wrap(err, "send message")
		}
	case "/extremism", "/extremism@lupanar_chatbot":
		elem := minust.Random()
		text := fmt.Sprintf("%d. %s", elem.ID, elem.Title)
		if _, err := reply.Text(ctx, text); err != nil {
			return errors.Wrap(err, "send message")
		}
	}
	return nil
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger, t *app.Telemetry) error {
		db, err := entdb.Open(ctx, os.Getenv("DATABASE_URL"), t)
		if err != nil {
			return errors.Wrap(err, "open database")
		}
		botToken := os.Getenv("BOT_TOKEN")
		if botToken == "" {
			return errors.New("BOT_TOKEN is empty")
		}
		appID, err := strconv.Atoi(os.Getenv("APP_ID"))
		if err != nil {
			return errors.Wrap(err, "parse APP_ID")
		}
		appHash := os.Getenv("APP_HASH")
		if appHash == "" {
			return errors.New("APP_HASH is empty")
		}
		waiter := floodwait.NewWaiter()
		dispatcher := tg.NewUpdateDispatcher()
		client := telegram.NewClient(appID, appHash, telegram.Options{
			Logger:         zctx.From(ctx).Named("tg"),
			TracerProvider: t.TracerProvider(),
			SessionStorage: entdb.NewSessionStorage(-1, db),
			UpdateHandler:  dispatcher,
			Middlewares: []telegram.Middleware{
				waiter,
			},
		})
		a := &Application{
			db:     db,
			api:    tg.NewClient(client),
			client: client,
			waiter: waiter,
			trace:  t.TracerProvider().Tracer("lupanar.bot"),
		}
		dispatcher.OnChannelParticipant(a.onChannelParticipant)
		dispatcher.OnNewMessage(a.onNewMessage)
		return a.Run(ctx)
	})
}
