// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/permissionnode"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
)

// PermissionNodeQuery is the builder for querying PermissionNode entities.
type PermissionNodeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PermissionNode
	// eager-loading edges.
	withRole    *RoleQuery
	withGuild   *GuildQuery
	withChannel *ChannelQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PermissionNodeQuery builder.
func (pnq *PermissionNodeQuery) Where(ps ...predicate.PermissionNode) *PermissionNodeQuery {
	pnq.predicates = append(pnq.predicates, ps...)
	return pnq
}

// Limit adds a limit step to the query.
func (pnq *PermissionNodeQuery) Limit(limit int) *PermissionNodeQuery {
	pnq.limit = &limit
	return pnq
}

// Offset adds an offset step to the query.
func (pnq *PermissionNodeQuery) Offset(offset int) *PermissionNodeQuery {
	pnq.offset = &offset
	return pnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pnq *PermissionNodeQuery) Unique(unique bool) *PermissionNodeQuery {
	pnq.unique = &unique
	return pnq
}

// Order adds an order step to the query.
func (pnq *PermissionNodeQuery) Order(o ...OrderFunc) *PermissionNodeQuery {
	pnq.order = append(pnq.order, o...)
	return pnq
}

// QueryRole chains the current query on the "role" edge.
func (pnq *PermissionNodeQuery) QueryRole() *RoleQuery {
	query := &RoleQuery{config: pnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permissionnode.Table, permissionnode.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, permissionnode.RoleTable, permissionnode.RoleColumn),
		)
		fromU = sqlgraph.SetNeighbors(pnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGuild chains the current query on the "guild" edge.
func (pnq *PermissionNodeQuery) QueryGuild() *GuildQuery {
	query := &GuildQuery{config: pnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permissionnode.Table, permissionnode.FieldID, selector),
			sqlgraph.To(guild.Table, guild.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, permissionnode.GuildTable, permissionnode.GuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(pnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChannel chains the current query on the "channel" edge.
func (pnq *PermissionNodeQuery) QueryChannel() *ChannelQuery {
	query := &ChannelQuery{config: pnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permissionnode.Table, permissionnode.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, permissionnode.ChannelTable, permissionnode.ChannelColumn),
		)
		fromU = sqlgraph.SetNeighbors(pnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PermissionNode entity from the query.
// Returns a *NotFoundError when no PermissionNode was found.
func (pnq *PermissionNodeQuery) First(ctx context.Context) (*PermissionNode, error) {
	nodes, err := pnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{permissionnode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pnq *PermissionNodeQuery) FirstX(ctx context.Context) *PermissionNode {
	node, err := pnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PermissionNode ID from the query.
// Returns a *NotFoundError when no PermissionNode ID was found.
func (pnq *PermissionNodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{permissionnode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pnq *PermissionNodeQuery) FirstIDX(ctx context.Context) int {
	id, err := pnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PermissionNode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PermissionNode entity is not found.
// Returns a *NotFoundError when no PermissionNode entities are found.
func (pnq *PermissionNodeQuery) Only(ctx context.Context) (*PermissionNode, error) {
	nodes, err := pnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{permissionnode.Label}
	default:
		return nil, &NotSingularError{permissionnode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pnq *PermissionNodeQuery) OnlyX(ctx context.Context) *PermissionNode {
	node, err := pnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PermissionNode ID in the query.
// Returns a *NotSingularError when exactly one PermissionNode ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pnq *PermissionNodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = &NotSingularError{permissionnode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pnq *PermissionNodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PermissionNodes.
func (pnq *PermissionNodeQuery) All(ctx context.Context) ([]*PermissionNode, error) {
	if err := pnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pnq *PermissionNodeQuery) AllX(ctx context.Context) []*PermissionNode {
	nodes, err := pnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PermissionNode IDs.
func (pnq *PermissionNodeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pnq.Select(permissionnode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pnq *PermissionNodeQuery) IDsX(ctx context.Context) []int {
	ids, err := pnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pnq *PermissionNodeQuery) Count(ctx context.Context) (int, error) {
	if err := pnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pnq *PermissionNodeQuery) CountX(ctx context.Context) int {
	count, err := pnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pnq *PermissionNodeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pnq *PermissionNodeQuery) ExistX(ctx context.Context) bool {
	exist, err := pnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PermissionNodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pnq *PermissionNodeQuery) Clone() *PermissionNodeQuery {
	if pnq == nil {
		return nil
	}
	return &PermissionNodeQuery{
		config:      pnq.config,
		limit:       pnq.limit,
		offset:      pnq.offset,
		order:       append([]OrderFunc{}, pnq.order...),
		predicates:  append([]predicate.PermissionNode{}, pnq.predicates...),
		withRole:    pnq.withRole.Clone(),
		withGuild:   pnq.withGuild.Clone(),
		withChannel: pnq.withChannel.Clone(),
		// clone intermediate query.
		sql:  pnq.sql.Clone(),
		path: pnq.path,
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (pnq *PermissionNodeQuery) WithRole(opts ...func(*RoleQuery)) *PermissionNodeQuery {
	query := &RoleQuery{config: pnq.config}
	for _, opt := range opts {
		opt(query)
	}
	pnq.withRole = query
	return pnq
}

// WithGuild tells the query-builder to eager-load the nodes that are connected to
// the "guild" edge. The optional arguments are used to configure the query builder of the edge.
func (pnq *PermissionNodeQuery) WithGuild(opts ...func(*GuildQuery)) *PermissionNodeQuery {
	query := &GuildQuery{config: pnq.config}
	for _, opt := range opts {
		opt(query)
	}
	pnq.withGuild = query
	return pnq
}

// WithChannel tells the query-builder to eager-load the nodes that are connected to
// the "channel" edge. The optional arguments are used to configure the query builder of the edge.
func (pnq *PermissionNodeQuery) WithChannel(opts ...func(*ChannelQuery)) *PermissionNodeQuery {
	query := &ChannelQuery{config: pnq.config}
	for _, opt := range opts {
		opt(query)
	}
	pnq.withChannel = query
	return pnq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Node string `json:"node,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PermissionNode.Query().
//		GroupBy(permissionnode.FieldNode).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
//
func (pnq *PermissionNodeQuery) GroupBy(field string, fields ...string) *PermissionNodeGroupBy {
	group := &PermissionNodeGroupBy{config: pnq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pnq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Node string `json:"node,omitempty"`
//	}
//
//	client.PermissionNode.Query().
//		Select(permissionnode.FieldNode).
//		Scan(ctx, &v)
//
func (pnq *PermissionNodeQuery) Select(field string, fields ...string) *PermissionNodeSelect {
	pnq.fields = append([]string{field}, fields...)
	return &PermissionNodeSelect{PermissionNodeQuery: pnq}
}

func (pnq *PermissionNodeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pnq.fields {
		if !permissionnode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if pnq.path != nil {
		prev, err := pnq.path(ctx)
		if err != nil {
			return err
		}
		pnq.sql = prev
	}
	return nil
}

func (pnq *PermissionNodeQuery) sqlAll(ctx context.Context) ([]*PermissionNode, error) {
	var (
		nodes       = []*PermissionNode{}
		withFKs     = pnq.withFKs
		_spec       = pnq.querySpec()
		loadedTypes = [3]bool{
			pnq.withRole != nil,
			pnq.withGuild != nil,
			pnq.withChannel != nil,
		}
	)
	if pnq.withRole != nil || pnq.withGuild != nil || pnq.withChannel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, permissionnode.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PermissionNode{config: pnq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("entgen: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pnq.withRole; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*PermissionNode)
		for i := range nodes {
			if nodes[i].role_permission_node == nil {
				continue
			}
			fk := *nodes[i].role_permission_node
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(role.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "role_permission_node" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Role = n
			}
		}
	}

	if query := pnq.withGuild; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*PermissionNode)
		for i := range nodes {
			if nodes[i].guild_permission_node == nil {
				continue
			}
			fk := *nodes[i].guild_permission_node
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(guild.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "guild_permission_node" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Guild = n
			}
		}
	}

	if query := pnq.withChannel; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*PermissionNode)
		for i := range nodes {
			if nodes[i].channel_permission_node == nil {
				continue
			}
			fk := *nodes[i].channel_permission_node
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(channel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "channel_permission_node" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Channel = n
			}
		}
	}

	return nodes, nil
}

func (pnq *PermissionNodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pnq.querySpec()
	return sqlgraph.CountNodes(ctx, pnq.driver, _spec)
}

func (pnq *PermissionNodeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pnq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entgen: check existence: %w", err)
	}
	return n > 0, nil
}

func (pnq *PermissionNodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permissionnode.Table,
			Columns: permissionnode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permissionnode.FieldID,
			},
		},
		From:   pnq.sql,
		Unique: true,
	}
	if unique := pnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissionnode.FieldID)
		for i := range fields {
			if fields[i] != permissionnode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pnq *PermissionNodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pnq.driver.Dialect())
	t1 := builder.Table(permissionnode.Table)
	selector := builder.Select(t1.Columns(permissionnode.Columns...)...).From(t1)
	if pnq.sql != nil {
		selector = pnq.sql
		selector.Select(selector.Columns(permissionnode.Columns...)...)
	}
	for _, p := range pnq.predicates {
		p(selector)
	}
	for _, p := range pnq.order {
		p(selector)
	}
	if offset := pnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PermissionNodeGroupBy is the group-by builder for PermissionNode entities.
type PermissionNodeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pngb *PermissionNodeGroupBy) Aggregate(fns ...AggregateFunc) *PermissionNodeGroupBy {
	pngb.fns = append(pngb.fns, fns...)
	return pngb
}

// Scan applies the group-by query and scans the result into the given value.
func (pngb *PermissionNodeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pngb.path(ctx)
	if err != nil {
		return err
	}
	pngb.sql = query
	return pngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) StringsX(ctx context.Context) []string {
	v, err := pngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) StringX(ctx context.Context) string {
	v, err := pngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) IntsX(ctx context.Context) []int {
	v, err := pngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) IntX(ctx context.Context) int {
	v, err := pngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PermissionNodeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pngb *PermissionNodeGroupBy) BoolX(ctx context.Context) bool {
	v, err := pngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pngb *PermissionNodeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pngb.fields {
		if !permissionnode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pngb *PermissionNodeGroupBy) sqlQuery() *sql.Selector {
	selector := pngb.sql
	columns := make([]string, 0, len(pngb.fields)+len(pngb.fns))
	columns = append(columns, pngb.fields...)
	for _, fn := range pngb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pngb.fields...)
}

// PermissionNodeSelect is the builder for selecting fields of PermissionNode entities.
type PermissionNodeSelect struct {
	*PermissionNodeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pns *PermissionNodeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pns.prepareQuery(ctx); err != nil {
		return err
	}
	pns.sql = pns.PermissionNodeQuery.sqlQuery(ctx)
	return pns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pns *PermissionNodeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pns *PermissionNodeSelect) StringsX(ctx context.Context) []string {
	v, err := pns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pns *PermissionNodeSelect) StringX(ctx context.Context) string {
	v, err := pns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pns *PermissionNodeSelect) IntsX(ctx context.Context) []int {
	v, err := pns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pns *PermissionNodeSelect) IntX(ctx context.Context) int {
	v, err := pns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pns *PermissionNodeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pns *PermissionNodeSelect) Float64X(ctx context.Context) float64 {
	v, err := pns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("entgen: PermissionNodeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pns *PermissionNodeSelect) BoolsX(ctx context.Context) []bool {
	v, err := pns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pns *PermissionNodeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{permissionnode.Label}
	default:
		err = fmt.Errorf("entgen: PermissionNodeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pns *PermissionNodeSelect) BoolX(ctx context.Context) bool {
	v, err := pns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pns *PermissionNodeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pns.sqlQuery().Query()
	if err := pns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pns *PermissionNodeSelect) sqlQuery() sql.Querier {
	selector := pns.sql
	selector.Select(selector.Columns(pns.fields...)...)
	return selector
}
