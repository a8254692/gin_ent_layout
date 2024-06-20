package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SysPermissions holds the schema definition for the SysPermissions entity.
type SysPermissions struct {
	ent.Schema
}

func (SysPermissions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_permissions"},
	}
}

// Fields of the SysPermissions.
func (SysPermissions) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.Int64("menu_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Default(0).Comment("权限归属菜单"),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Comment("权限名称"),

		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Comment("标识"),

		field.String("path").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.String("command").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.Int32("create_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),

		field.Int32("update_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),
	}

}

// Edges of the SysPermissions.
func (SysPermissions) Edges() []ent.Edge {
	return nil
}
