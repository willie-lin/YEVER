package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Quser holds the schema definition for the Quser entity.
type Quser struct {
	ent.Schema
}

// Annotations of the User.
func (Quser) Annotations() []schema.Annotation {
	return []schema.Annotation{

		entsql.Annotation{Table: "quser"},
	}
}

// Fields of the Quser.
func (Quser) Fields() []ent.Field {

	return []ent.Field{
		field.String("id").Unique(),
		field.String("qq"),
		field.String("phone"),
	}
}

// Edges of the Quser.
func (Quser) Edges() []ent.Edge {
	return nil
}
