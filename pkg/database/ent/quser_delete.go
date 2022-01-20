// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/YEVER/pkg/database/ent/predicate"
	"github.com/willie-lin/YEVER/pkg/database/ent/quser"
)

// QuserDelete is the builder for deleting a Quser entity.
type QuserDelete struct {
	config
	hooks    []Hook
	mutation *QuserMutation
}

// Where appends a list predicates to the QuserDelete builder.
func (qd *QuserDelete) Where(ps ...predicate.Quser) *QuserDelete {
	qd.mutation.Where(ps...)
	return qd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qd *QuserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qd.hooks) == 0 {
		affected, err = qd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qd.mutation = mutation
			affected, err = qd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qd.hooks) - 1; i >= 0; i-- {
			if qd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (qd *QuserDelete) ExecX(ctx context.Context) int {
	n, err := qd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qd *QuserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: quser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: quser.FieldID,
			},
		},
	}
	if ps := qd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, qd.driver, _spec)
}

// QuserDeleteOne is the builder for deleting a single Quser entity.
type QuserDeleteOne struct {
	qd *QuserDelete
}

// Exec executes the deletion query.
func (qdo *QuserDeleteOne) Exec(ctx context.Context) error {
	n, err := qdo.qd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{quser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qdo *QuserDeleteOne) ExecX(ctx context.Context) {
	qdo.qd.ExecX(ctx)
}