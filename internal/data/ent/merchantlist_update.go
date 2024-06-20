// Code generated by ent, DO NOT EDIT.

package ent

import (
	"back-end/merchant/internal/data/ent/merchantlist"
	"back-end/merchant/internal/data/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MerchantListUpdate is the builder for updating MerchantList entities.
type MerchantListUpdate struct {
	config
	hooks    []Hook
	mutation *MerchantListMutation
}

// Where appends a list predicates to the MerchantListUpdate builder.
func (mlu *MerchantListUpdate) Where(ps ...predicate.MerchantList) *MerchantListUpdate {
	mlu.mutation.Where(ps...)
	return mlu
}

// SetGUID sets the "guid" field.
func (mlu *MerchantListUpdate) SetGUID(s string) *MerchantListUpdate {
	mlu.mutation.SetGUID(s)
	return mlu
}

// SetNillableGUID sets the "guid" field if the given value is not nil.
func (mlu *MerchantListUpdate) SetNillableGUID(s *string) *MerchantListUpdate {
	if s != nil {
		mlu.SetGUID(*s)
	}
	return mlu
}

// SetName sets the "name" field.
func (mlu *MerchantListUpdate) SetName(s string) *MerchantListUpdate {
	mlu.mutation.SetName(s)
	return mlu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mlu *MerchantListUpdate) SetNillableName(s *string) *MerchantListUpdate {
	if s != nil {
		mlu.SetName(*s)
	}
	return mlu
}

// SetHost sets the "host" field.
func (mlu *MerchantListUpdate) SetHost(s string) *MerchantListUpdate {
	mlu.mutation.SetHost(s)
	return mlu
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (mlu *MerchantListUpdate) SetNillableHost(s *string) *MerchantListUpdate {
	if s != nil {
		mlu.SetHost(*s)
	}
	return mlu
}

// SetManageHost sets the "manage_host" field.
func (mlu *MerchantListUpdate) SetManageHost(s string) *MerchantListUpdate {
	mlu.mutation.SetManageHost(s)
	return mlu
}

// SetNillableManageHost sets the "manage_host" field if the given value is not nil.
func (mlu *MerchantListUpdate) SetNillableManageHost(s *string) *MerchantListUpdate {
	if s != nil {
		mlu.SetManageHost(*s)
	}
	return mlu
}

// SetStatus sets the "status" field.
func (mlu *MerchantListUpdate) SetStatus(i int16) *MerchantListUpdate {
	mlu.mutation.ResetStatus()
	mlu.mutation.SetStatus(i)
	return mlu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mlu *MerchantListUpdate) SetNillableStatus(i *int16) *MerchantListUpdate {
	if i != nil {
		mlu.SetStatus(*i)
	}
	return mlu
}

// AddStatus adds i to the "status" field.
func (mlu *MerchantListUpdate) AddStatus(i int16) *MerchantListUpdate {
	mlu.mutation.AddStatus(i)
	return mlu
}

// SetCreateTime sets the "create_time" field.
func (mlu *MerchantListUpdate) SetCreateTime(i int32) *MerchantListUpdate {
	mlu.mutation.ResetCreateTime()
	mlu.mutation.SetCreateTime(i)
	return mlu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mlu *MerchantListUpdate) SetNillableCreateTime(i *int32) *MerchantListUpdate {
	if i != nil {
		mlu.SetCreateTime(*i)
	}
	return mlu
}

// AddCreateTime adds i to the "create_time" field.
func (mlu *MerchantListUpdate) AddCreateTime(i int32) *MerchantListUpdate {
	mlu.mutation.AddCreateTime(i)
	return mlu
}

// SetUpdateTime sets the "update_time" field.
func (mlu *MerchantListUpdate) SetUpdateTime(i int32) *MerchantListUpdate {
	mlu.mutation.ResetUpdateTime()
	mlu.mutation.SetUpdateTime(i)
	return mlu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mlu *MerchantListUpdate) SetNillableUpdateTime(i *int32) *MerchantListUpdate {
	if i != nil {
		mlu.SetUpdateTime(*i)
	}
	return mlu
}

// AddUpdateTime adds i to the "update_time" field.
func (mlu *MerchantListUpdate) AddUpdateTime(i int32) *MerchantListUpdate {
	mlu.mutation.AddUpdateTime(i)
	return mlu
}

// Mutation returns the MerchantListMutation object of the builder.
func (mlu *MerchantListUpdate) Mutation() *MerchantListMutation {
	return mlu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mlu *MerchantListUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mlu.sqlSave, mlu.mutation, mlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mlu *MerchantListUpdate) SaveX(ctx context.Context) int {
	affected, err := mlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mlu *MerchantListUpdate) Exec(ctx context.Context) error {
	_, err := mlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlu *MerchantListUpdate) ExecX(ctx context.Context) {
	if err := mlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mlu *MerchantListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(merchantlist.Table, merchantlist.Columns, sqlgraph.NewFieldSpec(merchantlist.FieldID, field.TypeInt32))
	if ps := mlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mlu.mutation.GUID(); ok {
		_spec.SetField(merchantlist.FieldGUID, field.TypeString, value)
	}
	if value, ok := mlu.mutation.Name(); ok {
		_spec.SetField(merchantlist.FieldName, field.TypeString, value)
	}
	if value, ok := mlu.mutation.Host(); ok {
		_spec.SetField(merchantlist.FieldHost, field.TypeString, value)
	}
	if value, ok := mlu.mutation.ManageHost(); ok {
		_spec.SetField(merchantlist.FieldManageHost, field.TypeString, value)
	}
	if value, ok := mlu.mutation.Status(); ok {
		_spec.SetField(merchantlist.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := mlu.mutation.AddedStatus(); ok {
		_spec.AddField(merchantlist.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := mlu.mutation.CreateTime(); ok {
		_spec.SetField(merchantlist.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := mlu.mutation.AddedCreateTime(); ok {
		_spec.AddField(merchantlist.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := mlu.mutation.UpdateTime(); ok {
		_spec.SetField(merchantlist.FieldUpdateTime, field.TypeInt32, value)
	}
	if value, ok := mlu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(merchantlist.FieldUpdateTime, field.TypeInt32, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchantlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mlu.mutation.done = true
	return n, nil
}

// MerchantListUpdateOne is the builder for updating a single MerchantList entity.
type MerchantListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MerchantListMutation
}

// SetGUID sets the "guid" field.
func (mluo *MerchantListUpdateOne) SetGUID(s string) *MerchantListUpdateOne {
	mluo.mutation.SetGUID(s)
	return mluo
}

// SetNillableGUID sets the "guid" field if the given value is not nil.
func (mluo *MerchantListUpdateOne) SetNillableGUID(s *string) *MerchantListUpdateOne {
	if s != nil {
		mluo.SetGUID(*s)
	}
	return mluo
}

// SetName sets the "name" field.
func (mluo *MerchantListUpdateOne) SetName(s string) *MerchantListUpdateOne {
	mluo.mutation.SetName(s)
	return mluo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mluo *MerchantListUpdateOne) SetNillableName(s *string) *MerchantListUpdateOne {
	if s != nil {
		mluo.SetName(*s)
	}
	return mluo
}

// SetHost sets the "host" field.
func (mluo *MerchantListUpdateOne) SetHost(s string) *MerchantListUpdateOne {
	mluo.mutation.SetHost(s)
	return mluo
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (mluo *MerchantListUpdateOne) SetNillableHost(s *string) *MerchantListUpdateOne {
	if s != nil {
		mluo.SetHost(*s)
	}
	return mluo
}

// SetManageHost sets the "manage_host" field.
func (mluo *MerchantListUpdateOne) SetManageHost(s string) *MerchantListUpdateOne {
	mluo.mutation.SetManageHost(s)
	return mluo
}

// SetNillableManageHost sets the "manage_host" field if the given value is not nil.
func (mluo *MerchantListUpdateOne) SetNillableManageHost(s *string) *MerchantListUpdateOne {
	if s != nil {
		mluo.SetManageHost(*s)
	}
	return mluo
}

// SetStatus sets the "status" field.
func (mluo *MerchantListUpdateOne) SetStatus(i int16) *MerchantListUpdateOne {
	mluo.mutation.ResetStatus()
	mluo.mutation.SetStatus(i)
	return mluo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mluo *MerchantListUpdateOne) SetNillableStatus(i *int16) *MerchantListUpdateOne {
	if i != nil {
		mluo.SetStatus(*i)
	}
	return mluo
}

// AddStatus adds i to the "status" field.
func (mluo *MerchantListUpdateOne) AddStatus(i int16) *MerchantListUpdateOne {
	mluo.mutation.AddStatus(i)
	return mluo
}

// SetCreateTime sets the "create_time" field.
func (mluo *MerchantListUpdateOne) SetCreateTime(i int32) *MerchantListUpdateOne {
	mluo.mutation.ResetCreateTime()
	mluo.mutation.SetCreateTime(i)
	return mluo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mluo *MerchantListUpdateOne) SetNillableCreateTime(i *int32) *MerchantListUpdateOne {
	if i != nil {
		mluo.SetCreateTime(*i)
	}
	return mluo
}

// AddCreateTime adds i to the "create_time" field.
func (mluo *MerchantListUpdateOne) AddCreateTime(i int32) *MerchantListUpdateOne {
	mluo.mutation.AddCreateTime(i)
	return mluo
}

// SetUpdateTime sets the "update_time" field.
func (mluo *MerchantListUpdateOne) SetUpdateTime(i int32) *MerchantListUpdateOne {
	mluo.mutation.ResetUpdateTime()
	mluo.mutation.SetUpdateTime(i)
	return mluo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mluo *MerchantListUpdateOne) SetNillableUpdateTime(i *int32) *MerchantListUpdateOne {
	if i != nil {
		mluo.SetUpdateTime(*i)
	}
	return mluo
}

// AddUpdateTime adds i to the "update_time" field.
func (mluo *MerchantListUpdateOne) AddUpdateTime(i int32) *MerchantListUpdateOne {
	mluo.mutation.AddUpdateTime(i)
	return mluo
}

// Mutation returns the MerchantListMutation object of the builder.
func (mluo *MerchantListUpdateOne) Mutation() *MerchantListMutation {
	return mluo.mutation
}

// Where appends a list predicates to the MerchantListUpdate builder.
func (mluo *MerchantListUpdateOne) Where(ps ...predicate.MerchantList) *MerchantListUpdateOne {
	mluo.mutation.Where(ps...)
	return mluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mluo *MerchantListUpdateOne) Select(field string, fields ...string) *MerchantListUpdateOne {
	mluo.fields = append([]string{field}, fields...)
	return mluo
}

// Save executes the query and returns the updated MerchantList entity.
func (mluo *MerchantListUpdateOne) Save(ctx context.Context) (*MerchantList, error) {
	return withHooks(ctx, mluo.sqlSave, mluo.mutation, mluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mluo *MerchantListUpdateOne) SaveX(ctx context.Context) *MerchantList {
	node, err := mluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mluo *MerchantListUpdateOne) Exec(ctx context.Context) error {
	_, err := mluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mluo *MerchantListUpdateOne) ExecX(ctx context.Context) {
	if err := mluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mluo *MerchantListUpdateOne) sqlSave(ctx context.Context) (_node *MerchantList, err error) {
	_spec := sqlgraph.NewUpdateSpec(merchantlist.Table, merchantlist.Columns, sqlgraph.NewFieldSpec(merchantlist.FieldID, field.TypeInt32))
	id, ok := mluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MerchantList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchantlist.FieldID)
		for _, f := range fields {
			if !merchantlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != merchantlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mluo.mutation.GUID(); ok {
		_spec.SetField(merchantlist.FieldGUID, field.TypeString, value)
	}
	if value, ok := mluo.mutation.Name(); ok {
		_spec.SetField(merchantlist.FieldName, field.TypeString, value)
	}
	if value, ok := mluo.mutation.Host(); ok {
		_spec.SetField(merchantlist.FieldHost, field.TypeString, value)
	}
	if value, ok := mluo.mutation.ManageHost(); ok {
		_spec.SetField(merchantlist.FieldManageHost, field.TypeString, value)
	}
	if value, ok := mluo.mutation.Status(); ok {
		_spec.SetField(merchantlist.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := mluo.mutation.AddedStatus(); ok {
		_spec.AddField(merchantlist.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := mluo.mutation.CreateTime(); ok {
		_spec.SetField(merchantlist.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := mluo.mutation.AddedCreateTime(); ok {
		_spec.AddField(merchantlist.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := mluo.mutation.UpdateTime(); ok {
		_spec.SetField(merchantlist.FieldUpdateTime, field.TypeInt32, value)
	}
	if value, ok := mluo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(merchantlist.FieldUpdateTime, field.TypeInt32, value)
	}
	_node = &MerchantList{config: mluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchantlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mluo.mutation.done = true
	return _node, nil
}
