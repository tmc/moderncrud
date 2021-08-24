package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Widget holds the schema definition for the Widget entity.
type Widget struct {
	ent.Schema
}

// Fields of the Widget.
func (Widget) Fields() []ent.Field {
	return []ent.Field{
		field.Text("note").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			Values("draft", "completed").
			Default("draft"),
		field.Int("priority").
			Default(0),
	}
}

// Edges of the Widget.
func (Widget) Edges() []ent.Edge {
	return nil
}
