// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/ent/entgen/usermeta"
)

// UserMetaUpdate is the builder for updating UserMeta entities.
type UserMetaUpdate struct {
	config
	hooks    []Hook
	mutation *UserMetaMutation
}

// Where adds a new predicate for the UserMetaUpdate builder.
func (umu *UserMetaUpdate) Where(ps ...predicate.UserMeta) *UserMetaUpdate {
	umu.mutation.predicates = append(umu.mutation.predicates, ps...)
	return umu
}

// SetMeta sets the "meta" field.
func (umu *UserMetaUpdate) SetMeta(s string) *UserMetaUpdate {
	umu.mutation.SetMeta(s)
	return umu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (umu *UserMetaUpdate) SetUserID(id uint64) *UserMetaUpdate {
	umu.mutation.SetUserID(id)
	return umu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (umu *UserMetaUpdate) SetNillableUserID(id *uint64) *UserMetaUpdate {
	if id != nil {
		umu = umu.SetUserID(*id)
	}
	return umu
}

// SetUser sets the "user" edge to the User entity.
func (umu *UserMetaUpdate) SetUser(u *User) *UserMetaUpdate {
	return umu.SetUserID(u.ID)
}

// Mutation returns the UserMetaMutation object of the builder.
func (umu *UserMetaUpdate) Mutation() *UserMetaMutation {
	return umu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (umu *UserMetaUpdate) ClearUser() *UserMetaUpdate {
	umu.mutation.ClearUser()
	return umu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (umu *UserMetaUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(umu.hooks) == 0 {
		affected, err = umu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMetaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			umu.mutation = mutation
			affected, err = umu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(umu.hooks) - 1; i >= 0; i-- {
			mut = umu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, umu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (umu *UserMetaUpdate) SaveX(ctx context.Context) int {
	affected, err := umu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (umu *UserMetaUpdate) Exec(ctx context.Context) error {
	_, err := umu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umu *UserMetaUpdate) ExecX(ctx context.Context) {
	if err := umu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (umu *UserMetaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usermeta.Table,
			Columns: usermeta.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: usermeta.FieldID,
			},
		},
	}
	if ps := umu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := umu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usermeta.FieldMeta,
		})
	}
	if umu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermeta.UserTable,
			Columns: []string{usermeta.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := umu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermeta.UserTable,
			Columns: []string{usermeta.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, umu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usermeta.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserMetaUpdateOne is the builder for updating a single UserMeta entity.
type UserMetaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMetaMutation
}

// SetMeta sets the "meta" field.
func (umuo *UserMetaUpdateOne) SetMeta(s string) *UserMetaUpdateOne {
	umuo.mutation.SetMeta(s)
	return umuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (umuo *UserMetaUpdateOne) SetUserID(id uint64) *UserMetaUpdateOne {
	umuo.mutation.SetUserID(id)
	return umuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (umuo *UserMetaUpdateOne) SetNillableUserID(id *uint64) *UserMetaUpdateOne {
	if id != nil {
		umuo = umuo.SetUserID(*id)
	}
	return umuo
}

// SetUser sets the "user" edge to the User entity.
func (umuo *UserMetaUpdateOne) SetUser(u *User) *UserMetaUpdateOne {
	return umuo.SetUserID(u.ID)
}

// Mutation returns the UserMetaMutation object of the builder.
func (umuo *UserMetaUpdateOne) Mutation() *UserMetaMutation {
	return umuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (umuo *UserMetaUpdateOne) ClearUser() *UserMetaUpdateOne {
	umuo.mutation.ClearUser()
	return umuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (umuo *UserMetaUpdateOne) Select(field string, fields ...string) *UserMetaUpdateOne {
	umuo.fields = append([]string{field}, fields...)
	return umuo
}

// Save executes the query and returns the updated UserMeta entity.
func (umuo *UserMetaUpdateOne) Save(ctx context.Context) (*UserMeta, error) {
	var (
		err  error
		node *UserMeta
	)
	if len(umuo.hooks) == 0 {
		node, err = umuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMetaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			umuo.mutation = mutation
			node, err = umuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(umuo.hooks) - 1; i >= 0; i-- {
			mut = umuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, umuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (umuo *UserMetaUpdateOne) SaveX(ctx context.Context) *UserMeta {
	node, err := umuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (umuo *UserMetaUpdateOne) Exec(ctx context.Context) error {
	_, err := umuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umuo *UserMetaUpdateOne) ExecX(ctx context.Context) {
	if err := umuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (umuo *UserMetaUpdateOne) sqlSave(ctx context.Context) (_node *UserMeta, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usermeta.Table,
			Columns: usermeta.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: usermeta.FieldID,
			},
		},
	}
	id, ok := umuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserMeta.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := umuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usermeta.FieldID)
		for _, f := range fields {
			if !usermeta.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != usermeta.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := umuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := umuo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usermeta.FieldMeta,
		})
	}
	if umuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermeta.UserTable,
			Columns: []string{usermeta.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := umuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermeta.UserTable,
			Columns: []string{usermeta.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserMeta{config: umuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, umuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usermeta.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
