// Code generated by ent, DO NOT EDIT.

package ent

import (
	"back-end/merchant/internal/data/ent/predicate"
	"back-end/merchant/internal/data/ent/sysroleusers"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysRoleUsersDelete is the builder for deleting a SysRoleUsers entity.
type SysRoleUsersDelete struct {
	config
	hooks    []Hook
	mutation *SysRoleUsersMutation
}

// Where appends a list predicates to the SysRoleUsersDelete builder.
func (srud *SysRoleUsersDelete) Where(ps ...predicate.SysRoleUsers) *SysRoleUsersDelete {
	srud.mutation.Where(ps...)
	return srud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srud *SysRoleUsersDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, srud.sqlExec, srud.mutation, srud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (srud *SysRoleUsersDelete) ExecX(ctx context.Context) int {
	n, err := srud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srud *SysRoleUsersDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysroleusers.Table, sqlgraph.NewFieldSpec(sysroleusers.FieldID, field.TypeInt32))
	if ps := srud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, srud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	srud.mutation.done = true
	return affected, err
}

// SysRoleUsersDeleteOne is the builder for deleting a single SysRoleUsers entity.
type SysRoleUsersDeleteOne struct {
	srud *SysRoleUsersDelete
}

// Where appends a list predicates to the SysRoleUsersDelete builder.
func (srudo *SysRoleUsersDeleteOne) Where(ps ...predicate.SysRoleUsers) *SysRoleUsersDeleteOne {
	srudo.srud.mutation.Where(ps...)
	return srudo
}

// Exec executes the deletion query.
func (srudo *SysRoleUsersDeleteOne) Exec(ctx context.Context) error {
	n, err := srudo.srud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysroleusers.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srudo *SysRoleUsersDeleteOne) ExecX(ctx context.Context) {
	if err := srudo.Exec(ctx); err != nil {
		panic(err)
	}
}