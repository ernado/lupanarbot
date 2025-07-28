package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Try struct {
	ent.Schema
}

// Fields of the Try.
func (Try) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int64("user_id"),
		field.Time("created_at"),
		field.Enum("type").Values("Extremism", "Constitution", "CriminalCode"),
	}
}

func (Try) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "type").Unique(),
	}
}
