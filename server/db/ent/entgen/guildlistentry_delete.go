// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/guildlistentry"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// GuildListEntryDelete is the builder for deleting a GuildListEntry entity.
type GuildListEntryDelete struct {
	config
	hooks    []Hook
	mutation *GuildListEntryMutation
}

// Where adds a new predicate to the GuildListEntryDelete builder.
func (gled *GuildListEntryDelete) Where(ps ...predicate.GuildListEntry) *GuildListEntryDelete {
	gled.mutation.predicates = append(gled.mutation.predicates, ps...)
	return gled
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gled *GuildListEntryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gled.hooks) == 0 {
		affected, err = gled.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildListEntryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gled.mutation = mutation
			affected, err = gled.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gled.hooks) - 1; i >= 0; i-- {
			mut = gled.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gled.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gled *GuildListEntryDelete) ExecX(ctx context.Context) int {
	n, err := gled.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gled *GuildListEntryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: guildlistentry.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: guildlistentry.FieldID,
			},
		},
	}
	if ps := gled.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gled.driver, _spec)
}

// GuildListEntryDeleteOne is the builder for deleting a single GuildListEntry entity.
type GuildListEntryDeleteOne struct {
	gled *GuildListEntryDelete
}

// Exec executes the deletion query.
func (gledo *GuildListEntryDeleteOne) Exec(ctx context.Context) error {
	n, err := gledo.gled.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{guildlistentry.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gledo *GuildListEntryDeleteOne) ExecX(ctx context.Context) {
	gledo.gled.ExecX(ctx)
}