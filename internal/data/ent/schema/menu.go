package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("menu_name").NotEmpty(),
		field.Bool("leaf").Default(true),
		field.String("path").NotEmpty().Comment("菜单路径"),
		field.String("component").NotEmpty().Comment("前端组件"),
		field.Int("parent_id").Optional(),
		field.String("menu_desc").Optional(),
		field.Float("sort").Comment("排序，规则为 3.1 3表示父级菜单级别，1表示子级菜单序号"),
		field.Uint16("level").Comment("菜单级别"),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Menu.Type).
			From("parent").
			Field("parent_id").
			Unique(),
	}
}
