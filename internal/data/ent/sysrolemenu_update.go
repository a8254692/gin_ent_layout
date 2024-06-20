// Code generated by ent, DO NOT EDIT.

package ent

import (
	"back-end/merchant/internal/data/ent/predicate"
	"back-end/merchant/internal/data/ent/sysrolemenu"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysRoleMenuUpdate is the builder for updating SysRoleMenu entities.
type SysRoleMenuUpdate struct {
	config
	hooks    []Hook
	mutation *SysRoleMenuMutation
}

// Where appends a list predicates to the SysRoleMenuUpdate builder.
func (srmu *SysRoleMenuUpdate) Where(ps ...predicate.SysRoleMenu) *SysRoleMenuUpdate {
	srmu.mutation.Where(ps...)
	return srmu
}

// SetRoleID sets the "role_id" field.
func (srmu *SysRoleMenuUpdate) SetRoleID(i int64) *SysRoleMenuUpdate {
	srmu.mutation.ResetRoleID()
	srmu.mutation.SetRoleID(i)
	return srmu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableRoleID(i *int64) *SysRoleMenuUpdate {
	if i != nil {
		srmu.SetRoleID(*i)
	}
	return srmu
}

// AddRoleID adds i to the "role_id" field.
func (srmu *SysRoleMenuUpdate) AddRoleID(i int64) *SysRoleMenuUpdate {
	srmu.mutation.AddRoleID(i)
	return srmu
}

// SetMenuID sets the "menu_id" field.
func (srmu *SysRoleMenuUpdate) SetMenuID(i int64) *SysRoleMenuUpdate {
	srmu.mutation.ResetMenuID()
	srmu.mutation.SetMenuID(i)
	return srmu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableMenuID(i *int64) *SysRoleMenuUpdate {
	if i != nil {
		srmu.SetMenuID(*i)
	}
	return srmu
}

// AddMenuID adds i to the "menu_id" field.
func (srmu *SysRoleMenuUpdate) AddMenuID(i int64) *SysRoleMenuUpdate {
	srmu.mutation.AddMenuID(i)
	return srmu
}

// SetCreateTime sets the "create_time" field.
func (srmu *SysRoleMenuUpdate) SetCreateTime(i int32) *SysRoleMenuUpdate {
	srmu.mutation.ResetCreateTime()
	srmu.mutation.SetCreateTime(i)
	return srmu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableCreateTime(i *int32) *SysRoleMenuUpdate {
	if i != nil {
		srmu.SetCreateTime(*i)
	}
	return srmu
}

// AddCreateTime adds i to the "create_time" field.
func (srmu *SysRoleMenuUpdate) AddCreateTime(i int32) *SysRoleMenuUpdate {
	srmu.mutation.AddCreateTime(i)
	return srmu
}

// SetUpdateTime sets the "update_time" field.
func (srmu *SysRoleMenuUpdate) SetUpdateTime(i int32) *SysRoleMenuUpdate {
	srmu.mutation.ResetUpdateTime()
	srmu.mutation.SetUpdateTime(i)
	return srmu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableUpdateTime(i *int32) *SysRoleMenuUpdate {
	if i != nil {
		srmu.SetUpdateTime(*i)
	}
	return srmu
}

// AddUpdateTime adds i to the "update_time" field.
func (srmu *SysRoleMenuUpdate) AddUpdateTime(i int32) *SysRoleMenuUpdate {
	srmu.mutation.AddUpdateTime(i)
	return srmu
}

// Mutation returns the SysRoleMenuMutation object of the builder.
func (srmu *SysRoleMenuUpdate) Mutation() *SysRoleMenuMutation {
	return srmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (srmu *SysRoleMenuUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, srmu.sqlSave, srmu.mutation, srmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srmu *SysRoleMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := srmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (srmu *SysRoleMenuUpdate) Exec(ctx context.Context) error {
	_, err := srmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srmu *SysRoleMenuUpdate) ExecX(ctx context.Context) {
	if err := srmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (srmu *SysRoleMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sysrolemenu.Table, sysrolemenu.Columns, sqlgraph.NewFieldSpec(sysrolemenu.FieldID, field.TypeInt64))
	if ps := srmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srmu.mutation.RoleID(); ok {
		_spec.SetField(sysrolemenu.FieldRoleID, field.TypeInt64, value)
	}
	if value, ok := srmu.mutation.AddedRoleID(); ok {
		_spec.AddField(sysrolemenu.FieldRoleID, field.TypeInt64, value)
	}
	if value, ok := srmu.mutation.MenuID(); ok {
		_spec.SetField(sysrolemenu.FieldMenuID, field.TypeInt64, value)
	}
	if value, ok := srmu.mutation.AddedMenuID(); ok {
		_spec.AddField(sysrolemenu.FieldMenuID, field.TypeInt64, value)
	}
	if value, ok := srmu.mutation.CreateTime(); ok {
		_spec.SetField(sysrolemenu.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := srmu.mutation.AddedCreateTime(); ok {
		_spec.AddField(sysrolemenu.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := srmu.mutation.UpdateTime(); ok {
		_spec.SetField(sysrolemenu.FieldUpdateTime, field.TypeInt32, value)
	}
	if value, ok := srmu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(sysrolemenu.FieldUpdateTime, field.TypeInt32, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, srmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	srmu.mutation.done = true
	return n, nil
}

// SysRoleMenuUpdateOne is the builder for updating a single SysRoleMenu entity.
type SysRoleMenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysRoleMenuMutation
}

// SetRoleID sets the "role_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetRoleID(i int64) *SysRoleMenuUpdateOne {
	srmuo.mutation.ResetRoleID()
	srmuo.mutation.SetRoleID(i)
	return srmuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableRoleID(i *int64) *SysRoleMenuUpdateOne {
	if i != nil {
		srmuo.SetRoleID(*i)
	}
	return srmuo
}

// AddRoleID adds i to the "role_id" field.
func (srmuo *SysRoleMenuUpdateOne) AddRoleID(i int64) *SysRoleMenuUpdateOne {
	srmuo.mutation.AddRoleID(i)
	return srmuo
}

// SetMenuID sets the "menu_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetMenuID(i int64) *SysRoleMenuUpdateOne {
	srmuo.mutation.ResetMenuID()
	srmuo.mutation.SetMenuID(i)
	return srmuo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableMenuID(i *int64) *SysRoleMenuUpdateOne {
	if i != nil {
		srmuo.SetMenuID(*i)
	}
	return srmuo
}

// AddMenuID adds i to the "menu_id" field.
func (srmuo *SysRoleMenuUpdateOne) AddMenuID(i int64) *SysRoleMenuUpdateOne {
	srmuo.mutation.AddMenuID(i)
	return srmuo
}

// SetCreateTime sets the "create_time" field.
func (srmuo *SysRoleMenuUpdateOne) SetCreateTime(i int32) *SysRoleMenuUpdateOne {
	srmuo.mutation.ResetCreateTime()
	srmuo.mutation.SetCreateTime(i)
	return srmuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableCreateTime(i *int32) *SysRoleMenuUpdateOne {
	if i != nil {
		srmuo.SetCreateTime(*i)
	}
	return srmuo
}

// AddCreateTime adds i to the "create_time" field.
func (srmuo *SysRoleMenuUpdateOne) AddCreateTime(i int32) *SysRoleMenuUpdateOne {
	srmuo.mutation.AddCreateTime(i)
	return srmuo
}

// SetUpdateTime sets the "update_time" field.
func (srmuo *SysRoleMenuUpdateOne) SetUpdateTime(i int32) *SysRoleMenuUpdateOne {
	srmuo.mutation.ResetUpdateTime()
	srmuo.mutation.SetUpdateTime(i)
	return srmuo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableUpdateTime(i *int32) *SysRoleMenuUpdateOne {
	if i != nil {
		srmuo.SetUpdateTime(*i)
	}
	return srmuo
}

// AddUpdateTime adds i to the "update_time" field.
func (srmuo *SysRoleMenuUpdateOne) AddUpdateTime(i int32) *SysRoleMenuUpdateOne {
	srmuo.mutation.AddUpdateTime(i)
	return srmuo
}

// Mutation returns the SysRoleMenuMutation object of the builder.
func (srmuo *SysRoleMenuUpdateOne) Mutation() *SysRoleMenuMutation {
	return srmuo.mutation
}

// Where appends a list predicates to the SysRoleMenuUpdate builder.
func (srmuo *SysRoleMenuUpdateOne) Where(ps ...predicate.SysRoleMenu) *SysRoleMenuUpdateOne {
	srmuo.mutation.Where(ps...)
	return srmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (srmuo *SysRoleMenuUpdateOne) Select(field string, fields ...string) *SysRoleMenuUpdateOne {
	srmuo.fields = append([]string{field}, fields...)
	return srmuo
}

// Save executes the query and returns the updated SysRoleMenu entity.
func (srmuo *SysRoleMenuUpdateOne) Save(ctx context.Context) (*SysRoleMenu, error) {
	return withHooks(ctx, srmuo.sqlSave, srmuo.mutation, srmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srmuo *SysRoleMenuUpdateOne) SaveX(ctx context.Context) *SysRoleMenu {
	node, err := srmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (srmuo *SysRoleMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := srmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srmuo *SysRoleMenuUpdateOne) ExecX(ctx context.Context) {
	if err := srmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (srmuo *SysRoleMenuUpdateOne) sqlSave(ctx context.Context) (_node *SysRoleMenu, err error) {
	_spec := sqlgraph.NewUpdateSpec(sysrolemenu.Table, sysrolemenu.Columns, sqlgraph.NewFieldSpec(sysrolemenu.FieldID, field.TypeInt64))
	id, ok := srmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysRoleMenu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := srmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrolemenu.FieldID)
		for _, f := range fields {
			if !sysrolemenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysrolemenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := srmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srmuo.mutation.RoleID(); ok {
		_spec.SetField(sysrolemenu.FieldRoleID, field.TypeInt64, value)
	}
	if value, ok := srmuo.mutation.AddedRoleID(); ok {
		_spec.AddField(sysrolemenu.FieldRoleID, field.TypeInt64, value)
	}
	if value, ok := srmuo.mutation.MenuID(); ok {
		_spec.SetField(sysrolemenu.FieldMenuID, field.TypeInt64, value)
	}
	if value, ok := srmuo.mutation.AddedMenuID(); ok {
		_spec.AddField(sysrolemenu.FieldMenuID, field.TypeInt64, value)
	}
	if value, ok := srmuo.mutation.CreateTime(); ok {
		_spec.SetField(sysrolemenu.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := srmuo.mutation.AddedCreateTime(); ok {
		_spec.AddField(sysrolemenu.FieldCreateTime, field.TypeInt32, value)
	}
	if value, ok := srmuo.mutation.UpdateTime(); ok {
		_spec.SetField(sysrolemenu.FieldUpdateTime, field.TypeInt32, value)
	}
	if value, ok := srmuo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(sysrolemenu.FieldUpdateTime, field.TypeInt32, value)
	}
	_node = &SysRoleMenu{config: srmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, srmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	srmuo.mutation.done = true
	return _node, nil
}
