// Code generated by ent, DO NOT EDIT.

package sysrolepermissions

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sysrolepermissions type in the database.
	Label = "sys_role_permissions"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldPermissionID holds the string denoting the permission_id field in the database.
	FieldPermissionID = "permission_id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the sysrolepermissions in the database.
	Table = "sys_role_permissions"
)

// Columns holds all SQL columns for sysrolepermissions fields.
var Columns = []string{
	FieldID,
	FieldRoleID,
	FieldPermissionID,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime int32
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime int32
)

// OrderOption defines the ordering options for the SysRolePermissions queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByPermissionID orders the results by the permission_id field.
func ByPermissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPermissionID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}