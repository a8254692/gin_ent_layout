package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SysRoleUsers holds the schema definition for the SysRoleUsers entity.
type SysRoleUsers struct {
	ent.Schema
}

func (SysRoleUsers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_role_users"},
	}
}

// Fields of the SysRoleUsers.
func (SysRoleUsers) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Unique(),

		field.Int64("role_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}),

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}),

		field.Int32("create_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}),

		field.Int32("update_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),
	}

}

// Edges of the SysRoleUsers.
func (SysRoleUsers) Edges() []ent.Edge {
	return nil
}
