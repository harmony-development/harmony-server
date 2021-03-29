// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/emote"
	"github.com/harmony-development/legato/server/db/ent/entgen/emotepack"
	"github.com/harmony-development/legato/server/db/ent/entgen/file"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// EmoteQuery is the builder for querying Emote entities.
type EmoteQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Emote
	// eager-loading edges.
	withEmotepack *EmotePackQuery
	withFile      *FileQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmoteQuery builder.
func (eq *EmoteQuery) Where(ps ...predicate.Emote) *EmoteQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *EmoteQuery) Limit(limit int) *EmoteQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *EmoteQuery) Offset(offset int) *EmoteQuery {
	eq.offset = &offset
	return eq
}

// Order adds an order step to the query.
func (eq *EmoteQuery) Order(o ...OrderFunc) *EmoteQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryEmotepack chains the current query on the "emotepack" edge.
func (eq *EmoteQuery) QueryEmotepack() *EmotePackQuery {
	query := &EmotePackQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(emote.Table, emote.FieldID, selector),
			sqlgraph.To(emotepack.Table, emotepack.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, emote.EmotepackTable, emote.EmotepackColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFile chains the current query on the "file" edge.
func (eq *EmoteQuery) QueryFile() *FileQuery {
	query := &FileQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(emote.Table, emote.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, emote.FileTable, emote.FileColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Emote entity from the query.
// Returns a *NotFoundError when no Emote was found.
func (eq *EmoteQuery) First(ctx context.Context) (*Emote, error) {
	nodes, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EmoteQuery) FirstX(ctx context.Context) *Emote {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Emote ID from the query.
// Returns a *NotFoundError when no Emote ID was found.
func (eq *EmoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EmoteQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Emote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Emote entity is not found.
// Returns a *NotFoundError when no Emote entities are found.
func (eq *EmoteQuery) Only(ctx context.Context) (*Emote, error) {
	nodes, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emote.Label}
	default:
		return nil, &NotSingularError{emote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EmoteQuery) OnlyX(ctx context.Context) *Emote {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Emote ID in the query.
// Returns a *NotSingularError when exactly one Emote ID is not found.
// Returns a *NotFoundError when no entities are found.
func (eq *EmoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = &NotSingularError{emote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EmoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Emotes.
func (eq *EmoteQuery) All(ctx context.Context) ([]*Emote, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *EmoteQuery) AllX(ctx context.Context) []*Emote {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Emote IDs.
func (eq *EmoteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := eq.Select(emote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EmoteQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EmoteQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EmoteQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EmoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EmoteQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EmoteQuery) Clone() *EmoteQuery {
	if eq == nil {
		return nil
	}
	return &EmoteQuery{
		config:        eq.config,
		limit:         eq.limit,
		offset:        eq.offset,
		order:         append([]OrderFunc{}, eq.order...),
		predicates:    append([]predicate.Emote{}, eq.predicates...),
		withEmotepack: eq.withEmotepack.Clone(),
		withFile:      eq.withFile.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithEmotepack tells the query-builder to eager-load the nodes that are connected to
// the "emotepack" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EmoteQuery) WithEmotepack(opts ...func(*EmotePackQuery)) *EmoteQuery {
	query := &EmotePackQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withEmotepack = query
	return eq
}

// WithFile tells the query-builder to eager-load the nodes that are connected to
// the "file" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EmoteQuery) WithFile(opts ...func(*FileQuery)) *EmoteQuery {
	query := &FileQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withFile = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Emote.Query().
//		GroupBy(emote.FieldName).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
//
func (eq *EmoteQuery) GroupBy(field string, fields ...string) *EmoteGroupBy {
	group := &EmoteGroupBy{config: eq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Emote.Query().
//		Select(emote.FieldName).
//		Scan(ctx, &v)
//
func (eq *EmoteQuery) Select(field string, fields ...string) *EmoteSelect {
	eq.fields = append([]string{field}, fields...)
	return &EmoteSelect{EmoteQuery: eq}
}

func (eq *EmoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eq.fields {
		if !emote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EmoteQuery) sqlAll(ctx context.Context) ([]*Emote, error) {
	var (
		nodes       = []*Emote{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [2]bool{
			eq.withEmotepack != nil,
			eq.withFile != nil,
		}
	)
	if eq.withEmotepack != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, emote.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Emote{config: eq.config}
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
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := eq.withEmotepack; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*Emote)
		for i := range nodes {
			fk := nodes[i].emote_pack_emote
			if fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(emotepack.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "emote_pack_emote" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Emotepack = n
			}
		}
	}

	if query := eq.withFile; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Emote)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.File(func(s *sql.Selector) {
			s.Where(sql.InValues(emote.FileColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.emote_file
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "emote_file" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "emote_file" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.File = n
		}
	}

	return nodes, nil
}

func (eq *EmoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EmoteQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := eq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entgen: check existence: %w", err)
	}
	return n > 0, nil
}

func (eq *EmoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emote.Table,
			Columns: emote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: emote.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if fields := eq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emote.FieldID)
		for i := range fields {
			if fields[i] != emote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, emote.ValidColumn)
			}
		}
	}
	return _spec
}

func (eq *EmoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(emote.Table)
	selector := builder.Select(t1.Columns(emote.Columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(emote.Columns...)...)
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector, emote.ValidColumn)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmoteGroupBy is the group-by builder for Emote entities.
type EmoteGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EmoteGroupBy) Aggregate(fns ...AggregateFunc) *EmoteGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scans the result into the given value.
func (egb *EmoteGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (egb *EmoteGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := egb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("entgen: EmoteGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (egb *EmoteGroupBy) StringsX(ctx context.Context) []string {
	v, err := egb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = egb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (egb *EmoteGroupBy) StringX(ctx context.Context) string {
	v, err := egb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("entgen: EmoteGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (egb *EmoteGroupBy) IntsX(ctx context.Context) []int {
	v, err := egb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = egb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (egb *EmoteGroupBy) IntX(ctx context.Context) int {
	v, err := egb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("entgen: EmoteGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (egb *EmoteGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := egb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = egb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (egb *EmoteGroupBy) Float64X(ctx context.Context) float64 {
	v, err := egb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("entgen: EmoteGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (egb *EmoteGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := egb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EmoteGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = egb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (egb *EmoteGroupBy) BoolX(ctx context.Context) bool {
	v, err := egb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (egb *EmoteGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range egb.fields {
		if !emote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := egb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *EmoteGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql
	columns := make([]string, 0, len(egb.fields)+len(egb.fns))
	columns = append(columns, egb.fields...)
	for _, fn := range egb.fns {
		columns = append(columns, fn(selector, emote.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(egb.fields...)
}

// EmoteSelect is the builder for selecting fields of Emote entities.
type EmoteSelect struct {
	*EmoteQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (es *EmoteSelect) Scan(ctx context.Context, v interface{}) error {
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	es.sql = es.EmoteQuery.sqlQuery(ctx)
	return es.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (es *EmoteSelect) ScanX(ctx context.Context, v interface{}) {
	if err := es.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) Strings(ctx context.Context) ([]string, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("entgen: EmoteSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (es *EmoteSelect) StringsX(ctx context.Context) []string {
	v, err := es.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = es.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (es *EmoteSelect) StringX(ctx context.Context) string {
	v, err := es.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) Ints(ctx context.Context) ([]int, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("entgen: EmoteSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (es *EmoteSelect) IntsX(ctx context.Context) []int {
	v, err := es.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = es.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (es *EmoteSelect) IntX(ctx context.Context) int {
	v, err := es.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("entgen: EmoteSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (es *EmoteSelect) Float64sX(ctx context.Context) []float64 {
	v, err := es.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = es.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (es *EmoteSelect) Float64X(ctx context.Context) float64 {
	v, err := es.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("entgen: EmoteSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (es *EmoteSelect) BoolsX(ctx context.Context) []bool {
	v, err := es.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (es *EmoteSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = es.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emote.Label}
	default:
		err = fmt.Errorf("entgen: EmoteSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (es *EmoteSelect) BoolX(ctx context.Context) bool {
	v, err := es.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (es *EmoteSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := es.sqlQuery().Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (es *EmoteSelect) sqlQuery() sql.Querier {
	selector := es.sql
	selector.Select(selector.Columns(es.fields...)...)
	return selector
}
