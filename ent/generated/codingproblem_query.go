// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingtestcase"
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

// CodingProblemQuery is the builder for querying CodingProblem entities.
type CodingProblemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CodingProblem
	// eager-loading edges.
	withDrafts      *CodingDraftQuery
	withTestCases   *CodingTestCaseQuery
	withSubmissions *CodingSubmissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CodingProblemQuery builder.
func (cpq *CodingProblemQuery) Where(ps ...predicate.CodingProblem) *CodingProblemQuery {
	cpq.predicates = append(cpq.predicates, ps...)
	return cpq
}

// Limit adds a limit step to the query.
func (cpq *CodingProblemQuery) Limit(limit int) *CodingProblemQuery {
	cpq.limit = &limit
	return cpq
}

// Offset adds an offset step to the query.
func (cpq *CodingProblemQuery) Offset(offset int) *CodingProblemQuery {
	cpq.offset = &offset
	return cpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cpq *CodingProblemQuery) Unique(unique bool) *CodingProblemQuery {
	cpq.unique = &unique
	return cpq
}

// Order adds an order step to the query.
func (cpq *CodingProblemQuery) Order(o ...OrderFunc) *CodingProblemQuery {
	cpq.order = append(cpq.order, o...)
	return cpq
}

// QueryDrafts chains the current query on the "drafts" edge.
func (cpq *CodingProblemQuery) QueryDrafts() *CodingDraftQuery {
	query := &CodingDraftQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblem.Table, codingproblem.FieldID, selector),
			sqlgraph.To(codingdraft.Table, codingdraft.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, codingproblem.DraftsTable, codingproblem.DraftsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTestCases chains the current query on the "test_cases" edge.
func (cpq *CodingProblemQuery) QueryTestCases() *CodingTestCaseQuery {
	query := &CodingTestCaseQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblem.Table, codingproblem.FieldID, selector),
			sqlgraph.To(codingtestcase.Table, codingtestcase.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, codingproblem.TestCasesTable, codingproblem.TestCasesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubmissions chains the current query on the "submissions" edge.
func (cpq *CodingProblemQuery) QuerySubmissions() *CodingSubmissionQuery {
	query := &CodingSubmissionQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblem.Table, codingproblem.FieldID, selector),
			sqlgraph.To(codingsubmission.Table, codingsubmission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, codingproblem.SubmissionsTable, codingproblem.SubmissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CodingProblem entity from the query.
// Returns a *NotFoundError when no CodingProblem was found.
func (cpq *CodingProblemQuery) First(ctx context.Context) (*CodingProblem, error) {
	nodes, err := cpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{codingproblem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpq *CodingProblemQuery) FirstX(ctx context.Context) *CodingProblem {
	node, err := cpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CodingProblem ID from the query.
// Returns a *NotFoundError when no CodingProblem ID was found.
func (cpq *CodingProblemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{codingproblem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cpq *CodingProblemQuery) FirstIDX(ctx context.Context) int {
	id, err := cpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CodingProblem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CodingProblem entity is not found.
// Returns a *NotFoundError when no CodingProblem entities are found.
func (cpq *CodingProblemQuery) Only(ctx context.Context) (*CodingProblem, error) {
	nodes, err := cpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{codingproblem.Label}
	default:
		return nil, &NotSingularError{codingproblem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpq *CodingProblemQuery) OnlyX(ctx context.Context) *CodingProblem {
	node, err := cpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CodingProblem ID in the query.
// Returns a *NotSingularError when exactly one CodingProblem ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cpq *CodingProblemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = &NotSingularError{codingproblem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cpq *CodingProblemQuery) OnlyIDX(ctx context.Context) int {
	id, err := cpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CodingProblems.
func (cpq *CodingProblemQuery) All(ctx context.Context) ([]*CodingProblem, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpq *CodingProblemQuery) AllX(ctx context.Context) []*CodingProblem {
	nodes, err := cpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CodingProblem IDs.
func (cpq *CodingProblemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cpq.Select(codingproblem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpq *CodingProblemQuery) IDsX(ctx context.Context) []int {
	ids, err := cpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpq *CodingProblemQuery) Count(ctx context.Context) (int, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpq *CodingProblemQuery) CountX(ctx context.Context) int {
	count, err := cpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpq *CodingProblemQuery) Exist(ctx context.Context) (bool, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpq *CodingProblemQuery) ExistX(ctx context.Context) bool {
	exist, err := cpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CodingProblemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpq *CodingProblemQuery) Clone() *CodingProblemQuery {
	if cpq == nil {
		return nil
	}
	return &CodingProblemQuery{
		config:          cpq.config,
		limit:           cpq.limit,
		offset:          cpq.offset,
		order:           append([]OrderFunc{}, cpq.order...),
		predicates:      append([]predicate.CodingProblem{}, cpq.predicates...),
		withDrafts:      cpq.withDrafts.Clone(),
		withTestCases:   cpq.withTestCases.Clone(),
		withSubmissions: cpq.withSubmissions.Clone(),
		// clone intermediate query.
		sql:  cpq.sql.Clone(),
		path: cpq.path,
	}
}

// WithDrafts tells the query-builder to eager-load the nodes that are connected to
// the "drafts" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CodingProblemQuery) WithDrafts(opts ...func(*CodingDraftQuery)) *CodingProblemQuery {
	query := &CodingDraftQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withDrafts = query
	return cpq
}

// WithTestCases tells the query-builder to eager-load the nodes that are connected to
// the "test_cases" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CodingProblemQuery) WithTestCases(opts ...func(*CodingTestCaseQuery)) *CodingProblemQuery {
	query := &CodingTestCaseQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withTestCases = query
	return cpq
}

// WithSubmissions tells the query-builder to eager-load the nodes that are connected to
// the "submissions" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CodingProblemQuery) WithSubmissions(opts ...func(*CodingSubmissionQuery)) *CodingProblemQuery {
	query := &CodingSubmissionQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withSubmissions = query
	return cpq
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
//	client.CodingProblem.Query().
//		GroupBy(codingproblem.FieldName).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
//
func (cpq *CodingProblemQuery) GroupBy(field string, fields ...string) *CodingProblemGroupBy {
	group := &CodingProblemGroupBy{config: cpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cpq.sqlQuery(ctx), nil
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
//	client.CodingProblem.Query().
//		Select(codingproblem.FieldName).
//		Scan(ctx, &v)
//
func (cpq *CodingProblemQuery) Select(fields ...string) *CodingProblemSelect {
	cpq.fields = append(cpq.fields, fields...)
	return &CodingProblemSelect{CodingProblemQuery: cpq}
}

func (cpq *CodingProblemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cpq.fields {
		if !codingproblem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if cpq.path != nil {
		prev, err := cpq.path(ctx)
		if err != nil {
			return err
		}
		cpq.sql = prev
	}
	if codingproblem.Policy == nil {
		return errors.New("generated: uninitialized codingproblem.Policy (forgotten import generated/runtime?)")
	}
	if err := codingproblem.Policy.EvalQuery(ctx, cpq); err != nil {
		return err
	}
	return nil
}

func (cpq *CodingProblemQuery) sqlAll(ctx context.Context) ([]*CodingProblem, error) {
	var (
		nodes       = []*CodingProblem{}
		_spec       = cpq.querySpec()
		loadedTypes = [3]bool{
			cpq.withDrafts != nil,
			cpq.withTestCases != nil,
			cpq.withSubmissions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CodingProblem{config: cpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, cpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cpq.withDrafts; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*CodingProblem)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Drafts = []*CodingDraft{}
		}
		query.withFKs = true
		query.Where(predicate.CodingDraft(func(s *sql.Selector) {
			s.Where(sql.InValues(codingproblem.DraftsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.coding_draft_coding_problem
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "coding_draft_coding_problem" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coding_draft_coding_problem" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Drafts = append(node.Edges.Drafts, n)
		}
	}

	if query := cpq.withTestCases; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*CodingProblem)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TestCases = []*CodingTestCase{}
		}
		query.withFKs = true
		query.Where(predicate.CodingTestCase(func(s *sql.Selector) {
			s.Where(sql.InValues(codingproblem.TestCasesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.coding_problem_test_cases
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "coding_problem_test_cases" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coding_problem_test_cases" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.TestCases = append(node.Edges.TestCases, n)
		}
	}

	if query := cpq.withSubmissions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*CodingProblem)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Submissions = []*CodingSubmission{}
		}
		query.withFKs = true
		query.Where(predicate.CodingSubmission(func(s *sql.Selector) {
			s.Where(sql.InValues(codingproblem.SubmissionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.coding_submission_coding_problem
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "coding_submission_coding_problem" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coding_submission_coding_problem" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Submissions = append(node.Edges.Submissions, n)
		}
	}

	return nodes, nil
}

func (cpq *CodingProblemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cpq.querySpec()
	return sqlgraph.CountNodes(ctx, cpq.driver, _spec)
}

func (cpq *CodingProblemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("generated: check existence: %w", err)
	}
	return n > 0, nil
}

func (cpq *CodingProblemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingproblem.Table,
			Columns: codingproblem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingproblem.FieldID,
			},
		},
		From:   cpq.sql,
		Unique: true,
	}
	if unique := cpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codingproblem.FieldID)
		for i := range fields {
			if fields[i] != codingproblem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cpq *CodingProblemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cpq.driver.Dialect())
	t1 := builder.Table(codingproblem.Table)
	columns := cpq.fields
	if len(columns) == 0 {
		columns = codingproblem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cpq.sql != nil {
		selector = cpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range cpq.predicates {
		p(selector)
	}
	for _, p := range cpq.order {
		p(selector)
	}
	if offset := cpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CodingProblemGroupBy is the group-by builder for CodingProblem entities.
type CodingProblemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpgb *CodingProblemGroupBy) Aggregate(fns ...AggregateFunc) *CodingProblemGroupBy {
	cpgb.fns = append(cpgb.fns, fns...)
	return cpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cpgb *CodingProblemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cpgb.path(ctx)
	if err != nil {
		return err
	}
	cpgb.sql = query
	return cpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) StringsX(ctx context.Context) []string {
	v, err := cpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) StringX(ctx context.Context) string {
	v, err := cpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) IntsX(ctx context.Context) []int {
	v, err := cpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) IntX(ctx context.Context) int {
	v, err := cpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("generated: CodingProblemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CodingProblemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cpgb *CodingProblemGroupBy) BoolX(ctx context.Context) bool {
	v, err := cpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cpgb *CodingProblemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cpgb.fields {
		if !codingproblem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpgb *CodingProblemGroupBy) sqlQuery() *sql.Selector {
	selector := cpgb.sql.Select()
	aggregation := make([]string, 0, len(cpgb.fns))
	for _, fn := range cpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cpgb.fields)+len(cpgb.fns))
		for _, f := range cpgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cpgb.fields...)...)
}

// CodingProblemSelect is the builder for selecting fields of CodingProblem entities.
type CodingProblemSelect struct {
	*CodingProblemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cps *CodingProblemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cps.prepareQuery(ctx); err != nil {
		return err
	}
	cps.sql = cps.CodingProblemQuery.sqlQuery(ctx)
	return cps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cps *CodingProblemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("generated: CodingProblemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cps *CodingProblemSelect) StringsX(ctx context.Context) []string {
	v, err := cps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cps *CodingProblemSelect) StringX(ctx context.Context) string {
	v, err := cps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("generated: CodingProblemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cps *CodingProblemSelect) IntsX(ctx context.Context) []int {
	v, err := cps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cps *CodingProblemSelect) IntX(ctx context.Context) int {
	v, err := cps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("generated: CodingProblemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cps *CodingProblemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cps *CodingProblemSelect) Float64X(ctx context.Context) float64 {
	v, err := cps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("generated: CodingProblemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cps *CodingProblemSelect) BoolsX(ctx context.Context) []bool {
	v, err := cps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cps *CodingProblemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{codingproblem.Label}
	default:
		err = fmt.Errorf("generated: CodingProblemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cps *CodingProblemSelect) BoolX(ctx context.Context) bool {
	v, err := cps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cps *CodingProblemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cps.sql.Query()
	if err := cps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
