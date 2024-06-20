// Code generated by ent, DO NOT EDIT.

package merchantlist

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the merchantlist type in the database.
	Label = "merchant_list"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGUID holds the string denoting the guid field in the database.
	FieldGUID = "guid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldManageHost holds the string denoting the manage_host field in the database.
	FieldManageHost = "manage_host"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the merchantlist in the database.
	Table = "merchant_list"
)

// Columns holds all SQL columns for merchantlist fields.
var Columns = []string{
	FieldID,
	FieldGUID,
	FieldName,
	FieldHost,
	FieldManageHost,
	FieldStatus,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the MerchantList queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGUID orders the results by the guid field.
func ByGUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGUID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByHost orders the results by the host field.
func ByHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHost, opts...).ToFunc()
}

// ByManageHost orders the results by the manage_host field.
func ByManageHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManageHost, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}