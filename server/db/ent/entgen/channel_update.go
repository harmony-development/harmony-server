// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	v1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/permissionnode"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
)

// ChannelUpdate is the builder for updating Channel entities.
type ChannelUpdate struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where adds a new predicate for the ChannelUpdate builder.
func (cu *ChannelUpdate) Where(ps ...predicate.Channel) *ChannelUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ChannelUpdate) SetName(s string) *ChannelUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetKind sets the "kind" field.
func (cu *ChannelUpdate) SetKind(u uint64) *ChannelUpdate {
	cu.mutation.ResetKind()
	cu.mutation.SetKind(u)
	return cu
}

// AddKind adds u to the "kind" field.
func (cu *ChannelUpdate) AddKind(u uint64) *ChannelUpdate {
	cu.mutation.AddKind(u)
	return cu
}

// SetPosition sets the "position" field.
func (cu *ChannelUpdate) SetPosition(s string) *ChannelUpdate {
	cu.mutation.SetPosition(s)
	return cu
}

// SetMetadata sets the "metadata" field.
func (cu *ChannelUpdate) SetMetadata(v *v1.Metadata) *ChannelUpdate {
	cu.mutation.SetMetadata(v)
	return cu
}

// SetGuildID sets the "guild" edge to the Guild entity by ID.
func (cu *ChannelUpdate) SetGuildID(id uint64) *ChannelUpdate {
	cu.mutation.SetGuildID(id)
	return cu
}

// SetNillableGuildID sets the "guild" edge to the Guild entity by ID if the given value is not nil.
func (cu *ChannelUpdate) SetNillableGuildID(id *uint64) *ChannelUpdate {
	if id != nil {
		cu = cu.SetGuildID(*id)
	}
	return cu
}

// SetGuild sets the "guild" edge to the Guild entity.
func (cu *ChannelUpdate) SetGuild(g *Guild) *ChannelUpdate {
	return cu.SetGuildID(g.ID)
}

// AddMessageIDs adds the "message" edge to the Message entity by IDs.
func (cu *ChannelUpdate) AddMessageIDs(ids ...uint64) *ChannelUpdate {
	cu.mutation.AddMessageIDs(ids...)
	return cu
}

// AddMessage adds the "message" edges to the Message entity.
func (cu *ChannelUpdate) AddMessage(m ...*Message) *ChannelUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMessageIDs(ids...)
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (cu *ChannelUpdate) AddRoleIDs(ids ...uint64) *ChannelUpdate {
	cu.mutation.AddRoleIDs(ids...)
	return cu
}

// AddRole adds the "role" edges to the Role entity.
func (cu *ChannelUpdate) AddRole(r ...*Role) *ChannelUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddRoleIDs(ids...)
}

// AddPermissionNodeIDs adds the "permission_node" edge to the PermissionNode entity by IDs.
func (cu *ChannelUpdate) AddPermissionNodeIDs(ids ...int) *ChannelUpdate {
	cu.mutation.AddPermissionNodeIDs(ids...)
	return cu
}

// AddPermissionNode adds the "permission_node" edges to the PermissionNode entity.
func (cu *ChannelUpdate) AddPermissionNode(p ...*PermissionNode) *ChannelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPermissionNodeIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cu *ChannelUpdate) Mutation() *ChannelMutation {
	return cu.mutation
}

// ClearGuild clears the "guild" edge to the Guild entity.
func (cu *ChannelUpdate) ClearGuild() *ChannelUpdate {
	cu.mutation.ClearGuild()
	return cu
}

// ClearMessage clears all "message" edges to the Message entity.
func (cu *ChannelUpdate) ClearMessage() *ChannelUpdate {
	cu.mutation.ClearMessage()
	return cu
}

// RemoveMessageIDs removes the "message" edge to Message entities by IDs.
func (cu *ChannelUpdate) RemoveMessageIDs(ids ...uint64) *ChannelUpdate {
	cu.mutation.RemoveMessageIDs(ids...)
	return cu
}

// RemoveMessage removes "message" edges to Message entities.
func (cu *ChannelUpdate) RemoveMessage(m ...*Message) *ChannelUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMessageIDs(ids...)
}

// ClearRole clears all "role" edges to the Role entity.
func (cu *ChannelUpdate) ClearRole() *ChannelUpdate {
	cu.mutation.ClearRole()
	return cu
}

// RemoveRoleIDs removes the "role" edge to Role entities by IDs.
func (cu *ChannelUpdate) RemoveRoleIDs(ids ...uint64) *ChannelUpdate {
	cu.mutation.RemoveRoleIDs(ids...)
	return cu
}

// RemoveRole removes "role" edges to Role entities.
func (cu *ChannelUpdate) RemoveRole(r ...*Role) *ChannelUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveRoleIDs(ids...)
}

// ClearPermissionNode clears all "permission_node" edges to the PermissionNode entity.
func (cu *ChannelUpdate) ClearPermissionNode() *ChannelUpdate {
	cu.mutation.ClearPermissionNode()
	return cu
}

// RemovePermissionNodeIDs removes the "permission_node" edge to PermissionNode entities by IDs.
func (cu *ChannelUpdate) RemovePermissionNodeIDs(ids ...int) *ChannelUpdate {
	cu.mutation.RemovePermissionNodeIDs(ids...)
	return cu
}

// RemovePermissionNode removes "permission_node" edges to PermissionNode entities.
func (cu *ChannelUpdate) RemovePermissionNode(p ...*PermissionNode) *ChannelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePermissionNodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChannelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChannelUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChannelUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: channel.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldName,
		})
	}
	if value, ok := cu.mutation.Kind(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: channel.FieldKind,
		})
	}
	if value, ok := cu.mutation.AddedKind(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: channel.FieldKind,
		})
	}
	if value, ok := cu.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldPosition,
		})
	}
	if value, ok := cu.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: channel.FieldMetadata,
		})
	}
	if cu.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channel.GuildTable,
			Columns: []string{channel.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channel.GuildTable,
			Columns: []string{channel.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessageTable,
			Columns: []string{channel.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMessageIDs(); len(nodes) > 0 && !cu.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessageTable,
			Columns: []string{channel.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessageTable,
			Columns: []string{channel.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.RoleTable,
			Columns: []string{channel.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedRoleIDs(); len(nodes) > 0 && !cu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.RoleTable,
			Columns: []string{channel.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.RoleTable,
			Columns: []string{channel.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PermissionNodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.PermissionNodeTable,
			Columns: []string{channel.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPermissionNodeIDs(); len(nodes) > 0 && !cu.mutation.PermissionNodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.PermissionNodeTable,
			Columns: []string{channel.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PermissionNodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.PermissionNodeTable,
			Columns: []string{channel.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ChannelUpdateOne is the builder for updating a single Channel entity.
type ChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChannelMutation
}

// SetName sets the "name" field.
func (cuo *ChannelUpdateOne) SetName(s string) *ChannelUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetKind sets the "kind" field.
func (cuo *ChannelUpdateOne) SetKind(u uint64) *ChannelUpdateOne {
	cuo.mutation.ResetKind()
	cuo.mutation.SetKind(u)
	return cuo
}

// AddKind adds u to the "kind" field.
func (cuo *ChannelUpdateOne) AddKind(u uint64) *ChannelUpdateOne {
	cuo.mutation.AddKind(u)
	return cuo
}

// SetPosition sets the "position" field.
func (cuo *ChannelUpdateOne) SetPosition(s string) *ChannelUpdateOne {
	cuo.mutation.SetPosition(s)
	return cuo
}

// SetMetadata sets the "metadata" field.
func (cuo *ChannelUpdateOne) SetMetadata(v *v1.Metadata) *ChannelUpdateOne {
	cuo.mutation.SetMetadata(v)
	return cuo
}

// SetGuildID sets the "guild" edge to the Guild entity by ID.
func (cuo *ChannelUpdateOne) SetGuildID(id uint64) *ChannelUpdateOne {
	cuo.mutation.SetGuildID(id)
	return cuo
}

// SetNillableGuildID sets the "guild" edge to the Guild entity by ID if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableGuildID(id *uint64) *ChannelUpdateOne {
	if id != nil {
		cuo = cuo.SetGuildID(*id)
	}
	return cuo
}

// SetGuild sets the "guild" edge to the Guild entity.
func (cuo *ChannelUpdateOne) SetGuild(g *Guild) *ChannelUpdateOne {
	return cuo.SetGuildID(g.ID)
}

// AddMessageIDs adds the "message" edge to the Message entity by IDs.
func (cuo *ChannelUpdateOne) AddMessageIDs(ids ...uint64) *ChannelUpdateOne {
	cuo.mutation.AddMessageIDs(ids...)
	return cuo
}

// AddMessage adds the "message" edges to the Message entity.
func (cuo *ChannelUpdateOne) AddMessage(m ...*Message) *ChannelUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMessageIDs(ids...)
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (cuo *ChannelUpdateOne) AddRoleIDs(ids ...uint64) *ChannelUpdateOne {
	cuo.mutation.AddRoleIDs(ids...)
	return cuo
}

// AddRole adds the "role" edges to the Role entity.
func (cuo *ChannelUpdateOne) AddRole(r ...*Role) *ChannelUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddRoleIDs(ids...)
}

// AddPermissionNodeIDs adds the "permission_node" edge to the PermissionNode entity by IDs.
func (cuo *ChannelUpdateOne) AddPermissionNodeIDs(ids ...int) *ChannelUpdateOne {
	cuo.mutation.AddPermissionNodeIDs(ids...)
	return cuo
}

// AddPermissionNode adds the "permission_node" edges to the PermissionNode entity.
func (cuo *ChannelUpdateOne) AddPermissionNode(p ...*PermissionNode) *ChannelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPermissionNodeIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cuo *ChannelUpdateOne) Mutation() *ChannelMutation {
	return cuo.mutation
}

// ClearGuild clears the "guild" edge to the Guild entity.
func (cuo *ChannelUpdateOne) ClearGuild() *ChannelUpdateOne {
	cuo.mutation.ClearGuild()
	return cuo
}

// ClearMessage clears all "message" edges to the Message entity.
func (cuo *ChannelUpdateOne) ClearMessage() *ChannelUpdateOne {
	cuo.mutation.ClearMessage()
	return cuo
}

// RemoveMessageIDs removes the "message" edge to Message entities by IDs.
func (cuo *ChannelUpdateOne) RemoveMessageIDs(ids ...uint64) *ChannelUpdateOne {
	cuo.mutation.RemoveMessageIDs(ids...)
	return cuo
}

// RemoveMessage removes "message" edges to Message entities.
func (cuo *ChannelUpdateOne) RemoveMessage(m ...*Message) *ChannelUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMessageIDs(ids...)
}

// ClearRole clears all "role" edges to the Role entity.
func (cuo *ChannelUpdateOne) ClearRole() *ChannelUpdateOne {
	cuo.mutation.ClearRole()
	return cuo
}

// RemoveRoleIDs removes the "role" edge to Role entities by IDs.
func (cuo *ChannelUpdateOne) RemoveRoleIDs(ids ...uint64) *ChannelUpdateOne {
	cuo.mutation.RemoveRoleIDs(ids...)
	return cuo
}

// RemoveRole removes "role" edges to Role entities.
func (cuo *ChannelUpdateOne) RemoveRole(r ...*Role) *ChannelUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveRoleIDs(ids...)
}

// ClearPermissionNode clears all "permission_node" edges to the PermissionNode entity.
func (cuo *ChannelUpdateOne) ClearPermissionNode() *ChannelUpdateOne {
	cuo.mutation.ClearPermissionNode()
	return cuo
}

// RemovePermissionNodeIDs removes the "permission_node" edge to PermissionNode entities by IDs.
func (cuo *ChannelUpdateOne) RemovePermissionNodeIDs(ids ...int) *ChannelUpdateOne {
	cuo.mutation.RemovePermissionNodeIDs(ids...)
	return cuo
}

// RemovePermissionNode removes "permission_node" edges to PermissionNode entities.
func (cuo *ChannelUpdateOne) RemovePermissionNode(p ...*PermissionNode) *ChannelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePermissionNodeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChannelUpdateOne) Select(field string, fields ...string) *ChannelUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Channel entity.
func (cuo *ChannelUpdateOne) Save(ctx context.Context) (*Channel, error) {
	var (
		err  error
		node *Channel
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChannelUpdateOne) SaveX(ctx context.Context) *Channel {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChannelUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChannelUpdateOne) sqlSave(ctx context.Context) (_node *Channel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: channel.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Channel.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for _, f := range fields {
			if !channel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldName,
		})
	}
	if value, ok := cuo.mutation.Kind(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: channel.FieldKind,
		})
	}
	if value, ok := cuo.mutation.AddedKind(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: channel.FieldKind,
		})
	}
	if value, ok := cuo.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldPosition,
		})
	}
	if value, ok := cuo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: channel.FieldMetadata,
		})
	}
	if cuo.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channel.GuildTable,
			Columns: []string{channel.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channel.GuildTable,
			Columns: []string{channel.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessageTable,
			Columns: []string{channel.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMessageIDs(); len(nodes) > 0 && !cuo.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessageTable,
			Columns: []string{channel.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessageTable,
			Columns: []string{channel.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.RoleTable,
			Columns: []string{channel.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedRoleIDs(); len(nodes) > 0 && !cuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.RoleTable,
			Columns: []string{channel.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.RoleTable,
			Columns: []string{channel.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PermissionNodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.PermissionNodeTable,
			Columns: []string{channel.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPermissionNodeIDs(); len(nodes) > 0 && !cuo.mutation.PermissionNodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.PermissionNodeTable,
			Columns: []string{channel.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PermissionNodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.PermissionNodeTable,
			Columns: []string{channel.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Channel{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
