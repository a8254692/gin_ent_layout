package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SysRolePermissions holds the schema definition for the SysRolePermissions entity.
type SysRolePermissions struct {
	ent.Schema
}

func (SysRolePermissions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_role_permissions"},
	}
}

// Fields of the SysRolePermissions.
func (SysRolePermissions) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.Int64("role_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}),

		field.Int64("permission_id").SchemaType(map[string]string{
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

// Edges of the SysRolePermissions.
func (SysRolePermissions) Edges() []ent.Edge {
	return nil
}
