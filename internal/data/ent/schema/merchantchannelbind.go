package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MerchantChannelBind holds the schema definition for the MerchantChannelBind entity.
type MerchantChannelBind struct {
	ent.Schema
}

func (MerchantChannelBind) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "merchant_channel_bind"},
	}
}

// Fields of the MerchantChannelBind.
func (MerchantChannelBind) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Unique(),

		field.Int32("merchant_id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Comment("商户 Id"),

		field.Int32("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Comment("渠道 ID"),

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

// Edges of the MerchantChannelBind.
func (MerchantChannelBind) Edges() []ent.Edge {
	return nil
}
