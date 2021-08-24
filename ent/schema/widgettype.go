package schema

import "entgo.io/ent"

// WidgetType holds the schema definition for the WidgetType entity.
type WidgetType struct {
	ent.Schema
}

// Fields of the WidgetType.
func (WidgetType) Fields() []ent.Field {
	return nil
}

// Edges of the WidgetType.
func (WidgetType) Edges() []ent.Edge {
	return nil
}
