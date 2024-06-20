// Code generated by ent, DO NOT EDIT.

package ent

import (
	"back-end/merchant/internal/data/ent/predicate"
	"back-end/merchant/internal/data/ent/syspermissions"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysPermissionsDelete is the builder for deleting a SysPermissions entity.
type SysPermissionsDelete struct {
	config
	hooks    []Hook
	mutation *SysPermissionsMutation
}

// Where appends a list predicates to the SysPermissionsDelete builder.
func (spd *SysPermissionsDelete) Where(ps ...predicate.SysPermissions) *SysPermissionsDelete {
	spd.mutation.Where(ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *SysPermissionsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, spd.sqlExec, spd.mutation, spd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *SysPermissionsDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *SysPermissionsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(syspermissions.Table, sqlgraph.NewFieldSpec(syspermissions.FieldID, field.TypeInt64))
	if ps := spd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	spd.mutation.done = true
	return affected, err
}

// SysPermissionsDeleteOne is the builder for deleting a single SysPermissions entity.
type SysPermissionsDeleteOne struct {
	spd *SysPermissionsDelete
}

// Where appends a list predicates to the SysPermissionsDelete builder.
func (spdo *SysPermissionsDeleteOne) Where(ps ...predicate.SysPermissions) *SysPermissionsDeleteOne {
	spdo.spd.mutation.Where(ps...)
	return spdo
}

// Exec executes the deletion query.
func (spdo *SysPermissionsDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{syspermissions.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *SysPermissionsDeleteOne) ExecX(ctx context.Context) {
	if err := spdo.Exec(ctx); err != nil {
		panic(err)
	}
}