// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingproblemstaffdata"
	"170-ag/ent/generated/predicate"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CodingProblemStaffDataQuery is the builder for querying CodingProblemStaffData entities.
type CodingProblemStaffDataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CodingProblemStaffData
	// eager-loading edges.
	withCodingProblem *CodingProblemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CodingProblemStaffDataQuery builder.
func (cpsdq *CodingProblemStaffDataQuery) Where(ps ...predicate.CodingProblemStaffData) *CodingProblemStaffDataQuery {
	cpsdq.predicates = append(cpsdq.predicates, ps...)
	return cpsdq
}

// Limit adds a limit step to the query.
func (cpsdq *CodingProblemStaffDataQuery) Limit(limit int) *CodingProblemStaffDataQuery {
	cpsdq.limit = &limit
	return cpsdq
}

// Offset adds an offset step to the query.
func (cpsdq *CodingProblemStaffDataQuery) Offset(offset int) *CodingProblemStaffDataQuery {
	cpsdq.offset = &offset
	return cpsdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cpsdq *CodingProblemStaffDataQuery) Unique(unique bool) *CodingProblemStaffDataQuery {
	cpsdq.unique = &unique
	return cpsdq
}

// Order adds an order step to the query.
func (cpsdq *CodingProblemStaffDataQuery) Order(o ...OrderFunc) *CodingProblemStaffDataQuery {
	cpsdq.order = append(cpsdq.order, o...)
	return cpsdq
}

// QueryCodingProblem chains the current query on the "coding_problem" edge.
func (cpsdq *CodingProblemStaffDataQuery) QueryCodingProblem() *CodingProblemQuery {
	query := &CodingProblemQuery{config: cpsdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpsdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpsdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblemstaffdata.Table, codingproblemstaffdata.FieldID, selector),
			sqlgraph.To(codingproblem.Table, codingproblem.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, codingproblemstaffdata.CodingProblemTable, codingproblemstaffdata.CodingProblemColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpsdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CodingProblemStaffData entity from the query.
// Returns a *NotFoundError when no CodingProblemStaffData was found.
func (cpsdq *CodingProblemStaffDataQuery) First(ctx context.Context) (*CodingProblemStaffData, error) {
	nodes, err := cpsdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{codingproblemstaffdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) FirstX(ctx context.Context) *CodingProblemStaffData {
	node, err := cpsdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CodingProblemStaffData ID from the query.
// Returns a *NotFoundError when no CodingProblemStaffData ID was found.
func (cpsdq *CodingProblemStaffDataQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpsdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{codingproblemstaffdata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) FirstIDX(ctx context.Context) int {
	id, err := cpsdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CodingProblemStaffData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CodingProblemStaffData entity is not found.
// Returns a *NotFoundError when no CodingProblemStaffData entities are found.
func (cpsdq *CodingProblemStaffDataQuery) Only(ctx context.Context) (*CodingProblemStaffData, error) {
	nodes, err := cpsdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{codingproblemstaffdata.Label}
	default:
		return nil, &NotSingularError{codingproblemstaffdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) OnlyX(ctx context.Context) *CodingProblemStaffData {
	node, err := cpsdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CodingProblemStaffData ID in the query.
// Returns a *NotSingularError when exactly one CodingProblemStaffData ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cpsdq *CodingProblemStaffDataQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpsdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = &NotSingularError{codingproblemstaffdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) OnlyIDX(ctx context.Context) int {
	id, err := cpsdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CodingProblemStaffDataSlice.
func (cpsdq *CodingProblemStaffDataQuery) All(ctx context.Context) ([]*CodingProblemStaffData, error) {
	if err := cpsdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cpsdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) AllX(ctx context.Context) []*CodingProblemStaffData {
	nodes, err := cpsdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CodingProblemStaffData IDs.
func (cpsdq *CodingProblemStaffDataQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cpsdq.Select(codingproblemstaffdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) IDsX(ctx context.Context) []int {
	ids, err := cpsdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpsdq *CodingProblemStaffDataQuery) Count(ctx context.Context) (int, error) {
	if err := cpsdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cpsdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) CountX(ctx context.Context) int {
	count, err := cpsdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpsdq *CodingProblemStaffDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := cpsdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cpsdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpsdq *CodingProblemStaffDataQuery) ExistX(ctx context.Context) bool {
	exist, err := cpsdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CodingProblemStaffDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpsdq *CodingProblemStaffDataQuery) Clone() *CodingProblemStaffDataQuery {
	if cpsdq == nil {
		return nil
	}
	return &CodingProblemStaffDataQuery{
		config:            cpsdq.config,
		limit:             cpsdq.limit,
		offset:            cpsdq.offset,
		order:             append([]OrderFunc{}, cpsdq.order...),
		predicates:        append([]predicate.CodingProblemStaffData{}, cpsdq.predicates...),
		withCodingProblem: cpsdq.withCodingProblem.Clone(),
		// clone intermediate query.
		sql:  cpsdq.sql.Clone(),
		path: cpsdq.path,
	}
}

// WithCodingProblem tells the query-builder to eager-load the nodes that are connected to
// the "coding_problem" edge. The optional arguments are used to configure the query builder of the edge.
func (cpsdq *CodingProblemStaffDataQuery) WithCodingProblem(opts ...func(*CodingProblemQuery)) *CodingProblemStaffDataQuery {
	query := &CodingProblemQuery{config: cpsdq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpsdq.withCodingProblem = query
	return cpsdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Input string `json:"input,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CodingProblemStaffData.Query().
//		GroupBy(codingproblemstaffdata.FieldInput).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
//
func (cpsdq *CodingProblemStaffDataQuery) GroupBy(field string, fields ...string) *CodingProblemStaffDataGroupBy {
	group := &CodingProblemStaffDataGroupBy{config: cpsdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cpsdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cpsdq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Input string `json:"input,omitempty"`
//	}
//
//	client.CodingProblemStaffData.Query().
//		Select(codingproblemstaffdata.FieldInput).
//		Scan(ctx, &v)
//
func (cpsdq *CodingProblemStaffDataQuery) Select(fields ...string) *CodingProblemStaffDataSelect {
	cpsdq.fields = append(cpsdq.fields, fields...)
	return &CodingProblemStaffDataSelect{CodingProblemStaffDataQuery: cpsdq}
}

func (cpsdq *CodingProblemStaffDataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cpsdq.fields {
		if !codingproblemstaffdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if cpsdq.path != nil {
		prev, err := cpsdq.path(ctx)
		if err != nil {
			return err
		}
		cpsdq.sql = prev
	}
	return nil
}

func (cpsdq *CodingProblemStaffDataQuery) sqlAll(ctx context.Context) ([]*CodingProblemStaffData, error) {
	var (
		nodes       = []*CodingProblemStaffData{}
		_spec       = cpsdq.querySpec()
		loadedTypes = [1]bool{
			cpsdq.withCodingProblem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CodingProblemStaffData{config: cpsdq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("generated: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, cpsdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cpsdq.withCodingProblem; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*CodingProblemStaffData)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.CodingProblem(func(s *sql.Selector) {
			s.Where(sql.InValues(codingproblemstaffdata.CodingProblemColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.coding_problem_staff_data_coding_problem
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "coding_problem_staff_data_coding_problem" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coding_problem_staff_data_coding_problem" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.CodingProblem = n
		}
	}

	return nodes, nil
}

func (cpsdq *CodingProblemStaffDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cpsdq.querySpec()
	return sqlgraph.CountNodes(ctx, cpsdq.driver, _spec)
}

func (cpsdq *CodingProblemStaffDataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpsdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("generated: check existence: %w", err)
	}
	return n > 0, nil
}

func (cpsdq *CodingProblemStaffDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingproblemstaffdata.Table,
			Columns: codingproblemstaffdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingproblemstaffdata.FieldID,
			},
		},
		From:   cpsdq.sql,
		Unique: true,
	}
	if unique := cpsdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cpsdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codingproblemstaffdata.FieldID)
		for i := range fields {
			if fields[i] != codingproblemstaffdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cpsdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cpsdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cpsdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cpsdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cpsdq *CodingProblemStaffDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cpsdq.driver.Dialect())
	t1 := builder.Table(codingproblemstaffdata.Table)
	columns := cpsdq.fields
	if len(columns) == 0 {
		columns = codingproblemstaffdata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cpsdq.sql != nil {
		selector = cpsdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range cpsdq.predicates {
		p(selector)
	}
	for _, p := range cpsdq.order {
		p(selector)
	}
	if offset := cpsdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpsdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CodingProblemStaffDataGroupBy is the group-by builder for CodingProblemStaffData entities.
type CodingProblemStaffDataGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpsdgb *CodingProblemStaffDataGroupBy) Aggregate(fns ...AggregateFunc) *CodingProblemStaffDataGroupBy {
	cpsdgb.fns = append(cpsdgb.fns, fns...)
	return cpsdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cpsdgb *CodingProblemStaffDataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cpsdgb.path(ctx)
	if err != nil {
		return err
	}
	cpsdgb.sql = query
	return cpsdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cpsdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cpsdgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cpsdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) StringsX(ctx context.Context) []string {
	v, err := cpsdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cpsdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) StringX(ctx context.Context) string {
	v, err := cpsdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cpsdgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cpsdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) IntsX(ctx context.Context) []int {
	v, err := cpsdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cpsdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) IntX(ctx context.Context) int {
	v, err := cpsdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cpsdgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cpsdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cpsdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cpsdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cpsdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cpsdgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cpsdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cpsdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpsdgb *CodingProblemStaffDataGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cpsdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cpsdgb *CodingProblemStaffDataGroupBy) BoolX(ctx context.Context) bool {
	v, err := cpsdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cpsdgb *CodingProblemStaffDataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cpsdgb.fields {
		if !codingproblemstaffdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cpsdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cpsdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpsdgb *CodingProblemStaffDataGroupBy) sqlQuery() *sql.Selector {
	selector := cpsdgb.sql.Select()
	aggregation := make([]string, 0, len(cpsdgb.fns))
	for _, fn := range cpsdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cpsdgb.fields)+len(cpsdgb.fns))
		for _, f := range cpsdgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cpsdgb.fields...)...)
}

// CodingProblemStaffDataSelect is the builder for selecting fields of CodingProblemStaffData entities.
type CodingProblemStaffDataSelect struct {
	*CodingProblemStaffDataQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cpsds *CodingProblemStaffDataSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cpsds.prepareQuery(ctx); err != nil {
		return err
	}
	cpsds.sql = cpsds.CodingProblemStaffDataQuery.sqlQuery(ctx)
	return cpsds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cpsds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cpsds.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cpsds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) StringsX(ctx context.Context) []string {
	v, err := cpsds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cpsds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) StringX(ctx context.Context) string {
	v, err := cpsds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cpsds.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cpsds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) IntsX(ctx context.Context) []int {
	v, err := cpsds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cpsds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) IntX(ctx context.Context) int {
	v, err := cpsds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cpsds.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cpsds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cpsds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cpsds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) Float64X(ctx context.Context) float64 {
	v, err := cpsds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cpsds.fields) > 1 {
		return nil, errors.New("generated: CodingProblemStaffDataSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cpsds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) BoolsX(ctx context.Context) []bool {
	v, err := cpsds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cpsds *CodingProblemStaffDataSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cpsds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblemstaffdata.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemStaffDataSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cpsds *CodingProblemStaffDataSelect) BoolX(ctx context.Context) bool {
	v, err := cpsds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cpsds *CodingProblemStaffDataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cpsds.sql.Query()
	if err := cpsds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
