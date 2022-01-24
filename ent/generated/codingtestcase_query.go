// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/codingtestcasedata"
	"170-ag/ent/generated/predicate"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CodingTestCaseQuery is the builder for querying CodingTestCase entities.
type CodingTestCaseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CodingTestCase
	// eager-loading edges.
	withCodingProblem *CodingProblemQuery
	withData          *CodingTestCaseDataQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CodingTestCaseQuery builder.
func (ctcq *CodingTestCaseQuery) Where(ps ...predicate.CodingTestCase) *CodingTestCaseQuery {
	ctcq.predicates = append(ctcq.predicates, ps...)
	return ctcq
}

// Limit adds a limit step to the query.
func (ctcq *CodingTestCaseQuery) Limit(limit int) *CodingTestCaseQuery {
	ctcq.limit = &limit
	return ctcq
}

// Offset adds an offset step to the query.
func (ctcq *CodingTestCaseQuery) Offset(offset int) *CodingTestCaseQuery {
	ctcq.offset = &offset
	return ctcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctcq *CodingTestCaseQuery) Unique(unique bool) *CodingTestCaseQuery {
	ctcq.unique = &unique
	return ctcq
}

// Order adds an order step to the query.
func (ctcq *CodingTestCaseQuery) Order(o ...OrderFunc) *CodingTestCaseQuery {
	ctcq.order = append(ctcq.order, o...)
	return ctcq
}

// QueryCodingProblem chains the current query on the "coding_problem" edge.
func (ctcq *CodingTestCaseQuery) QueryCodingProblem() *CodingProblemQuery {
	query := &CodingProblemQuery{config: ctcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(codingtestcase.Table, codingtestcase.FieldID, selector),
			sqlgraph.To(codingproblem.Table, codingproblem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, codingtestcase.CodingProblemTable, codingtestcase.CodingProblemColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryData chains the current query on the "data" edge.
func (ctcq *CodingTestCaseQuery) QueryData() *CodingTestCaseDataQuery {
	query := &CodingTestCaseDataQuery{config: ctcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(codingtestcase.Table, codingtestcase.FieldID, selector),
			sqlgraph.To(codingtestcasedata.Table, codingtestcasedata.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, codingtestcase.DataTable, codingtestcase.DataColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CodingTestCase entity from the query.
// Returns a *NotFoundError when no CodingTestCase was found.
func (ctcq *CodingTestCaseQuery) First(ctx context.Context) (*CodingTestCase, error) {
	nodes, err := ctcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{codingtestcase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) FirstX(ctx context.Context) *CodingTestCase {
	node, err := ctcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CodingTestCase ID from the query.
// Returns a *NotFoundError when no CodingTestCase ID was found.
func (ctcq *CodingTestCaseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{codingtestcase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) FirstIDX(ctx context.Context) int {
	id, err := ctcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CodingTestCase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CodingTestCase entity is not found.
// Returns a *NotFoundError when no CodingTestCase entities are found.
func (ctcq *CodingTestCaseQuery) Only(ctx context.Context) (*CodingTestCase, error) {
	nodes, err := ctcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{codingtestcase.Label}
	default:
		return nil, &NotSingularError{codingtestcase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) OnlyX(ctx context.Context) *CodingTestCase {
	node, err := ctcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CodingTestCase ID in the query.
// Returns a *NotSingularError when exactly one CodingTestCase ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ctcq *CodingTestCaseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = &NotSingularError{codingtestcase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CodingTestCases.
func (ctcq *CodingTestCaseQuery) All(ctx context.Context) ([]*CodingTestCase, error) {
	if err := ctcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ctcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) AllX(ctx context.Context) []*CodingTestCase {
	nodes, err := ctcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CodingTestCase IDs.
func (ctcq *CodingTestCaseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ctcq.Select(codingtestcase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) IDsX(ctx context.Context) []int {
	ids, err := ctcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctcq *CodingTestCaseQuery) Count(ctx context.Context) (int, error) {
	if err := ctcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ctcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) CountX(ctx context.Context) int {
	count, err := ctcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctcq *CodingTestCaseQuery) Exist(ctx context.Context) (bool, error) {
	if err := ctcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ctcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ctcq *CodingTestCaseQuery) ExistX(ctx context.Context) bool {
	exist, err := ctcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CodingTestCaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctcq *CodingTestCaseQuery) Clone() *CodingTestCaseQuery {
	if ctcq == nil {
		return nil
	}
	return &CodingTestCaseQuery{
		config:            ctcq.config,
		limit:             ctcq.limit,
		offset:            ctcq.offset,
		order:             append([]OrderFunc{}, ctcq.order...),
		predicates:        append([]predicate.CodingTestCase{}, ctcq.predicates...),
		withCodingProblem: ctcq.withCodingProblem.Clone(),
		withData:          ctcq.withData.Clone(),
		// clone intermediate query.
		sql:  ctcq.sql.Clone(),
		path: ctcq.path,
	}
}

// WithCodingProblem tells the query-builder to eager-load the nodes that are connected to
// the "coding_problem" edge. The optional arguments are used to configure the query builder of the edge.
func (ctcq *CodingTestCaseQuery) WithCodingProblem(opts ...func(*CodingProblemQuery)) *CodingTestCaseQuery {
	query := &CodingProblemQuery{config: ctcq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctcq.withCodingProblem = query
	return ctcq
}

// WithData tells the query-builder to eager-load the nodes that are connected to
// the "data" edge. The optional arguments are used to configure the query builder of the edge.
func (ctcq *CodingTestCaseQuery) WithData(opts ...func(*CodingTestCaseDataQuery)) *CodingTestCaseQuery {
	query := &CodingTestCaseDataQuery{config: ctcq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctcq.withData = query
	return ctcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CodingTestCase.Query().
//		GroupBy(codingtestcase.FieldCreateTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
//
func (ctcq *CodingTestCaseQuery) GroupBy(field string, fields ...string) *CodingTestCaseGroupBy {
	group := &CodingTestCaseGroupBy{config: ctcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.CodingTestCase.Query().
//		Select(codingtestcase.FieldCreateTime).
//		Scan(ctx, &v)
//
func (ctcq *CodingTestCaseQuery) Select(fields ...string) *CodingTestCaseSelect {
	ctcq.fields = append(ctcq.fields, fields...)
	return &CodingTestCaseSelect{CodingTestCaseQuery: ctcq}
}

func (ctcq *CodingTestCaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ctcq.fields {
		if !codingtestcase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ctcq.path != nil {
		prev, err := ctcq.path(ctx)
		if err != nil {
			return err
		}
		ctcq.sql = prev
	}
	if codingtestcase.Policy == nil {
		return errors.New("generated: uninitialized codingtestcase.Policy (forgotten import generated/runtime?)")
	}
	if err := codingtestcase.Policy.EvalQuery(ctx, ctcq); err != nil {
		return err
	}
	return nil
}

func (ctcq *CodingTestCaseQuery) sqlAll(ctx context.Context) ([]*CodingTestCase, error) {
	var (
		nodes       = []*CodingTestCase{}
		withFKs     = ctcq.withFKs
		_spec       = ctcq.querySpec()
		loadedTypes = [2]bool{
			ctcq.withCodingProblem != nil,
			ctcq.withData != nil,
		}
	)
	if ctcq.withCodingProblem != nil || ctcq.withData != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, codingtestcase.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CodingTestCase{config: ctcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ctcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ctcq.withCodingProblem; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CodingTestCase)
		for i := range nodes {
			if nodes[i].coding_problem_test_cases == nil {
				continue
			}
			fk := *nodes[i].coding_problem_test_cases
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(codingproblem.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coding_problem_test_cases" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CodingProblem = n
			}
		}
	}

	if query := ctcq.withData; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CodingTestCase)
		for i := range nodes {
			if nodes[i].coding_test_case_data_test_case == nil {
				continue
			}
			fk := *nodes[i].coding_test_case_data_test_case
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(codingtestcasedata.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coding_test_case_data_test_case" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Data = n
			}
		}
	}

	return nodes, nil
}

func (ctcq *CodingTestCaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctcq.querySpec()
	_spec.Node.Columns = ctcq.fields
	if len(ctcq.fields) > 0 {
		_spec.Unique = ctcq.unique != nil && *ctcq.unique
	}
	return sqlgraph.CountNodes(ctx, ctcq.driver, _spec)
}

func (ctcq *CodingTestCaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ctcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("generated: check existence: %w", err)
	}
	return n > 0, nil
}

func (ctcq *CodingTestCaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingtestcase.Table,
			Columns: codingtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingtestcase.FieldID,
			},
		},
		From:   ctcq.sql,
		Unique: true,
	}
	if unique := ctcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ctcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codingtestcase.FieldID)
		for i := range fields {
			if fields[i] != codingtestcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctcq *CodingTestCaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctcq.driver.Dialect())
	t1 := builder.Table(codingtestcase.Table)
	columns := ctcq.fields
	if len(columns) == 0 {
		columns = codingtestcase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctcq.sql != nil {
		selector = ctcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctcq.unique != nil && *ctcq.unique {
		selector.Distinct()
	}
	for _, p := range ctcq.predicates {
		p(selector)
	}
	for _, p := range ctcq.order {
		p(selector)
	}
	if offset := ctcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CodingTestCaseGroupBy is the group-by builder for CodingTestCase entities.
type CodingTestCaseGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctcgb *CodingTestCaseGroupBy) Aggregate(fns ...AggregateFunc) *CodingTestCaseGroupBy {
	ctcgb.fns = append(ctcgb.fns, fns...)
	return ctcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ctcgb *CodingTestCaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ctcgb.path(ctx)
	if err != nil {
		return err
	}
	ctcgb.sql = query
	return ctcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ctcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ctcgb.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ctcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) StringsX(ctx context.Context) []string {
	v, err := ctcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ctcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) StringX(ctx context.Context) string {
	v, err := ctcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ctcgb.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ctcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) IntsX(ctx context.Context) []int {
	v, err := ctcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ctcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) IntX(ctx context.Context) int {
	v, err := ctcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ctcgb.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ctcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ctcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ctcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ctcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ctcgb.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ctcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ctcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctcgb *CodingTestCaseGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ctcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ctcgb *CodingTestCaseGroupBy) BoolX(ctx context.Context) bool {
	v, err := ctcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ctcgb *CodingTestCaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ctcgb.fields {
		if !codingtestcase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ctcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ctcgb *CodingTestCaseGroupBy) sqlQuery() *sql.Selector {
	selector := ctcgb.sql.Select()
	aggregation := make([]string, 0, len(ctcgb.fns))
	for _, fn := range ctcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ctcgb.fields)+len(ctcgb.fns))
		for _, f := range ctcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ctcgb.fields...)...)
}

// CodingTestCaseSelect is the builder for selecting fields of CodingTestCase entities.
type CodingTestCaseSelect struct {
	*CodingTestCaseQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ctcs *CodingTestCaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ctcs.prepareQuery(ctx); err != nil {
		return err
	}
	ctcs.sql = ctcs.CodingTestCaseQuery.sqlQuery(ctx)
	return ctcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ctcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ctcs.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ctcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) StringsX(ctx context.Context) []string {
	v, err := ctcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ctcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) StringX(ctx context.Context) string {
	v, err := ctcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ctcs.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ctcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) IntsX(ctx context.Context) []int {
	v, err := ctcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ctcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) IntX(ctx context.Context) int {
	v, err := ctcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ctcs.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ctcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ctcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ctcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) Float64X(ctx context.Context) float64 {
	v, err := ctcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ctcs.fields) > 1 {
		return nil, errors.New("generated: CodingTestCaseSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ctcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) BoolsX(ctx context.Context) []bool {
	v, err := ctcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ctcs *CodingTestCaseSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ctcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingtestcase.Label}
	default:
		err = fmt.Errorf("generated: CodingTestCaseSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ctcs *CodingTestCaseSelect) BoolX(ctx context.Context) bool {
	v, err := ctcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ctcs *CodingTestCaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ctcs.sql.Query()
	if err := ctcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
