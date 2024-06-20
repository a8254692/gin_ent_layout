package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SysMenu holds the schema definition for the SysMenu entity.
type SysMenu struct {
	ent.Schema
}

func (SysMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menu"},
	}
}

// Fields of the SysMenu.
func (SysMenu) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.Int64("parent_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Default(0).Comment("上级菜单"),

		field.String("title").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Comment("标题"),

		field.String("icon").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Optional().Comment("图标"),

		field.String("uri").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Optional().Comment("路径"),

		field.Int8("show").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Default(1).Comment("是否展示:1是,0否"),

		field.Int32("sort").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0).Comment("排序"),

		field.Int32("create_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),

		field.Int32("update_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),
	}

}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return nil
}
