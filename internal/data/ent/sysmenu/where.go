// Code generated by ent, DO NOT EDIT.

package sysmenu

import (
	"back-end/merchant/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldID, id))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldParentID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldTitle, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIcon, v))
}

// URI applies equality check predicate on the "uri" field. It's identical to URIEQ.
func URI(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldURI, v))
}

// Show applies equality check predicate on the "show" field. It's identical to ShowEQ.
func Show(v int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldShow, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldSort, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldUpdateTime, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v int64) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldParentID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldTitle, v))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldIcon, v))
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldIcon))
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldIcon))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldIcon, v))
}

// URIEQ applies the EQ predicate on the "uri" field.
func URIEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldURI, v))
}

// URINEQ applies the NEQ predicate on the "uri" field.
func URINEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldURI, v))
}

// URIIn applies the In predicate on the "uri" field.
func URIIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldURI, vs...))
}

// URINotIn applies the NotIn predicate on the "uri" field.
func URINotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldURI, vs...))
}

// URIGT applies the GT predicate on the "uri" field.
func URIGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldURI, v))
}

// URIGTE applies the GTE predicate on the "uri" field.
func URIGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldURI, v))
}

// URILT applies the LT predicate on the "uri" field.
func URILT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldURI, v))
}

// URILTE applies the LTE predicate on the "uri" field.
func URILTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldURI, v))
}

// URIContains applies the Contains predicate on the "uri" field.
func URIContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldURI, v))
}

// URIHasPrefix applies the HasPrefix predicate on the "uri" field.
func URIHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldURI, v))
}

// URIHasSuffix applies the HasSuffix predicate on the "uri" field.
func URIHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldURI, v))
}

// URIIsNil applies the IsNil predicate on the "uri" field.
func URIIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldURI))
}

// URINotNil applies the NotNil predicate on the "uri" field.
func URINotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldURI))
}

// URIEqualFold applies the EqualFold predicate on the "uri" field.
func URIEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldURI, v))
}

// URIContainsFold applies the ContainsFold predicate on the "uri" field.
func URIContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldURI, v))
}

// ShowEQ applies the EQ predicate on the "show" field.
func ShowEQ(v int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldShow, v))
}

// ShowNEQ applies the NEQ predicate on the "show" field.
func ShowNEQ(v int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldShow, v))
}

// ShowIn applies the In predicate on the "show" field.
func ShowIn(vs ...int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldShow, vs...))
}

// ShowNotIn applies the NotIn predicate on the "show" field.
func ShowNotIn(vs ...int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldShow, vs...))
}

// ShowGT applies the GT predicate on the "show" field.
func ShowGT(v int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldShow, v))
}

// ShowGTE applies the GTE predicate on the "show" field.
func ShowGTE(v int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldShow, v))
}

// ShowLT applies the LT predicate on the "show" field.
func ShowLT(v int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldShow, v))
}

// ShowLTE applies the LTE predicate on the "show" field.
func ShowLTE(v int8) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldShow, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldSort, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldUpdateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(sql.NotPredicates(p))
}