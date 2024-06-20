// Code generated by ent, DO NOT EDIT.

package ent

import (
	"back-end/merchant/internal/data/ent/merchantlist"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MerchantListCreate is the builder for creating a MerchantList entity.
type MerchantListCreate struct {
	config
	mutation *MerchantListMutation
	hooks    []Hook
}

// SetGUID sets the "guid" field.
func (mlc *MerchantListCreate) SetGUID(s string) *MerchantListCreate {
	mlc.mutation.SetGUID(s)
	return mlc
}

// SetName sets the "name" field.
func (mlc *MerchantListCreate) SetName(s string) *MerchantListCreate {
	mlc.mutation.SetName(s)
	return mlc
}

// SetHost sets the "host" field.
func (mlc *MerchantListCreate) SetHost(s string) *MerchantListCreate {
	mlc.mutation.SetHost(s)
	return mlc
}

// SetManageHost sets the "manage_host" field.
func (mlc *MerchantListCreate) SetManageHost(s string) *MerchantListCreate {
	mlc.mutation.SetManageHost(s)
	return mlc
}

// SetStatus sets the "status" field.
func (mlc *MerchantListCreate) SetStatus(i int16) *MerchantListCreate {
	mlc.mutation.SetStatus(i)
	return mlc
}

// SetCreateTime sets the "create_time" field.
func (mlc *MerchantListCreate) SetCreateTime(i int32) *MerchantListCreate {
	mlc.mutation.SetCreateTime(i)
	return mlc
}

// SetUpdateTime sets the "update_time" field.
func (mlc *MerchantListCreate) SetUpdateTime(i int32) *MerchantListCreate {
	mlc.mutation.SetUpdateTime(i)
	return mlc
}

// SetID sets the "id" field.
func (mlc *MerchantListCreate) SetID(i int32) *MerchantListCreate {
	mlc.mutation.SetID(i)
	return mlc
}

// Mutation returns the MerchantListMutation object of the builder.
func (mlc *MerchantListCreate) Mutation() *MerchantListMutation {
	return mlc.mutation
}

// Save creates the MerchantList in the database.
func (mlc *MerchantListCreate) Save(ctx context.Context) (*MerchantList, error) {
	return withHooks(ctx, mlc.sqlSave, mlc.mutation, mlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mlc *MerchantListCreate) SaveX(ctx context.Context) *MerchantList {
	v, err := mlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mlc *MerchantListCreate) Exec(ctx context.Context) error {
	_, err := mlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlc *MerchantListCreate) ExecX(ctx context.Context) {
	if err := mlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mlc *MerchantListCreate) check() error {
	if _, ok := mlc.mutation.GUID(); !ok {
		return &ValidationError{Name: "guid", err: errors.New(`ent: missing required field "MerchantList.guid"`)}
	}
	if _, ok := mlc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MerchantList.name"`)}
	}
	if _, ok := mlc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "MerchantList.host"`)}
	}
	if _, ok := mlc.mutation.ManageHost(); !ok {
		return &ValidationError{Name: "manage_host", err: errors.New(`ent: missing required field "MerchantList.manage_host"`)}
	}
	if _, ok := mlc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "MerchantList.status"`)}
	}
	if _, ok := mlc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "MerchantList.create_time"`)}
	}
	if _, ok := mlc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "MerchantList.update_time"`)}
	}
	return nil
}

func (mlc *MerchantListCreate) sqlSave(ctx context.Context) (*MerchantList, error) {
	if err := mlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	mlc.mutation.id = &_node.ID
	mlc.mutation.done = true
	return _node, nil
}

func (mlc *MerchantListCreate) createSpec() (*MerchantList, *sqlgraph.CreateSpec) {
	var (
		_node = &MerchantList{config: mlc.config}
		_spec = sqlgraph.NewCreateSpec(merchantlist.Table, sqlgraph.NewFieldSpec(merchantlist.FieldID, field.TypeInt32))
	)
	if id, ok := mlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mlc.mutation.GUID(); ok {
		_spec.SetField(merchantlist.FieldGUID, field.TypeString, value)
		_node.GUID = value
	}
	if value, ok := mlc.mutation.Name(); ok {
		_spec.SetField(merchantlist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mlc.mutation.Host(); ok {
		_spec.SetField(merchantlist.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := mlc.mutation.ManageHost(); ok {
		_spec.SetField(merchantlist.FieldManageHost, field.TypeString, value)
		_node.ManageHost = value
	}
	if value, ok := mlc.mutation.Status(); ok {
		_spec.SetField(merchantlist.FieldStatus, field.TypeInt16, value)
		_node.Status = value
	}
	if value, ok := mlc.mutation.CreateTime(); ok {
		_spec.SetField(merchantlist.FieldCreateTime, field.TypeInt32, value)
		_node.CreateTime = value
	}
	if value, ok := mlc.mutation.UpdateTime(); ok {
		_spec.SetField(merchantlist.FieldUpdateTime, field.TypeInt32, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// MerchantListCreateBulk is the builder for creating many MerchantList entities in bulk.
type MerchantListCreateBulk struct {
	config
	err      error
	builders []*MerchantListCreate
}

// Save creates the MerchantList entities in the database.
func (mlcb *MerchantListCreateBulk) Save(ctx context.Context) ([]*MerchantList, error) {
	if mlcb.err != nil {
		return nil, mlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mlcb.builders))
	nodes := make([]*MerchantList, len(mlcb.builders))
	mutators := make([]Mutator, len(mlcb.builders))
	for i := range mlcb.builders {
		func(i int, root context.Context) {
			builder := mlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MerchantListMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mlcb *MerchantListCreateBulk) SaveX(ctx context.Context) []*MerchantList {
	v, err := mlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mlcb *MerchantListCreateBulk) Exec(ctx context.Context) error {
	_, err := mlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlcb *MerchantListCreateBulk) ExecX(ctx context.Context) {
	if err := mlcb.Exec(ctx); err != nil {
		panic(err)
	}
}