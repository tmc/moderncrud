package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// WidgetType holds the schema definition for the WidgetType entity.
type WidgetType struct {
	ent.Schema
}

// Fields of the WidgetType.
func (WidgetType) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name"),
	}
}

// Edges of the WidgetType.
func (WidgetType) Edges() []ent.Edge {
	return nil
}
