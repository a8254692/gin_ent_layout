// Code generated by ent, DO NOT EDIT.

package ent

import (
	"back-end/merchant/internal/data/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// GUID holds the value of the "guid" field.
	GUID string `json:"guid,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// NickName holds the value of the "nick_name" field.
	NickName string `json:"nick_name,omitempty"`
	// MerchantID holds the value of the "merchant_id" field.
	MerchantID int64 `json:"merchant_id,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// IsMaster holds the value of the "is_master" field.
	IsMaster int16 `json:"is_master,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime int32 `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime int32 `json:"update_time,omitempty"`
	// LastLoginTime holds the value of the "last_login_time" field.
	LastLoginTime int32 `json:"last_login_time,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldMerchantID, user.FieldIsMaster, user.FieldCreateTime, user.FieldUpdateTime, user.FieldLastLoginTime:
			values[i] = new(sql.NullInt64)
		case user.FieldGUID, user.FieldUsername, user.FieldPassword, user.FieldNickName, user.FieldRole, user.FieldAvatar, user.FieldEmail, user.FieldPhone:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldGUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field guid", values[i])
			} else if value.Valid {
				u.GUID = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				u.NickName = value.String
			}
		case user.FieldMerchantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_id", values[i])
			} else if value.Valid {
				u.MerchantID = value.Int64
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldIsMaster:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_master", values[i])
			} else if value.Valid {
				u.IsMaster = int16(value.Int64)
			}
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = int32(value.Int64)
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = int32(value.Int64)
			}
		case user.FieldLastLoginTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_time", values[i])
			} else if value.Valid {
				u.LastLoginTime = int32(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("guid=")
	builder.WriteString(u.GUID)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("nick_name=")
	builder.WriteString(u.NickName)
	builder.WriteString(", ")
	builder.WriteString("merchant_id=")
	builder.WriteString(fmt.Sprintf("%v", u.MerchantID))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(u.Role)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("is_master=")
	builder.WriteString(fmt.Sprintf("%v", u.IsMaster))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(fmt.Sprintf("%v", u.CreateTime))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdateTime))
	builder.WriteString(", ")
	builder.WriteString("last_login_time=")
	builder.WriteString(fmt.Sprintf("%v", u.LastLoginTime))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User