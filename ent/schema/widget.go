package schema

import "entgo.io/ent"

// Widget holds the schema definition for the Widget entity.
type Widget struct {
	ent.Schema
}

// Fields of the Widget.
func (Widget) Fields() []ent.Field {
	return nil
}

// Edges of the Widget.
func (Widget) Edges() []ent.Edge {
	return nil
}
