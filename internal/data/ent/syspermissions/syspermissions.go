// Code generated by ent, DO NOT EDIT.

package syspermissions

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the syspermissions type in the database.
	Label = "sys_permissions"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMenuID holds the string denoting the menu_id field in the database.
	FieldMenuID = "menu_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldCommand holds the string denoting the command field in the database.
	FieldCommand = "command"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the syspermissions in the database.
	Table = "sys_permissions"
)

// Columns holds all SQL columns for syspermissions fields.
var Columns = []string{
	FieldID,
	FieldMenuID,
	FieldName,
	FieldCode,
	FieldPath,
	FieldCommand,
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

var (
	// DefaultMenuID holds the default value on creation for the "menu_id" field.
	DefaultMenuID int64
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime int32
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime int32
)

// OrderOption defines the ordering options for the SysPermissions queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMenuID orders the results by the menu_id field.
func ByMenuID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByCommand orders the results by the command field.
func ByCommand(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommand, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
