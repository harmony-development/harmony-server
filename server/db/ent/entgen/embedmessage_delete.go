// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/embedmessage"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// EmbedMessageDelete is the builder for deleting a EmbedMessage entity.
type EmbedMessageDelete struct {
	config
	hooks    []Hook
	mutation *EmbedMessageMutation
}

// Where adds a new predicate to the EmbedMessageDelete builder.
func (emd *EmbedMessageDelete) Where(ps ...predicate.EmbedMessage) *EmbedMessageDelete {
	emd.mutation.predicates = append(emd.mutation.predicates, ps...)
	return emd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (emd *EmbedMessageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(emd.hooks) == 0 {
		affected, err = emd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmbedMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			emd.mutation = mutation
			affected, err = emd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(emd.hooks) - 1; i >= 0; i-- {
			mut = emd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, emd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (emd *EmbedMessageDelete) ExecX(ctx context.Context) int {
	n, err := emd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (emd *EmbedMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: embedmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: embedmessage.FieldID,
			},
		},
	}
	if ps := emd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, emd.driver, _spec)
}

// EmbedMessageDeleteOne is the builder for deleting a single EmbedMessage entity.
type EmbedMessageDeleteOne struct {
	emd *EmbedMessageDelete
}

// Exec executes the deletion query.
func (emdo *EmbedMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := emdo.emd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{embedmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (emdo *EmbedMessageDeleteOne) ExecX(ctx context.Context) {
	emdo.emd.ExecX(ctx)
}
