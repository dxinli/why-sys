package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
	mixin.Time
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").
			NotEmpty().
			Unique(),
		field.String("password").
			NotEmpty(),
		field.Uint8("age").Range(0, 200).Optional(),
		field.Enum("gender").Values("Male", "Female").Default("Male"),
		field.String("email").
			NotEmpty().
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
