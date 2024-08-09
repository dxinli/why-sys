package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_name").NotEmpty(),
		field.String("role_desc").Optional(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
