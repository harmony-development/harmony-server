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
	"github.com/harmony-development/legato/server/db/ent/entgen/foreignuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// ForeignUserQuery is the builder for querying ForeignUser entities.
type ForeignUserQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.ForeignUser
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ForeignUserQuery builder.
func (fuq *ForeignUserQuery) Where(ps ...predicate.ForeignUser) *ForeignUserQuery {
	fuq.predicates = append(fuq.predicates, ps...)
	return fuq
}

// Limit adds a limit step to the query.
func (fuq *ForeignUserQuery) Limit(limit int) *ForeignUserQuery {
	fuq.limit = &limit
	return fuq
}

// Offset adds an offset step to the query.
func (fuq *ForeignUserQuery) Offset(offset int) *ForeignUserQuery {
	fuq.offset = &offset
	return fuq
}

// Order adds an order step to the query.
func (fuq *ForeignUserQuery) Order(o ...OrderFunc) *ForeignUserQuery {
	fuq.order = append(fuq.order, o...)
	return fuq
}

// QueryUser chains the current query on the "user" edge.
func (fuq *ForeignUserQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: fuq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(foreignuser.Table, foreignuser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, foreignuser.UserTable, foreignuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(fuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ForeignUser entity from the query.
// Returns a *NotFoundError when no ForeignUser was found.
func (fuq *ForeignUserQuery) First(ctx context.Context) (*ForeignUser, error) {
	nodes, err := fuq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{foreignuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fuq *ForeignUserQuery) FirstX(ctx context.Context) *ForeignUser {
	node, err := fuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ForeignUser ID from the query.
// Returns a *NotFoundError when no ForeignUser ID was found.
func (fuq *ForeignUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fuq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{foreignuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fuq *ForeignUserQuery) FirstIDX(ctx context.Context) int {
	id, err := fuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ForeignUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ForeignUser entity is not found.
// Returns a *NotFoundError when no ForeignUser entities are found.
func (fuq *ForeignUserQuery) Only(ctx context.Context) (*ForeignUser, error) {
	nodes, err := fuq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{foreignuser.Label}
	default:
		return nil, &NotSingularError{foreignuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fuq *ForeignUserQuery) OnlyX(ctx context.Context) *ForeignUser {
	node, err := fuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ForeignUser ID in the query.
// Returns a *NotSingularError when exactly one ForeignUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (fuq *ForeignUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fuq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = &NotSingularError{foreignuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fuq *ForeignUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := fuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ForeignUsers.
func (fuq *ForeignUserQuery) All(ctx context.Context) ([]*ForeignUser, error) {
	if err := fuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fuq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fuq *ForeignUserQuery) AllX(ctx context.Context) []*ForeignUser {
	nodes, err := fuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ForeignUser IDs.
func (fuq *ForeignUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fuq.Select(foreignuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fuq *ForeignUserQuery) IDsX(ctx context.Context) []int {
	ids, err := fuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fuq *ForeignUserQuery) Count(ctx context.Context) (int, error) {
	if err := fuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fuq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fuq *ForeignUserQuery) CountX(ctx context.Context) int {
	count, err := fuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fuq *ForeignUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := fuq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fuq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fuq *ForeignUserQuery) ExistX(ctx context.Context) bool {
	exist, err := fuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ForeignUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fuq *ForeignUserQuery) Clone() *ForeignUserQuery {
	if fuq == nil {
		return nil
	}
	return &ForeignUserQuery{
		config:     fuq.config,
		limit:      fuq.limit,
		offset:     fuq.offset,
		order:      append([]OrderFunc{}, fuq.order...),
		predicates: append([]predicate.ForeignUser{}, fuq.predicates...),
		withUser:   fuq.withUser.Clone(),
		// clone intermediate query.
		sql:  fuq.sql.Clone(),
		path: fuq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (fuq *ForeignUserQuery) WithUser(opts ...func(*UserQuery)) *ForeignUserQuery {
	query := &UserQuery{config: fuq.config}
	for _, opt := range opts {
		opt(query)
	}
	fuq.withUser = query
	return fuq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Foreignid uint64 `json:"foreignid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ForeignUser.Query().
//		GroupBy(foreignuser.FieldForeignid).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
//
func (fuq *ForeignUserQuery) GroupBy(field string, fields ...string) *ForeignUserGroupBy {
	group := &ForeignUserGroupBy{config: fuq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fuq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Foreignid uint64 `json:"foreignid,omitempty"`
//	}
//
//	client.ForeignUser.Query().
//		Select(foreignuser.FieldForeignid).
//		Scan(ctx, &v)
//
func (fuq *ForeignUserQuery) Select(field string, fields ...string) *ForeignUserSelect {
	fuq.fields = append([]string{field}, fields...)
	return &ForeignUserSelect{ForeignUserQuery: fuq}
}

func (fuq *ForeignUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fuq.fields {
		if !foreignuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if fuq.path != nil {
		prev, err := fuq.path(ctx)
		if err != nil {
			return err
		}
		fuq.sql = prev
	}
	return nil
}

func (fuq *ForeignUserQuery) sqlAll(ctx context.Context) ([]*ForeignUser, error) {
	var (
		nodes       = []*ForeignUser{}
		withFKs     = fuq.withFKs
		_spec       = fuq.querySpec()
		loadedTypes = [1]bool{
			fuq.withUser != nil,
		}
	)
	if fuq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, foreignuser.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ForeignUser{config: fuq.config}
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
	if err := sqlgraph.QueryNodes(ctx, fuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fuq.withUser; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*ForeignUser)
		for i := range nodes {
			fk := nodes[i].user_foreign_user
			if fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_foreign_user" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (fuq *ForeignUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fuq.querySpec()
	return sqlgraph.CountNodes(ctx, fuq.driver, _spec)
}

func (fuq *ForeignUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fuq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entgen: check existence: %w", err)
	}
	return n > 0, nil
}

func (fuq *ForeignUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   foreignuser.Table,
			Columns: foreignuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: foreignuser.FieldID,
			},
		},
		From:   fuq.sql,
		Unique: true,
	}
	if fields := fuq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, foreignuser.FieldID)
		for i := range fields {
			if fields[i] != foreignuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fuq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fuq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, foreignuser.ValidColumn)
			}
		}
	}
	return _spec
}

func (fuq *ForeignUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fuq.driver.Dialect())
	t1 := builder.Table(foreignuser.Table)
	selector := builder.Select(t1.Columns(foreignuser.Columns...)...).From(t1)
	if fuq.sql != nil {
		selector = fuq.sql
		selector.Select(selector.Columns(foreignuser.Columns...)...)
	}
	for _, p := range fuq.predicates {
		p(selector)
	}
	for _, p := range fuq.order {
		p(selector, foreignuser.ValidColumn)
	}
	if offset := fuq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fuq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForeignUserGroupBy is the group-by builder for ForeignUser entities.
type ForeignUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fugb *ForeignUserGroupBy) Aggregate(fns ...AggregateFunc) *ForeignUserGroupBy {
	fugb.fns = append(fugb.fns, fns...)
	return fugb
}

// Scan applies the group-by query and scans the result into the given value.
func (fugb *ForeignUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fugb.path(ctx)
	if err != nil {
		return err
	}
	fugb.sql = query
	return fugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fugb.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := fugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) StringX(ctx context.Context) string {
	v, err := fugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fugb.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := fugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) IntX(ctx context.Context) int {
	v, err := fugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fugb.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fugb.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fugb *ForeignUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fugb *ForeignUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := fugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fugb *ForeignUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fugb.fields {
		if !foreignuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fugb *ForeignUserGroupBy) sqlQuery() *sql.Selector {
	selector := fugb.sql
	columns := make([]string, 0, len(fugb.fields)+len(fugb.fns))
	columns = append(columns, fugb.fields...)
	for _, fn := range fugb.fns {
		columns = append(columns, fn(selector, foreignuser.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(fugb.fields...)
}

// ForeignUserSelect is the builder for selecting fields of ForeignUser entities.
type ForeignUserSelect struct {
	*ForeignUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fus *ForeignUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fus.prepareQuery(ctx); err != nil {
		return err
	}
	fus.sql = fus.ForeignUserQuery.sqlQuery(ctx)
	return fus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fus *ForeignUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fus.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fus *ForeignUserSelect) StringsX(ctx context.Context) []string {
	v, err := fus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fus *ForeignUserSelect) StringX(ctx context.Context) string {
	v, err := fus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fus.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fus *ForeignUserSelect) IntsX(ctx context.Context) []int {
	v, err := fus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fus *ForeignUserSelect) IntX(ctx context.Context) int {
	v, err := fus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fus.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fus *ForeignUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fus *ForeignUserSelect) Float64X(ctx context.Context) float64 {
	v, err := fus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fus.fields) > 1 {
		return nil, errors.New("entgen: ForeignUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fus *ForeignUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := fus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fus *ForeignUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{foreignuser.Label}
	default:
		err = fmt.Errorf("entgen: ForeignUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fus *ForeignUserSelect) BoolX(ctx context.Context) bool {
	v, err := fus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fus *ForeignUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fus.sqlQuery().Query()
	if err := fus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fus *ForeignUserSelect) sqlQuery() sql.Querier {
	selector := fus.sql
	selector.Select(selector.Columns(fus.fields...)...)
	return selector
}
