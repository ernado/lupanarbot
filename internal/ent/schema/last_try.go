package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Try struct {
	ent.Schema
}

// Fields of the Try.
func (Try) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Time("created_at"),
		field.Enum("type").Values("Extremism", "Constitution", "CriminalCode"),
	}
}

func (Try) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "type").Unique(),
	}
}
