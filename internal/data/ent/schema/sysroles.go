package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SysRoles holds the schema definition for the SysRoles entity.
type SysRoles struct {
	ent.Schema
}

func (SysRoles) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_roles"},
	}
}

// Fields of the SysRoles.
func (SysRoles) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.String("guid").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.Int64("merchant_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Comment("商户ID"),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Comment("角色名称"),

		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Comment("角色标识"),

		field.Int32("create_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),

		field.Int32("update_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),
	}

}

// Edges of the SysRoles.
func (SysRoles) Edges() []ent.Edge {
	return nil
}
