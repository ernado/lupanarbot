// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ernado/lupanarbot/internal/ent/predicate"
	"github.com/ernado/lupanarbot/internal/ent/telegramchannel"
)

// TelegramChannelUpdate is the builder for updating TelegramChannel entities.
type TelegramChannelUpdate struct {
	config
	hooks    []Hook
	mutation *TelegramChannelMutation
}

// Where appends a list predicates to the TelegramChannelUpdate builder.
func (tcu *TelegramChannelUpdate) Where(ps ...predicate.TelegramChannel) *TelegramChannelUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetAccessHash sets the "access_hash" field.
func (tcu *TelegramChannelUpdate) SetAccessHash(i int64) *TelegramChannelUpdate {
	tcu.mutation.ResetAccessHash()
	tcu.mutation.SetAccessHash(i)
	return tcu
}

// SetNillableAccessHash sets the "access_hash" field if the given value is not nil.
func (tcu *TelegramChannelUpdate) SetNillableAccessHash(i *int64) *TelegramChannelUpdate {
	if i != nil {
		tcu.SetAccessHash(*i)
	}
	return tcu
}

// AddAccessHash adds i to the "access_hash" field.
func (tcu *TelegramChannelUpdate) AddAccessHash(i int64) *TelegramChannelUpdate {
	tcu.mutation.AddAccessHash(i)
	return tcu
}

// SetTitle sets the "title" field.
func (tcu *TelegramChannelUpdate) SetTitle(s string) *TelegramChannelUpdate {
	tcu.mutation.SetTitle(s)
	return tcu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tcu *TelegramChannelUpdate) SetNillableTitle(s *string) *TelegramChannelUpdate {
	if s != nil {
		tcu.SetTitle(*s)
	}
	return tcu
}

// SetSaveRecords sets the "save_records" field.
func (tcu *TelegramChannelUpdate) SetSaveRecords(b bool) *TelegramChannelUpdate {
	tcu.mutation.SetSaveRecords(b)
	return tcu
}

// SetNillableSaveRecords sets the "save_records" field if the given value is not nil.
func (tcu *TelegramChannelUpdate) SetNillableSaveRecords(b *bool) *TelegramChannelUpdate {
	if b != nil {
		tcu.SetSaveRecords(*b)
	}
	return tcu
}

// ClearSaveRecords clears the value of the "save_records" field.
func (tcu *TelegramChannelUpdate) ClearSaveRecords() *TelegramChannelUpdate {
	tcu.mutation.ClearSaveRecords()
	return tcu
}

// SetSaveFavoriteRecords sets the "save_favorite_records" field.
func (tcu *TelegramChannelUpdate) SetSaveFavoriteRecords(b bool) *TelegramChannelUpdate {
	tcu.mutation.SetSaveFavoriteRecords(b)
	return tcu
}

// SetNillableSaveFavoriteRecords sets the "save_favorite_records" field if the given value is not nil.
func (tcu *TelegramChannelUpdate) SetNillableSaveFavoriteRecords(b *bool) *TelegramChannelUpdate {
	if b != nil {
		tcu.SetSaveFavoriteRecords(*b)
	}
	return tcu
}

// ClearSaveFavoriteRecords clears the value of the "save_favorite_records" field.
func (tcu *TelegramChannelUpdate) ClearSaveFavoriteRecords() *TelegramChannelUpdate {
	tcu.mutation.ClearSaveFavoriteRecords()
	return tcu
}

// SetActive sets the "active" field.
func (tcu *TelegramChannelUpdate) SetActive(b bool) *TelegramChannelUpdate {
	tcu.mutation.SetActive(b)
	return tcu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (tcu *TelegramChannelUpdate) SetNillableActive(b *bool) *TelegramChannelUpdate {
	if b != nil {
		tcu.SetActive(*b)
	}
	return tcu
}

// Mutation returns the TelegramChannelMutation object of the builder.
func (tcu *TelegramChannelUpdate) Mutation() *TelegramChannelMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TelegramChannelUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tcu.sqlSave, tcu.mutation, tcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TelegramChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TelegramChannelUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TelegramChannelUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tcu *TelegramChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(telegramchannel.Table, telegramchannel.Columns, sqlgraph.NewFieldSpec(telegramchannel.FieldID, field.TypeInt64))
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.AccessHash(); ok {
		_spec.SetField(telegramchannel.FieldAccessHash, field.TypeInt64, value)
	}
	if value, ok := tcu.mutation.AddedAccessHash(); ok {
		_spec.AddField(telegramchannel.FieldAccessHash, field.TypeInt64, value)
	}
	if value, ok := tcu.mutation.Title(); ok {
		_spec.SetField(telegramchannel.FieldTitle, field.TypeString, value)
	}
	if value, ok := tcu.mutation.SaveRecords(); ok {
		_spec.SetField(telegramchannel.FieldSaveRecords, field.TypeBool, value)
	}
	if tcu.mutation.SaveRecordsCleared() {
		_spec.ClearField(telegramchannel.FieldSaveRecords, field.TypeBool)
	}
	if value, ok := tcu.mutation.SaveFavoriteRecords(); ok {
		_spec.SetField(telegramchannel.FieldSaveFavoriteRecords, field.TypeBool, value)
	}
	if tcu.mutation.SaveFavoriteRecordsCleared() {
		_spec.ClearField(telegramchannel.FieldSaveFavoriteRecords, field.TypeBool)
	}
	if value, ok := tcu.mutation.Active(); ok {
		_spec.SetField(telegramchannel.FieldActive, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcu.mutation.done = true
	return n, nil
}

// TelegramChannelUpdateOne is the builder for updating a single TelegramChannel entity.
type TelegramChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TelegramChannelMutation
}

// SetAccessHash sets the "access_hash" field.
func (tcuo *TelegramChannelUpdateOne) SetAccessHash(i int64) *TelegramChannelUpdateOne {
	tcuo.mutation.ResetAccessHash()
	tcuo.mutation.SetAccessHash(i)
	return tcuo
}

// SetNillableAccessHash sets the "access_hash" field if the given value is not nil.
func (tcuo *TelegramChannelUpdateOne) SetNillableAccessHash(i *int64) *TelegramChannelUpdateOne {
	if i != nil {
		tcuo.SetAccessHash(*i)
	}
	return tcuo
}

// AddAccessHash adds i to the "access_hash" field.
func (tcuo *TelegramChannelUpdateOne) AddAccessHash(i int64) *TelegramChannelUpdateOne {
	tcuo.mutation.AddAccessHash(i)
	return tcuo
}

// SetTitle sets the "title" field.
func (tcuo *TelegramChannelUpdateOne) SetTitle(s string) *TelegramChannelUpdateOne {
	tcuo.mutation.SetTitle(s)
	return tcuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tcuo *TelegramChannelUpdateOne) SetNillableTitle(s *string) *TelegramChannelUpdateOne {
	if s != nil {
		tcuo.SetTitle(*s)
	}
	return tcuo
}

// SetSaveRecords sets the "save_records" field.
func (tcuo *TelegramChannelUpdateOne) SetSaveRecords(b bool) *TelegramChannelUpdateOne {
	tcuo.mutation.SetSaveRecords(b)
	return tcuo
}

// SetNillableSaveRecords sets the "save_records" field if the given value is not nil.
func (tcuo *TelegramChannelUpdateOne) SetNillableSaveRecords(b *bool) *TelegramChannelUpdateOne {
	if b != nil {
		tcuo.SetSaveRecords(*b)
	}
	return tcuo
}

// ClearSaveRecords clears the value of the "save_records" field.
func (tcuo *TelegramChannelUpdateOne) ClearSaveRecords() *TelegramChannelUpdateOne {
	tcuo.mutation.ClearSaveRecords()
	return tcuo
}

// SetSaveFavoriteRecords sets the "save_favorite_records" field.
func (tcuo *TelegramChannelUpdateOne) SetSaveFavoriteRecords(b bool) *TelegramChannelUpdateOne {
	tcuo.mutation.SetSaveFavoriteRecords(b)
	return tcuo
}

// SetNillableSaveFavoriteRecords sets the "save_favorite_records" field if the given value is not nil.
func (tcuo *TelegramChannelUpdateOne) SetNillableSaveFavoriteRecords(b *bool) *TelegramChannelUpdateOne {
	if b != nil {
		tcuo.SetSaveFavoriteRecords(*b)
	}
	return tcuo
}

// ClearSaveFavoriteRecords clears the value of the "save_favorite_records" field.
func (tcuo *TelegramChannelUpdateOne) ClearSaveFavoriteRecords() *TelegramChannelUpdateOne {
	tcuo.mutation.ClearSaveFavoriteRecords()
	return tcuo
}

// SetActive sets the "active" field.
func (tcuo *TelegramChannelUpdateOne) SetActive(b bool) *TelegramChannelUpdateOne {
	tcuo.mutation.SetActive(b)
	return tcuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (tcuo *TelegramChannelUpdateOne) SetNillableActive(b *bool) *TelegramChannelUpdateOne {
	if b != nil {
		tcuo.SetActive(*b)
	}
	return tcuo
}

// Mutation returns the TelegramChannelMutation object of the builder.
func (tcuo *TelegramChannelUpdateOne) Mutation() *TelegramChannelMutation {
	return tcuo.mutation
}

// Where appends a list predicates to the TelegramChannelUpdate builder.
func (tcuo *TelegramChannelUpdateOne) Where(ps ...predicate.TelegramChannel) *TelegramChannelUpdateOne {
	tcuo.mutation.Where(ps...)
	return tcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TelegramChannelUpdateOne) Select(field string, fields ...string) *TelegramChannelUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TelegramChannel entity.
func (tcuo *TelegramChannelUpdateOne) Save(ctx context.Context) (*TelegramChannel, error) {
	return withHooks(ctx, tcuo.sqlSave, tcuo.mutation, tcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TelegramChannelUpdateOne) SaveX(ctx context.Context) *TelegramChannel {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TelegramChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TelegramChannelUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tcuo *TelegramChannelUpdateOne) sqlSave(ctx context.Context) (_node *TelegramChannel, err error) {
	_spec := sqlgraph.NewUpdateSpec(telegramchannel.Table, telegramchannel.Columns, sqlgraph.NewFieldSpec(telegramchannel.FieldID, field.TypeInt64))
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TelegramChannel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, telegramchannel.FieldID)
		for _, f := range fields {
			if !telegramchannel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != telegramchannel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.AccessHash(); ok {
		_spec.SetField(telegramchannel.FieldAccessHash, field.TypeInt64, value)
	}
	if value, ok := tcuo.mutation.AddedAccessHash(); ok {
		_spec.AddField(telegramchannel.FieldAccessHash, field.TypeInt64, value)
	}
	if value, ok := tcuo.mutation.Title(); ok {
		_spec.SetField(telegramchannel.FieldTitle, field.TypeString, value)
	}
	if value, ok := tcuo.mutation.SaveRecords(); ok {
		_spec.SetField(telegramchannel.FieldSaveRecords, field.TypeBool, value)
	}
	if tcuo.mutation.SaveRecordsCleared() {
		_spec.ClearField(telegramchannel.FieldSaveRecords, field.TypeBool)
	}
	if value, ok := tcuo.mutation.SaveFavoriteRecords(); ok {
		_spec.SetField(telegramchannel.FieldSaveFavoriteRecords, field.TypeBool, value)
	}
	if tcuo.mutation.SaveFavoriteRecordsCleared() {
		_spec.ClearField(telegramchannel.FieldSaveFavoriteRecords, field.TypeBool)
	}
	if value, ok := tcuo.mutation.Active(); ok {
		_spec.SetField(telegramchannel.FieldActive, field.TypeBool, value)
	}
	_node = &TelegramChannel{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcuo.mutation.done = true
	return _node, nil
}
