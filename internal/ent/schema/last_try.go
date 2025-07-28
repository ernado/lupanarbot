package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type LastTry struct {
	ent.Schema
}

// Fields of the LastTry.
func (LastTry) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Time("try"),
		field.Enum("type").Values("Extremism", "Constitution", "CriminalCode"),
	}
}
