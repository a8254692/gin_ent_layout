package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.String("guid").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.String("username").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.String("password").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.String("nick_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.Int64("merchant_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}),

		field.String("role").SchemaType(map[string]string{
			dialect.MySQL: "varchar(200)", // Override MySQL.
		}),

		field.String("avatar").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Optional(),

		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.String("phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.Int16("is_master").SchemaType(map[string]string{
			dialect.MySQL: "smallint", // Override MySQL.
		}),

		field.Int32("create_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}),

		field.Int32("update_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}),

		field.Int32("last_login_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}),
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
