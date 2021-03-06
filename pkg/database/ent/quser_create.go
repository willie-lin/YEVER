// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/YEVER/pkg/database/ent/quser"
)

// QuserCreate is the builder for creating a Quser entity.
type QuserCreate struct {
	config
	mutation *QuserMutation
	hooks    []Hook
}

// SetQq sets the "qq" field.
func (qc *QuserCreate) SetQq(s string) *QuserCreate {
	qc.mutation.SetQq(s)
	return qc
}

// SetPhone sets the "phone" field.
func (qc *QuserCreate) SetPhone(s string) *QuserCreate {
	qc.mutation.SetPhone(s)
	return qc
}

// SetID sets the "id" field.
func (qc *QuserCreate) SetID(s string) *QuserCreate {
	qc.mutation.SetID(s)
	return qc
}

// Mutation returns the QuserMutation object of the builder.
func (qc *QuserCreate) Mutation() *QuserMutation {
	return qc.mutation
}

// Save creates the Quser in the database.
func (qc *QuserCreate) Save(ctx context.Context) (*Quser, error) {
	var (
		err  error
		node *Quser
	)
	if len(qc.hooks) == 0 {
		if err = qc.check(); err != nil {
			return nil, err
		}
		node, err = qc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qc.check(); err != nil {
				return nil, err
			}
			qc.mutation = mutation
			if node, err = qc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(qc.hooks) - 1; i >= 0; i-- {
			if qc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QuserCreate) SaveX(ctx context.Context) *Quser {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qc *QuserCreate) Exec(ctx context.Context) error {
	_, err := qc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qc *QuserCreate) ExecX(ctx context.Context) {
	if err := qc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QuserCreate) check() error {
	if _, ok := qc.mutation.Qq(); !ok {
		return &ValidationError{Name: "qq", err: errors.New(`ent: missing required field "Quser.qq"`)}
	}
	if _, ok := qc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Quser.phone"`)}
	}
	return nil
}

func (qc *QuserCreate) sqlSave(ctx context.Context) (*Quser, error) {
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Quser.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (qc *QuserCreate) createSpec() (*Quser, *sqlgraph.CreateSpec) {
	var (
		_node = &Quser{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: quser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: quser.FieldID,
			},
		}
	)
	if id, ok := qc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := qc.mutation.Qq(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: quser.FieldQq,
		})
		_node.Qq = value
	}
	if value, ok := qc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: quser.FieldPhone,
		})
		_node.Phone = value
	}
	return _node, _spec
}

// QuserCreateBulk is the builder for creating many Quser entities in bulk.
type QuserCreateBulk struct {
	config
	builders []*QuserCreate
}

// Save creates the Quser entities in the database.
func (qcb *QuserCreateBulk) Save(ctx context.Context) ([]*Quser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Quser, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QuserCreateBulk) SaveX(ctx context.Context) []*Quser {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcb *QuserCreateBulk) Exec(ctx context.Context) error {
	_, err := qcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcb *QuserCreateBulk) ExecX(ctx context.Context) {
	if err := qcb.Exec(ctx); err != nil {
		panic(err)
	}
}
