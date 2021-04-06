// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/filehash"
)

// FileHashCreate is the builder for creating a FileHash entity.
type FileHashCreate struct {
	config
	mutation *FileHashMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (fhc *FileHashCreate) SetHash(b []byte) *FileHashCreate {
	fhc.mutation.SetHash(b)
	return fhc
}

// SetFileid sets the "fileid" field.
func (fhc *FileHashCreate) SetFileid(s string) *FileHashCreate {
	fhc.mutation.SetFileid(s)
	return fhc
}

// Mutation returns the FileHashMutation object of the builder.
func (fhc *FileHashCreate) Mutation() *FileHashMutation {
	return fhc.mutation
}

// Save creates the FileHash in the database.
func (fhc *FileHashCreate) Save(ctx context.Context) (*FileHash, error) {
	var (
		err  error
		node *FileHash
	)
	if len(fhc.hooks) == 0 {
		if err = fhc.check(); err != nil {
			return nil, err
		}
		node, err = fhc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileHashMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fhc.check(); err != nil {
				return nil, err
			}
			fhc.mutation = mutation
			node, err = fhc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fhc.hooks) - 1; i >= 0; i-- {
			mut = fhc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fhc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fhc *FileHashCreate) SaveX(ctx context.Context) *FileHash {
	v, err := fhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (fhc *FileHashCreate) check() error {
	if _, ok := fhc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New("entgen: missing required field \"hash\"")}
	}
	if _, ok := fhc.mutation.Fileid(); !ok {
		return &ValidationError{Name: "fileid", err: errors.New("entgen: missing required field \"fileid\"")}
	}
	return nil
}

func (fhc *FileHashCreate) sqlSave(ctx context.Context) (*FileHash, error) {
	_node, _spec := fhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fhc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fhc *FileHashCreate) createSpec() (*FileHash, *sqlgraph.CreateSpec) {
	var (
		_node = &FileHash{config: fhc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: filehash.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filehash.FieldID,
			},
		}
	)
	if value, ok := fhc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: filehash.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := fhc.mutation.Fileid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filehash.FieldFileid,
		})
		_node.Fileid = value
	}
	return _node, _spec
}

// FileHashCreateBulk is the builder for creating many FileHash entities in bulk.
type FileHashCreateBulk struct {
	config
	builders []*FileHashCreate
}

// Save creates the FileHash entities in the database.
func (fhcb *FileHashCreateBulk) Save(ctx context.Context) ([]*FileHash, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fhcb.builders))
	nodes := make([]*FileHash, len(fhcb.builders))
	mutators := make([]Mutator, len(fhcb.builders))
	for i := range fhcb.builders {
		func(i int, root context.Context) {
			builder := fhcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileHashMutation)
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
					_, err = mutators[i+1].Mutate(root, fhcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fhcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fhcb *FileHashCreateBulk) SaveX(ctx context.Context) []*FileHash {
	v, err := fhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
