// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGUID holds the string denoting the guid field in the database.
	FieldGUID = "guid"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// FieldMerchantID holds the string denoting the merchant_id field in the database.
	FieldMerchantID = "merchant_id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldIsMaster holds the string denoting the is_master field in the database.
	FieldIsMaster = "is_master"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldGUID,
	FieldUsername,
	FieldPassword,
	FieldNickName,
	FieldMerchantID,
	FieldRole,
	FieldAvatar,
	FieldEmail,
	FieldPhone,
	FieldIsMaster,
	FieldCreateTime,
	FieldUpdateTime,
	FieldLastLoginTime,
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

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGUID orders the results by the guid field.
func ByGUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGUID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByNickName orders the results by the nick_name field.
func ByNickName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickName, opts...).ToFunc()
}

// ByMerchantID orders the results by the merchant_id field.
func ByMerchantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMerchantID, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByIsMaster orders the results by the is_master field.
func ByIsMaster(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsMaster, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByLastLoginTime orders the results by the last_login_time field.
func ByLastLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginTime, opts...).ToFunc()
}