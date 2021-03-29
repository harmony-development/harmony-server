// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/localuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/session"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// LocalUserUpdate is the builder for updating LocalUser entities.
type LocalUserUpdate struct {
	config
	hooks    []Hook
	mutation *LocalUserMutation
}

// Where adds a new predicate for the LocalUserUpdate builder.
func (luu *LocalUserUpdate) Where(ps ...predicate.LocalUser) *LocalUserUpdate {
	luu.mutation.predicates = append(luu.mutation.predicates, ps...)
	return luu
}

// SetEmail sets the "email" field.
func (luu *LocalUserUpdate) SetEmail(s string) *LocalUserUpdate {
	luu.mutation.SetEmail(s)
	return luu
}

// SetPassword sets the "password" field.
func (luu *LocalUserUpdate) SetPassword(b []byte) *LocalUserUpdate {
	luu.mutation.SetPassword(b)
	return luu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (luu *LocalUserUpdate) SetUserID(id uint64) *LocalUserUpdate {
	luu.mutation.SetUserID(id)
	return luu
}

// SetUser sets the "user" edge to the User entity.
func (luu *LocalUserUpdate) SetUser(u *User) *LocalUserUpdate {
	return luu.SetUserID(u.ID)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (luu *LocalUserUpdate) AddSessionIDs(ids ...int) *LocalUserUpdate {
	luu.mutation.AddSessionIDs(ids...)
	return luu
}

// AddSessions adds the "sessions" edges to the Session entity.
func (luu *LocalUserUpdate) AddSessions(s ...*Session) *LocalUserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return luu.AddSessionIDs(ids...)
}

// Mutation returns the LocalUserMutation object of the builder.
func (luu *LocalUserUpdate) Mutation() *LocalUserMutation {
	return luu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (luu *LocalUserUpdate) ClearUser() *LocalUserUpdate {
	luu.mutation.ClearUser()
	return luu
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (luu *LocalUserUpdate) ClearSessions() *LocalUserUpdate {
	luu.mutation.ClearSessions()
	return luu
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (luu *LocalUserUpdate) RemoveSessionIDs(ids ...int) *LocalUserUpdate {
	luu.mutation.RemoveSessionIDs(ids...)
	return luu
}

// RemoveSessions removes "sessions" edges to Session entities.
func (luu *LocalUserUpdate) RemoveSessions(s ...*Session) *LocalUserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return luu.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (luu *LocalUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(luu.hooks) == 0 {
		if err = luu.check(); err != nil {
			return 0, err
		}
		affected, err = luu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LocalUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luu.check(); err != nil {
				return 0, err
			}
			luu.mutation = mutation
			affected, err = luu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(luu.hooks) - 1; i >= 0; i-- {
			mut = luu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (luu *LocalUserUpdate) SaveX(ctx context.Context) int {
	affected, err := luu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (luu *LocalUserUpdate) Exec(ctx context.Context) error {
	_, err := luu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luu *LocalUserUpdate) ExecX(ctx context.Context) {
	if err := luu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luu *LocalUserUpdate) check() error {
	if v, ok := luu.mutation.Email(); ok {
		if err := localuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("entgen: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := luu.mutation.UserID(); luu.mutation.UserCleared() && !ok {
		return errors.New("entgen: clearing a required unique edge \"user\"")
	}
	return nil
}

func (luu *LocalUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   localuser.Table,
			Columns: localuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: localuser.FieldID,
			},
		},
	}
	if ps := luu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: localuser.FieldEmail,
		})
	}
	if value, ok := luu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: localuser.FieldPassword,
		})
	}
	if luu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   localuser.UserTable,
			Columns: []string{localuser.UserColumn},
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
	if nodes := luu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   localuser.UserTable,
			Columns: []string{localuser.UserColumn},
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
	if luu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   localuser.SessionsTable,
			Columns: []string{localuser.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luu.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !luu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   localuser.SessionsTable,
			Columns: []string{localuser.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luu.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   localuser.SessionsTable,
			Columns: []string{localuser.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, luu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{localuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LocalUserUpdateOne is the builder for updating a single LocalUser entity.
type LocalUserUpdateOne struct {
	config
	hooks    []Hook
	mutation *LocalUserMutation
}

// SetEmail sets the "email" field.
func (luuo *LocalUserUpdateOne) SetEmail(s string) *LocalUserUpdateOne {
	luuo.mutation.SetEmail(s)
	return luuo
}

// SetPassword sets the "password" field.
func (luuo *LocalUserUpdateOne) SetPassword(b []byte) *LocalUserUpdateOne {
	luuo.mutation.SetPassword(b)
	return luuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (luuo *LocalUserUpdateOne) SetUserID(id uint64) *LocalUserUpdateOne {
	luuo.mutation.SetUserID(id)
	return luuo
}

// SetUser sets the "user" edge to the User entity.
func (luuo *LocalUserUpdateOne) SetUser(u *User) *LocalUserUpdateOne {
	return luuo.SetUserID(u.ID)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (luuo *LocalUserUpdateOne) AddSessionIDs(ids ...int) *LocalUserUpdateOne {
	luuo.mutation.AddSessionIDs(ids...)
	return luuo
}

// AddSessions adds the "sessions" edges to the Session entity.
func (luuo *LocalUserUpdateOne) AddSessions(s ...*Session) *LocalUserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return luuo.AddSessionIDs(ids...)
}

// Mutation returns the LocalUserMutation object of the builder.
func (luuo *LocalUserUpdateOne) Mutation() *LocalUserMutation {
	return luuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (luuo *LocalUserUpdateOne) ClearUser() *LocalUserUpdateOne {
	luuo.mutation.ClearUser()
	return luuo
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (luuo *LocalUserUpdateOne) ClearSessions() *LocalUserUpdateOne {
	luuo.mutation.ClearSessions()
	return luuo
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (luuo *LocalUserUpdateOne) RemoveSessionIDs(ids ...int) *LocalUserUpdateOne {
	luuo.mutation.RemoveSessionIDs(ids...)
	return luuo
}

// RemoveSessions removes "sessions" edges to Session entities.
func (luuo *LocalUserUpdateOne) RemoveSessions(s ...*Session) *LocalUserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return luuo.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the updated LocalUser entity.
func (luuo *LocalUserUpdateOne) Save(ctx context.Context) (*LocalUser, error) {
	var (
		err  error
		node *LocalUser
	)
	if len(luuo.hooks) == 0 {
		if err = luuo.check(); err != nil {
			return nil, err
		}
		node, err = luuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LocalUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luuo.check(); err != nil {
				return nil, err
			}
			luuo.mutation = mutation
			node, err = luuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luuo.hooks) - 1; i >= 0; i-- {
			mut = luuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luuo *LocalUserUpdateOne) SaveX(ctx context.Context) *LocalUser {
	node, err := luuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luuo *LocalUserUpdateOne) Exec(ctx context.Context) error {
	_, err := luuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luuo *LocalUserUpdateOne) ExecX(ctx context.Context) {
	if err := luuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luuo *LocalUserUpdateOne) check() error {
	if v, ok := luuo.mutation.Email(); ok {
		if err := localuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("entgen: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := luuo.mutation.UserID(); luuo.mutation.UserCleared() && !ok {
		return errors.New("entgen: clearing a required unique edge \"user\"")
	}
	return nil
}

func (luuo *LocalUserUpdateOne) sqlSave(ctx context.Context) (_node *LocalUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   localuser.Table,
			Columns: localuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: localuser.FieldID,
			},
		},
	}
	id, ok := luuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LocalUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := luuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: localuser.FieldEmail,
		})
	}
	if value, ok := luuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: localuser.FieldPassword,
		})
	}
	if luuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   localuser.UserTable,
			Columns: []string{localuser.UserColumn},
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
	if nodes := luuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   localuser.UserTable,
			Columns: []string{localuser.UserColumn},
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
	if luuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   localuser.SessionsTable,
			Columns: []string{localuser.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luuo.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !luuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   localuser.SessionsTable,
			Columns: []string{localuser.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luuo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   localuser.SessionsTable,
			Columns: []string{localuser.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LocalUser{config: luuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{localuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
