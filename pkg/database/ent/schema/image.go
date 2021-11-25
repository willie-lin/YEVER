package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Annotations of the Image.
func (Image) Annotations() []schema.Annotation {
	return []schema.Annotation{

		entsql.Annotation{Table: "images"},
	}
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return nil
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return nil
}
