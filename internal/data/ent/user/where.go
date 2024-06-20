// Code generated by ent, DO NOT EDIT.

package user

import (
	"back-end/merchant/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// GUID applies equality check predicate on the "guid" field. It's identical to GUIDEQ.
func GUID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGUID, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// NickName applies equality check predicate on the "nick_name" field. It's identical to NickNameEQ.
func NickName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// MerchantID applies equality check predicate on the "merchant_id" field. It's identical to MerchantIDEQ.
func MerchantID(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMerchantID, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// IsMaster applies equality check predicate on the "is_master" field. It's identical to IsMasterEQ.
func IsMaster(v int16) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsMaster, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// LastLoginTime applies equality check predicate on the "last_login_time" field. It's identical to LastLoginTimeEQ.
func LastLoginTime(v int32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLoginTime, v))
}

// GUIDEQ applies the EQ predicate on the "guid" field.
func GUIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGUID, v))
}

// GUIDNEQ applies the NEQ predicate on the "guid" field.
func GUIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGUID, v))
}

// GUIDIn applies the In predicate on the "guid" field.
func GUIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGUID, vs...))
}

// GUIDNotIn applies the NotIn predicate on the "guid" field.
func GUIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGUID, vs...))
}

// GUIDGT applies the GT predicate on the "guid" field.
func GUIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGUID, v))
}

// GUIDGTE applies the GTE predicate on the "guid" field.
func GUIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGUID, v))
}

// GUIDLT applies the LT predicate on the "guid" field.
func GUIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGUID, v))
}

// GUIDLTE applies the LTE predicate on the "guid" field.
func GUIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGUID, v))
}

// GUIDContains applies the Contains predicate on the "guid" field.
func GUIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGUID, v))
}

// GUIDHasPrefix applies the HasPrefix predicate on the "guid" field.
func GUIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGUID, v))
}

// GUIDHasSuffix applies the HasSuffix predicate on the "guid" field.
func GUIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGUID, v))
}

// GUIDEqualFold applies the EqualFold predicate on the "guid" field.
func GUIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGUID, v))
}

// GUIDContainsFold applies the ContainsFold predicate on the "guid" field.
func GUIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGUID, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// NickNameEQ applies the EQ predicate on the "nick_name" field.
func NickNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// NickNameNEQ applies the NEQ predicate on the "nick_name" field.
func NickNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickName, v))
}

// NickNameIn applies the In predicate on the "nick_name" field.
func NickNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickName, vs...))
}

// NickNameNotIn applies the NotIn predicate on the "nick_name" field.
func NickNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickName, vs...))
}

// NickNameGT applies the GT predicate on the "nick_name" field.
func NickNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickName, v))
}

// NickNameGTE applies the GTE predicate on the "nick_name" field.
func NickNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickName, v))
}

// NickNameLT applies the LT predicate on the "nick_name" field.
func NickNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickName, v))
}

// NickNameLTE applies the LTE predicate on the "nick_name" field.
func NickNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickName, v))
}

// NickNameContains applies the Contains predicate on the "nick_name" field.
func NickNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickName, v))
}

// NickNameHasPrefix applies the HasPrefix predicate on the "nick_name" field.
func NickNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickName, v))
}

// NickNameHasSuffix applies the HasSuffix predicate on the "nick_name" field.
func NickNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickName, v))
}

// NickNameEqualFold applies the EqualFold predicate on the "nick_name" field.
func NickNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickName, v))
}

// NickNameContainsFold applies the ContainsFold predicate on the "nick_name" field.
func NickNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickName, v))
}

// MerchantIDEQ applies the EQ predicate on the "merchant_id" field.
func MerchantIDEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMerchantID, v))
}

// MerchantIDNEQ applies the NEQ predicate on the "merchant_id" field.
func MerchantIDNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMerchantID, v))
}

// MerchantIDIn applies the In predicate on the "merchant_id" field.
func MerchantIDIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldMerchantID, vs...))
}

// MerchantIDNotIn applies the NotIn predicate on the "merchant_id" field.
func MerchantIDNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMerchantID, vs...))
}

// MerchantIDGT applies the GT predicate on the "merchant_id" field.
func MerchantIDGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldMerchantID, v))
}

// MerchantIDGTE applies the GTE predicate on the "merchant_id" field.
func MerchantIDGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMerchantID, v))
}

// MerchantIDLT applies the LT predicate on the "merchant_id" field.
func MerchantIDLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldMerchantID, v))
}

// MerchantIDLTE applies the LTE predicate on the "merchant_id" field.
func MerchantIDLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMerchantID, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRole, v))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRole, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatar))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatar, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// IsMasterEQ applies the EQ predicate on the "is_master" field.
func IsMasterEQ(v int16) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsMaster, v))
}

// IsMasterNEQ applies the NEQ predicate on the "is_master" field.
func IsMasterNEQ(v int16) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsMaster, v))
}

// IsMasterIn applies the In predicate on the "is_master" field.
func IsMasterIn(vs ...int16) predicate.User {
	return predicate.User(sql.FieldIn(FieldIsMaster, vs...))
}

// IsMasterNotIn applies the NotIn predicate on the "is_master" field.
func IsMasterNotIn(vs ...int16) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldIsMaster, vs...))
}

// IsMasterGT applies the GT predicate on the "is_master" field.
func IsMasterGT(v int16) predicate.User {
	return predicate.User(sql.FieldGT(FieldIsMaster, v))
}

// IsMasterGTE applies the GTE predicate on the "is_master" field.
func IsMasterGTE(v int16) predicate.User {
	return predicate.User(sql.FieldGTE(FieldIsMaster, v))
}

// IsMasterLT applies the LT predicate on the "is_master" field.
func IsMasterLT(v int16) predicate.User {
	return predicate.User(sql.FieldLT(FieldIsMaster, v))
}

// IsMasterLTE applies the LTE predicate on the "is_master" field.
func IsMasterLTE(v int16) predicate.User {
	return predicate.User(sql.FieldLTE(FieldIsMaster, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int32) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int32) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int32) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int32) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int32) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int32) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int32) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int32) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int32) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int32) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int32) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int32) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int32) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int32) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdateTime, v))
}

// LastLoginTimeEQ applies the EQ predicate on the "last_login_time" field.
func LastLoginTimeEQ(v int32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLoginTime, v))
}

// LastLoginTimeNEQ applies the NEQ predicate on the "last_login_time" field.
func LastLoginTimeNEQ(v int32) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastLoginTime, v))
}

// LastLoginTimeIn applies the In predicate on the "last_login_time" field.
func LastLoginTimeIn(vs ...int32) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastLoginTime, vs...))
}

// LastLoginTimeNotIn applies the NotIn predicate on the "last_login_time" field.
func LastLoginTimeNotIn(vs ...int32) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastLoginTime, vs...))
}

// LastLoginTimeGT applies the GT predicate on the "last_login_time" field.
func LastLoginTimeGT(v int32) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastLoginTime, v))
}

// LastLoginTimeGTE applies the GTE predicate on the "last_login_time" field.
func LastLoginTimeGTE(v int32) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastLoginTime, v))
}

// LastLoginTimeLT applies the LT predicate on the "last_login_time" field.
func LastLoginTimeLT(v int32) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastLoginTime, v))
}

// LastLoginTimeLTE applies the LTE predicate on the "last_login_time" field.
func LastLoginTimeLTE(v int32) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastLoginTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}