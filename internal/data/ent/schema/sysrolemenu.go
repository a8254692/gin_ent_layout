package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SysRoleMenu holds the schema definition for the SysRoleMenu entity.
type SysRoleMenu struct {
	ent.Schema
}

func (SysRoleMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_role_menu"},
	}
}

// Fields of the SysRoleMenu.
func (SysRoleMenu) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.Int64("role_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}),

		field.Int64("menu_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}),

		field.Int32("create_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),

		field.Int32("update_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),
	}

}

// Edges of the SysRoleMenu.
func (SysRoleMenu) Edges() []ent.Edge {
	return nil
}
