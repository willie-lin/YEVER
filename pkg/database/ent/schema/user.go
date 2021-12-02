package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{

		entsql.Annotation{Table: "users"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("name").MaxLen(120).Unique(),
		field.Int("age").Positive(),
		field.String("password").Sensitive(),
		field.String("email").MaxLen(120),
		field.String("phone").MaxLen(13).MinLen(11).Default("12345678910"),
		field.String("description").MaxLen(2000).Default("others"),
		field.Time("created").Default(time.Now).Immutable(),
		field.Time("updated").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
