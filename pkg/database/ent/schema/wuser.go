package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Wuser holds the schema definition for the Wuser entity.
type Wuser struct {
	ent.Schema
}

func (Wuser) Annotations() []schema.Annotation {
	return []schema.Annotation{

		entsql.Annotation{Table: "wuser"},
	}
}

// Fields of the Wuser.
func (Wuser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("phone"),
		field.String("uid"),
	}
}

// Edges of the Wuser.
func (Wuser) Edges() []ent.Edge {
	return nil
}
