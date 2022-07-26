// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent/profitdetail"
)

// ProfitDetailDelete is the builder for deleting a ProfitDetail entity.
type ProfitDetailDelete struct {
	config
	hooks    []Hook
	mutation *ProfitDetailMutation
}

// Where appends a list predicates to the ProfitDetailDelete builder.
func (pdd *ProfitDetailDelete) Where(ps ...predicate.ProfitDetail) *ProfitDetailDelete {
	pdd.mutation.Where(ps...)
	return pdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pdd *ProfitDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pdd.hooks) == 0 {
		affected, err = pdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfitDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pdd.mutation = mutation
			affected, err = pdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pdd.hooks) - 1; i >= 0; i-- {
			if pdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdd *ProfitDetailDelete) ExecX(ctx context.Context) int {
	n, err := pdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pdd *ProfitDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: profitdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profitdetail.FieldID,
			},
		},
	}
	if ps := pdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pdd.driver, _spec)
}

// ProfitDetailDeleteOne is the builder for deleting a single ProfitDetail entity.
type ProfitDetailDeleteOne struct {
	pdd *ProfitDetailDelete
}

// Exec executes the deletion query.
func (pddo *ProfitDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := pddo.pdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{profitdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pddo *ProfitDetailDeleteOne) ExecX(ctx context.Context) {
	pddo.pdd.ExecX(ctx)
}
