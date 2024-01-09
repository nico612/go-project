package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_at").Default(time.Now()),
		field.Int("age").Positive(),
		field.String("username").Unique(),
		field.String("name").Default("unknown"),
	}
}

// Edges of the User. one to many cars
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),

		// Create an inverse-edge called "groups" of type `Group`
		// and reference it to the "users" edge (in Group schema)
		// explicitly using the `Ref` method.
		edge.From("groups", Group.Type).Ref("users"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("age", "name").Unique(),
	}
}
