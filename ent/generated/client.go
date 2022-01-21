// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"log"

	"170-ag/ent/generated/migrate"

	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingproblemstaffdata"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingsubmissionstaffdata"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CodingDraft is the client for interacting with the CodingDraft builders.
	CodingDraft *CodingDraftClient
	// CodingProblem is the client for interacting with the CodingProblem builders.
	CodingProblem *CodingProblemClient
	// CodingProblemStaffData is the client for interacting with the CodingProblemStaffData builders.
	CodingProblemStaffData *CodingProblemStaffDataClient
	// CodingSubmission is the client for interacting with the CodingSubmission builders.
	CodingSubmission *CodingSubmissionClient
	// CodingSubmissionStaffData is the client for interacting with the CodingSubmissionStaffData builders.
	CodingSubmissionStaffData *CodingSubmissionStaffDataClient
	// CodingTestCase is the client for interacting with the CodingTestCase builders.
	CodingTestCase *CodingTestCaseClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CodingDraft = NewCodingDraftClient(c.config)
	c.CodingProblem = NewCodingProblemClient(c.config)
	c.CodingProblemStaffData = NewCodingProblemStaffDataClient(c.config)
	c.CodingSubmission = NewCodingSubmissionClient(c.config)
	c.CodingSubmissionStaffData = NewCodingSubmissionStaffDataClient(c.config)
	c.CodingTestCase = NewCodingTestCaseClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("generated: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                       ctx,
		config:                    cfg,
		CodingDraft:               NewCodingDraftClient(cfg),
		CodingProblem:             NewCodingProblemClient(cfg),
		CodingProblemStaffData:    NewCodingProblemStaffDataClient(cfg),
		CodingSubmission:          NewCodingSubmissionClient(cfg),
		CodingSubmissionStaffData: NewCodingSubmissionStaffDataClient(cfg),
		CodingTestCase:            NewCodingTestCaseClient(cfg),
		User:                      NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:                    cfg,
		CodingDraft:               NewCodingDraftClient(cfg),
		CodingProblem:             NewCodingProblemClient(cfg),
		CodingProblemStaffData:    NewCodingProblemStaffDataClient(cfg),
		CodingSubmission:          NewCodingSubmissionClient(cfg),
		CodingSubmissionStaffData: NewCodingSubmissionStaffDataClient(cfg),
		CodingTestCase:            NewCodingTestCaseClient(cfg),
		User:                      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CodingDraft.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CodingDraft.Use(hooks...)
	c.CodingProblem.Use(hooks...)
	c.CodingProblemStaffData.Use(hooks...)
	c.CodingSubmission.Use(hooks...)
	c.CodingSubmissionStaffData.Use(hooks...)
	c.CodingTestCase.Use(hooks...)
	c.User.Use(hooks...)
}

// CodingDraftClient is a client for the CodingDraft schema.
type CodingDraftClient struct {
	config
}

// NewCodingDraftClient returns a client for the CodingDraft from the given config.
func NewCodingDraftClient(c config) *CodingDraftClient {
	return &CodingDraftClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `codingdraft.Hooks(f(g(h())))`.
func (c *CodingDraftClient) Use(hooks ...Hook) {
	c.hooks.CodingDraft = append(c.hooks.CodingDraft, hooks...)
}

// Create returns a create builder for CodingDraft.
func (c *CodingDraftClient) Create() *CodingDraftCreate {
	mutation := newCodingDraftMutation(c.config, OpCreate)
	return &CodingDraftCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CodingDraft entities.
func (c *CodingDraftClient) CreateBulk(builders ...*CodingDraftCreate) *CodingDraftCreateBulk {
	return &CodingDraftCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CodingDraft.
func (c *CodingDraftClient) Update() *CodingDraftUpdate {
	mutation := newCodingDraftMutation(c.config, OpUpdate)
	return &CodingDraftUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CodingDraftClient) UpdateOne(cd *CodingDraft) *CodingDraftUpdateOne {
	mutation := newCodingDraftMutation(c.config, OpUpdateOne, withCodingDraft(cd))
	return &CodingDraftUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CodingDraftClient) UpdateOneID(id int) *CodingDraftUpdateOne {
	mutation := newCodingDraftMutation(c.config, OpUpdateOne, withCodingDraftID(id))
	return &CodingDraftUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CodingDraft.
func (c *CodingDraftClient) Delete() *CodingDraftDelete {
	mutation := newCodingDraftMutation(c.config, OpDelete)
	return &CodingDraftDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CodingDraftClient) DeleteOne(cd *CodingDraft) *CodingDraftDeleteOne {
	return c.DeleteOneID(cd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CodingDraftClient) DeleteOneID(id int) *CodingDraftDeleteOne {
	builder := c.Delete().Where(codingdraft.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CodingDraftDeleteOne{builder}
}

// Query returns a query builder for CodingDraft.
func (c *CodingDraftClient) Query() *CodingDraftQuery {
	return &CodingDraftQuery{
		config: c.config,
	}
}

// Get returns a CodingDraft entity by its id.
func (c *CodingDraftClient) Get(ctx context.Context, id int) (*CodingDraft, error) {
	return c.Query().Where(codingdraft.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CodingDraftClient) GetX(ctx context.Context, id int) *CodingDraft {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a CodingDraft.
func (c *CodingDraftClient) QueryAuthor(cd *CodingDraft) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingdraft.Table, codingdraft.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, codingdraft.AuthorTable, codingdraft.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(cd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCodingProblem queries the coding_problem edge of a CodingDraft.
func (c *CodingDraftClient) QueryCodingProblem(cd *CodingDraft) *CodingProblemQuery {
	query := &CodingProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingdraft.Table, codingdraft.FieldID, id),
			sqlgraph.To(codingproblem.Table, codingproblem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, codingdraft.CodingProblemTable, codingdraft.CodingProblemColumn),
		)
		fromV = sqlgraph.Neighbors(cd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CodingDraftClient) Hooks() []Hook {
	hooks := c.hooks.CodingDraft
	return append(hooks[:len(hooks):len(hooks)], codingdraft.Hooks[:]...)
}

// CodingProblemClient is a client for the CodingProblem schema.
type CodingProblemClient struct {
	config
}

// NewCodingProblemClient returns a client for the CodingProblem from the given config.
func NewCodingProblemClient(c config) *CodingProblemClient {
	return &CodingProblemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `codingproblem.Hooks(f(g(h())))`.
func (c *CodingProblemClient) Use(hooks ...Hook) {
	c.hooks.CodingProblem = append(c.hooks.CodingProblem, hooks...)
}

// Create returns a create builder for CodingProblem.
func (c *CodingProblemClient) Create() *CodingProblemCreate {
	mutation := newCodingProblemMutation(c.config, OpCreate)
	return &CodingProblemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CodingProblem entities.
func (c *CodingProblemClient) CreateBulk(builders ...*CodingProblemCreate) *CodingProblemCreateBulk {
	return &CodingProblemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CodingProblem.
func (c *CodingProblemClient) Update() *CodingProblemUpdate {
	mutation := newCodingProblemMutation(c.config, OpUpdate)
	return &CodingProblemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CodingProblemClient) UpdateOne(cp *CodingProblem) *CodingProblemUpdateOne {
	mutation := newCodingProblemMutation(c.config, OpUpdateOne, withCodingProblem(cp))
	return &CodingProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CodingProblemClient) UpdateOneID(id int) *CodingProblemUpdateOne {
	mutation := newCodingProblemMutation(c.config, OpUpdateOne, withCodingProblemID(id))
	return &CodingProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CodingProblem.
func (c *CodingProblemClient) Delete() *CodingProblemDelete {
	mutation := newCodingProblemMutation(c.config, OpDelete)
	return &CodingProblemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CodingProblemClient) DeleteOne(cp *CodingProblem) *CodingProblemDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CodingProblemClient) DeleteOneID(id int) *CodingProblemDeleteOne {
	builder := c.Delete().Where(codingproblem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CodingProblemDeleteOne{builder}
}

// Query returns a query builder for CodingProblem.
func (c *CodingProblemClient) Query() *CodingProblemQuery {
	return &CodingProblemQuery{
		config: c.config,
	}
}

// Get returns a CodingProblem entity by its id.
func (c *CodingProblemClient) Get(ctx context.Context, id int) (*CodingProblem, error) {
	return c.Query().Where(codingproblem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CodingProblemClient) GetX(ctx context.Context, id int) *CodingProblem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDrafts queries the drafts edge of a CodingProblem.
func (c *CodingProblemClient) QueryDrafts(cp *CodingProblem) *CodingDraftQuery {
	query := &CodingDraftQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblem.Table, codingproblem.FieldID, id),
			sqlgraph.To(codingdraft.Table, codingdraft.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, codingproblem.DraftsTable, codingproblem.DraftsColumn),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStaffData queries the staff_data edge of a CodingProblem.
func (c *CodingProblemClient) QueryStaffData(cp *CodingProblem) *CodingProblemStaffDataQuery {
	query := &CodingProblemStaffDataQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblem.Table, codingproblem.FieldID, id),
			sqlgraph.To(codingproblemstaffdata.Table, codingproblemstaffdata.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, codingproblem.StaffDataTable, codingproblem.StaffDataColumn),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTestCases queries the test_cases edge of a CodingProblem.
func (c *CodingProblemClient) QueryTestCases(cp *CodingProblem) *CodingTestCaseQuery {
	query := &CodingTestCaseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblem.Table, codingproblem.FieldID, id),
			sqlgraph.To(codingtestcase.Table, codingtestcase.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, codingproblem.TestCasesTable, codingproblem.TestCasesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubmissions queries the submissions edge of a CodingProblem.
func (c *CodingProblemClient) QuerySubmissions(cp *CodingProblem) *CodingSubmissionQuery {
	query := &CodingSubmissionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblem.Table, codingproblem.FieldID, id),
			sqlgraph.To(codingsubmission.Table, codingsubmission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, codingproblem.SubmissionsTable, codingproblem.SubmissionsColumn),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CodingProblemClient) Hooks() []Hook {
	hooks := c.hooks.CodingProblem
	return append(hooks[:len(hooks):len(hooks)], codingproblem.Hooks[:]...)
}

// CodingProblemStaffDataClient is a client for the CodingProblemStaffData schema.
type CodingProblemStaffDataClient struct {
	config
}

// NewCodingProblemStaffDataClient returns a client for the CodingProblemStaffData from the given config.
func NewCodingProblemStaffDataClient(c config) *CodingProblemStaffDataClient {
	return &CodingProblemStaffDataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `codingproblemstaffdata.Hooks(f(g(h())))`.
func (c *CodingProblemStaffDataClient) Use(hooks ...Hook) {
	c.hooks.CodingProblemStaffData = append(c.hooks.CodingProblemStaffData, hooks...)
}

// Create returns a create builder for CodingProblemStaffData.
func (c *CodingProblemStaffDataClient) Create() *CodingProblemStaffDataCreate {
	mutation := newCodingProblemStaffDataMutation(c.config, OpCreate)
	return &CodingProblemStaffDataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CodingProblemStaffData entities.
func (c *CodingProblemStaffDataClient) CreateBulk(builders ...*CodingProblemStaffDataCreate) *CodingProblemStaffDataCreateBulk {
	return &CodingProblemStaffDataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CodingProblemStaffData.
func (c *CodingProblemStaffDataClient) Update() *CodingProblemStaffDataUpdate {
	mutation := newCodingProblemStaffDataMutation(c.config, OpUpdate)
	return &CodingProblemStaffDataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CodingProblemStaffDataClient) UpdateOne(cpsd *CodingProblemStaffData) *CodingProblemStaffDataUpdateOne {
	mutation := newCodingProblemStaffDataMutation(c.config, OpUpdateOne, withCodingProblemStaffData(cpsd))
	return &CodingProblemStaffDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CodingProblemStaffDataClient) UpdateOneID(id int) *CodingProblemStaffDataUpdateOne {
	mutation := newCodingProblemStaffDataMutation(c.config, OpUpdateOne, withCodingProblemStaffDataID(id))
	return &CodingProblemStaffDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CodingProblemStaffData.
func (c *CodingProblemStaffDataClient) Delete() *CodingProblemStaffDataDelete {
	mutation := newCodingProblemStaffDataMutation(c.config, OpDelete)
	return &CodingProblemStaffDataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CodingProblemStaffDataClient) DeleteOne(cpsd *CodingProblemStaffData) *CodingProblemStaffDataDeleteOne {
	return c.DeleteOneID(cpsd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CodingProblemStaffDataClient) DeleteOneID(id int) *CodingProblemStaffDataDeleteOne {
	builder := c.Delete().Where(codingproblemstaffdata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CodingProblemStaffDataDeleteOne{builder}
}

// Query returns a query builder for CodingProblemStaffData.
func (c *CodingProblemStaffDataClient) Query() *CodingProblemStaffDataQuery {
	return &CodingProblemStaffDataQuery{
		config: c.config,
	}
}

// Get returns a CodingProblemStaffData entity by its id.
func (c *CodingProblemStaffDataClient) Get(ctx context.Context, id int) (*CodingProblemStaffData, error) {
	return c.Query().Where(codingproblemstaffdata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CodingProblemStaffDataClient) GetX(ctx context.Context, id int) *CodingProblemStaffData {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCodingProblem queries the coding_problem edge of a CodingProblemStaffData.
func (c *CodingProblemStaffDataClient) QueryCodingProblem(cpsd *CodingProblemStaffData) *CodingProblemQuery {
	query := &CodingProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cpsd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingproblemstaffdata.Table, codingproblemstaffdata.FieldID, id),
			sqlgraph.To(codingproblem.Table, codingproblem.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, codingproblemstaffdata.CodingProblemTable, codingproblemstaffdata.CodingProblemColumn),
		)
		fromV = sqlgraph.Neighbors(cpsd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CodingProblemStaffDataClient) Hooks() []Hook {
	return c.hooks.CodingProblemStaffData
}

// CodingSubmissionClient is a client for the CodingSubmission schema.
type CodingSubmissionClient struct {
	config
}

// NewCodingSubmissionClient returns a client for the CodingSubmission from the given config.
func NewCodingSubmissionClient(c config) *CodingSubmissionClient {
	return &CodingSubmissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `codingsubmission.Hooks(f(g(h())))`.
func (c *CodingSubmissionClient) Use(hooks ...Hook) {
	c.hooks.CodingSubmission = append(c.hooks.CodingSubmission, hooks...)
}

// Create returns a create builder for CodingSubmission.
func (c *CodingSubmissionClient) Create() *CodingSubmissionCreate {
	mutation := newCodingSubmissionMutation(c.config, OpCreate)
	return &CodingSubmissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CodingSubmission entities.
func (c *CodingSubmissionClient) CreateBulk(builders ...*CodingSubmissionCreate) *CodingSubmissionCreateBulk {
	return &CodingSubmissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CodingSubmission.
func (c *CodingSubmissionClient) Update() *CodingSubmissionUpdate {
	mutation := newCodingSubmissionMutation(c.config, OpUpdate)
	return &CodingSubmissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CodingSubmissionClient) UpdateOne(cs *CodingSubmission) *CodingSubmissionUpdateOne {
	mutation := newCodingSubmissionMutation(c.config, OpUpdateOne, withCodingSubmission(cs))
	return &CodingSubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CodingSubmissionClient) UpdateOneID(id int) *CodingSubmissionUpdateOne {
	mutation := newCodingSubmissionMutation(c.config, OpUpdateOne, withCodingSubmissionID(id))
	return &CodingSubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CodingSubmission.
func (c *CodingSubmissionClient) Delete() *CodingSubmissionDelete {
	mutation := newCodingSubmissionMutation(c.config, OpDelete)
	return &CodingSubmissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CodingSubmissionClient) DeleteOne(cs *CodingSubmission) *CodingSubmissionDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CodingSubmissionClient) DeleteOneID(id int) *CodingSubmissionDeleteOne {
	builder := c.Delete().Where(codingsubmission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CodingSubmissionDeleteOne{builder}
}

// Query returns a query builder for CodingSubmission.
func (c *CodingSubmissionClient) Query() *CodingSubmissionQuery {
	return &CodingSubmissionQuery{
		config: c.config,
	}
}

// Get returns a CodingSubmission entity by its id.
func (c *CodingSubmissionClient) Get(ctx context.Context, id int) (*CodingSubmission, error) {
	return c.Query().Where(codingsubmission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CodingSubmissionClient) GetX(ctx context.Context, id int) *CodingSubmission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a CodingSubmission.
func (c *CodingSubmissionClient) QueryAuthor(cs *CodingSubmission) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingsubmission.Table, codingsubmission.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, codingsubmission.AuthorTable, codingsubmission.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCodingProblem queries the coding_problem edge of a CodingSubmission.
func (c *CodingSubmissionClient) QueryCodingProblem(cs *CodingSubmission) *CodingProblemQuery {
	query := &CodingProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingsubmission.Table, codingsubmission.FieldID, id),
			sqlgraph.To(codingproblem.Table, codingproblem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, codingsubmission.CodingProblemTable, codingsubmission.CodingProblemColumn),
		)
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStaffData queries the staff_data edge of a CodingSubmission.
func (c *CodingSubmissionClient) QueryStaffData(cs *CodingSubmission) *CodingSubmissionStaffDataQuery {
	query := &CodingSubmissionStaffDataQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingsubmission.Table, codingsubmission.FieldID, id),
			sqlgraph.To(codingsubmissionstaffdata.Table, codingsubmissionstaffdata.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, codingsubmission.StaffDataTable, codingsubmission.StaffDataColumn),
		)
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CodingSubmissionClient) Hooks() []Hook {
	hooks := c.hooks.CodingSubmission
	return append(hooks[:len(hooks):len(hooks)], codingsubmission.Hooks[:]...)
}

// CodingSubmissionStaffDataClient is a client for the CodingSubmissionStaffData schema.
type CodingSubmissionStaffDataClient struct {
	config
}

// NewCodingSubmissionStaffDataClient returns a client for the CodingSubmissionStaffData from the given config.
func NewCodingSubmissionStaffDataClient(c config) *CodingSubmissionStaffDataClient {
	return &CodingSubmissionStaffDataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `codingsubmissionstaffdata.Hooks(f(g(h())))`.
func (c *CodingSubmissionStaffDataClient) Use(hooks ...Hook) {
	c.hooks.CodingSubmissionStaffData = append(c.hooks.CodingSubmissionStaffData, hooks...)
}

// Create returns a create builder for CodingSubmissionStaffData.
func (c *CodingSubmissionStaffDataClient) Create() *CodingSubmissionStaffDataCreate {
	mutation := newCodingSubmissionStaffDataMutation(c.config, OpCreate)
	return &CodingSubmissionStaffDataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CodingSubmissionStaffData entities.
func (c *CodingSubmissionStaffDataClient) CreateBulk(builders ...*CodingSubmissionStaffDataCreate) *CodingSubmissionStaffDataCreateBulk {
	return &CodingSubmissionStaffDataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CodingSubmissionStaffData.
func (c *CodingSubmissionStaffDataClient) Update() *CodingSubmissionStaffDataUpdate {
	mutation := newCodingSubmissionStaffDataMutation(c.config, OpUpdate)
	return &CodingSubmissionStaffDataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CodingSubmissionStaffDataClient) UpdateOne(cssd *CodingSubmissionStaffData) *CodingSubmissionStaffDataUpdateOne {
	mutation := newCodingSubmissionStaffDataMutation(c.config, OpUpdateOne, withCodingSubmissionStaffData(cssd))
	return &CodingSubmissionStaffDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CodingSubmissionStaffDataClient) UpdateOneID(id int) *CodingSubmissionStaffDataUpdateOne {
	mutation := newCodingSubmissionStaffDataMutation(c.config, OpUpdateOne, withCodingSubmissionStaffDataID(id))
	return &CodingSubmissionStaffDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CodingSubmissionStaffData.
func (c *CodingSubmissionStaffDataClient) Delete() *CodingSubmissionStaffDataDelete {
	mutation := newCodingSubmissionStaffDataMutation(c.config, OpDelete)
	return &CodingSubmissionStaffDataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CodingSubmissionStaffDataClient) DeleteOne(cssd *CodingSubmissionStaffData) *CodingSubmissionStaffDataDeleteOne {
	return c.DeleteOneID(cssd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CodingSubmissionStaffDataClient) DeleteOneID(id int) *CodingSubmissionStaffDataDeleteOne {
	builder := c.Delete().Where(codingsubmissionstaffdata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CodingSubmissionStaffDataDeleteOne{builder}
}

// Query returns a query builder for CodingSubmissionStaffData.
func (c *CodingSubmissionStaffDataClient) Query() *CodingSubmissionStaffDataQuery {
	return &CodingSubmissionStaffDataQuery{
		config: c.config,
	}
}

// Get returns a CodingSubmissionStaffData entity by its id.
func (c *CodingSubmissionStaffDataClient) Get(ctx context.Context, id int) (*CodingSubmissionStaffData, error) {
	return c.Query().Where(codingsubmissionstaffdata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CodingSubmissionStaffDataClient) GetX(ctx context.Context, id int) *CodingSubmissionStaffData {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCodingSubmission queries the coding_submission edge of a CodingSubmissionStaffData.
func (c *CodingSubmissionStaffDataClient) QueryCodingSubmission(cssd *CodingSubmissionStaffData) *CodingSubmissionQuery {
	query := &CodingSubmissionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cssd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingsubmissionstaffdata.Table, codingsubmissionstaffdata.FieldID, id),
			sqlgraph.To(codingsubmission.Table, codingsubmission.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, codingsubmissionstaffdata.CodingSubmissionTable, codingsubmissionstaffdata.CodingSubmissionColumn),
		)
		fromV = sqlgraph.Neighbors(cssd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CodingSubmissionStaffDataClient) Hooks() []Hook {
	hooks := c.hooks.CodingSubmissionStaffData
	return append(hooks[:len(hooks):len(hooks)], codingsubmissionstaffdata.Hooks[:]...)
}

// CodingTestCaseClient is a client for the CodingTestCase schema.
type CodingTestCaseClient struct {
	config
}

// NewCodingTestCaseClient returns a client for the CodingTestCase from the given config.
func NewCodingTestCaseClient(c config) *CodingTestCaseClient {
	return &CodingTestCaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `codingtestcase.Hooks(f(g(h())))`.
func (c *CodingTestCaseClient) Use(hooks ...Hook) {
	c.hooks.CodingTestCase = append(c.hooks.CodingTestCase, hooks...)
}

// Create returns a create builder for CodingTestCase.
func (c *CodingTestCaseClient) Create() *CodingTestCaseCreate {
	mutation := newCodingTestCaseMutation(c.config, OpCreate)
	return &CodingTestCaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CodingTestCase entities.
func (c *CodingTestCaseClient) CreateBulk(builders ...*CodingTestCaseCreate) *CodingTestCaseCreateBulk {
	return &CodingTestCaseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CodingTestCase.
func (c *CodingTestCaseClient) Update() *CodingTestCaseUpdate {
	mutation := newCodingTestCaseMutation(c.config, OpUpdate)
	return &CodingTestCaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CodingTestCaseClient) UpdateOne(ctc *CodingTestCase) *CodingTestCaseUpdateOne {
	mutation := newCodingTestCaseMutation(c.config, OpUpdateOne, withCodingTestCase(ctc))
	return &CodingTestCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CodingTestCaseClient) UpdateOneID(id int) *CodingTestCaseUpdateOne {
	mutation := newCodingTestCaseMutation(c.config, OpUpdateOne, withCodingTestCaseID(id))
	return &CodingTestCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CodingTestCase.
func (c *CodingTestCaseClient) Delete() *CodingTestCaseDelete {
	mutation := newCodingTestCaseMutation(c.config, OpDelete)
	return &CodingTestCaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CodingTestCaseClient) DeleteOne(ctc *CodingTestCase) *CodingTestCaseDeleteOne {
	return c.DeleteOneID(ctc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CodingTestCaseClient) DeleteOneID(id int) *CodingTestCaseDeleteOne {
	builder := c.Delete().Where(codingtestcase.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CodingTestCaseDeleteOne{builder}
}

// Query returns a query builder for CodingTestCase.
func (c *CodingTestCaseClient) Query() *CodingTestCaseQuery {
	return &CodingTestCaseQuery{
		config: c.config,
	}
}

// Get returns a CodingTestCase entity by its id.
func (c *CodingTestCaseClient) Get(ctx context.Context, id int) (*CodingTestCase, error) {
	return c.Query().Where(codingtestcase.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CodingTestCaseClient) GetX(ctx context.Context, id int) *CodingTestCase {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCodingProblem queries the coding_problem edge of a CodingTestCase.
func (c *CodingTestCaseClient) QueryCodingProblem(ctc *CodingTestCase) *CodingProblemQuery {
	query := &CodingProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ctc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(codingtestcase.Table, codingtestcase.FieldID, id),
			sqlgraph.To(codingproblem.Table, codingproblem.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, codingtestcase.CodingProblemTable, codingtestcase.CodingProblemPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ctc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CodingTestCaseClient) Hooks() []Hook {
	hooks := c.hooks.CodingTestCase
	return append(hooks[:len(hooks):len(hooks)], codingtestcase.Hooks[:]...)
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDrafts queries the drafts edge of a User.
func (c *UserClient) QueryDrafts(u *User) *CodingDraftQuery {
	query := &CodingDraftQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(codingdraft.Table, codingdraft.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.DraftsTable, user.DraftsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}
