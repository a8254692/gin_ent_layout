package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MerchantList holds the schema definition for the MerchantList entity.
type MerchantList struct {
	ent.Schema
}

func (MerchantList) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "merchant_list"},
	}
}

// Fields of the MerchantList.
func (MerchantList) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Unique(),

		field.String("guid").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(200)", // Override MySQL.
		}).Comment(" 商户名称"),

		field.String("host").SchemaType(map[string]string{
			dialect.MySQL: "varchar(200)", // Override MySQL.
		}).Comment("站点域名"),

		field.String("manage_host").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.Int16("status").SchemaType(map[string]string{
			dialect.MySQL: "smallint", // Override MySQL.
		}),

		field.Int32("create_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}),

		field.Int32("update_time").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}),
	}

}

// Edges of the MerchantList.
func (MerchantList) Edges() []ent.Edge {
	return nil
}
