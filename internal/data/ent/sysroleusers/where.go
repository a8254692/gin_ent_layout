// Code generated by ent, DO NOT EDIT.

package sysroleusers

import (
	"back-end/merchant/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLTE(FieldID, id))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldRoleID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldUserID, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldUpdateTime, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLTE(FieldRoleID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLTE(FieldUserID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int32) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.FieldLTE(FieldUpdateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysRoleUsers) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysRoleUsers) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysRoleUsers) predicate.SysRoleUsers {
	return predicate.SysRoleUsers(sql.NotPredicates(p))
}