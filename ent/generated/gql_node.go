// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingsubmissionstaffdata"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/codingtestcasedata"
	"170-ag/ent/generated/user"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (cd *CodingDraft) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     cd.ID,
		Type:   "CodingDraft",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(cd.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cd.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cd.Code); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "code",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "author",
	}
	err = cd.QueryAuthor().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CodingProblem",
		Name: "coding_problem",
	}
	err = cd.QueryCodingProblem().
		Select(codingproblem.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (cp *CodingProblem) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     cp.ID,
		Type:   "CodingProblem",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(cp.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cp.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cp.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cp.Statement); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "statement",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cp.Skeleton); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "skeleton",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cp.Released); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "bool",
		Name:  "released",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cp.Deadline); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "time.Time",
		Name:  "deadline",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CodingDraft",
		Name: "drafts",
	}
	err = cp.QueryDrafts().
		Select(codingdraft.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CodingTestCase",
		Name: "test_cases",
	}
	err = cp.QueryTestCases().
		Select(codingtestcase.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "CodingSubmission",
		Name: "submissions",
	}
	err = cp.QuerySubmissions().
		Select(codingsubmission.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (cs *CodingSubmission) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     cs.ID,
		Type:   "CodingSubmission",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(cs.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cs.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cs.Code); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cs.Status); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "codingsubmission.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cs.Points); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "int",
		Name:  "points",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cs.Results); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "models.CodingSubmissionResults",
		Name:  "results",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "author",
	}
	err = cs.QueryAuthor().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CodingProblem",
		Name: "coding_problem",
	}
	err = cs.QueryCodingProblem().
		Select(codingproblem.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "CodingSubmissionStaffData",
		Name: "staff_data",
	}
	err = cs.QueryStaffData().
		Select(codingsubmissionstaffdata.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (cssd *CodingSubmissionStaffData) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     cssd.ID,
		Type:   "CodingSubmissionStaffData",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(cssd.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cssd.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cssd.ExecutionID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int64",
		Name:  "execution_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cssd.Input); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "input",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cssd.Output); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "output",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cssd.Stderr); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "stderr",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cssd.ExitError); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "exit_error",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CodingSubmission",
		Name: "coding_submission",
	}
	err = cssd.QueryCodingSubmission().
		Select(codingsubmission.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ctc *CodingTestCase) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ctc.ID,
		Type:   "CodingTestCase",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(ctc.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ctc.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ctc.Points); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "points",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ctc.Visibility); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "codingtestcase.Visibility",
		Name:  "visibility",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CodingProblem",
		Name: "coding_problem",
	}
	err = ctc.QueryCodingProblem().
		Select(codingproblem.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CodingTestCaseData",
		Name: "data",
	}
	err = ctc.QueryData().
		Select(codingtestcasedata.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ctcd *CodingTestCaseData) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ctcd.ID,
		Type:   "CodingTestCaseData",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(ctcd.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ctcd.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ctcd.Input); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "input",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ctcd.Output); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "output",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CodingTestCase",
		Name: "test_case",
	}
	err = ctcd.QueryTestCase().
		Select(codingtestcase.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(u.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Email); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.IsStaff); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "bool",
		Name:  "is_staff",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CodingDraft",
		Name: "drafts",
	}
	err = u.QueryDrafts().
		Select(codingdraft.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CodingSubmission",
		Name: "submissions",
	}
	err = u.QuerySubmissions().
		Select(codingsubmission.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case codingdraft.Table:
		n, err := c.CodingDraft.Query().
			Where(codingdraft.ID(id)).
			CollectFields(ctx, "CodingDraft").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case codingproblem.Table:
		n, err := c.CodingProblem.Query().
			Where(codingproblem.ID(id)).
			CollectFields(ctx, "CodingProblem").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case codingsubmission.Table:
		n, err := c.CodingSubmission.Query().
			Where(codingsubmission.ID(id)).
			CollectFields(ctx, "CodingSubmission").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case codingsubmissionstaffdata.Table:
		n, err := c.CodingSubmissionStaffData.Query().
			Where(codingsubmissionstaffdata.ID(id)).
			CollectFields(ctx, "CodingSubmissionStaffData").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case codingtestcase.Table:
		n, err := c.CodingTestCase.Query().
			Where(codingtestcase.ID(id)).
			CollectFields(ctx, "CodingTestCase").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case codingtestcasedata.Table:
		n, err := c.CodingTestCaseData.Query().
			Where(codingtestcasedata.ID(id)).
			CollectFields(ctx, "CodingTestCaseData").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		n, err := c.User.Query().
			Where(user.ID(id)).
			CollectFields(ctx, "User").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case codingdraft.Table:
		nodes, err := c.CodingDraft.Query().
			Where(codingdraft.IDIn(ids...)).
			CollectFields(ctx, "CodingDraft").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case codingproblem.Table:
		nodes, err := c.CodingProblem.Query().
			Where(codingproblem.IDIn(ids...)).
			CollectFields(ctx, "CodingProblem").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case codingsubmission.Table:
		nodes, err := c.CodingSubmission.Query().
			Where(codingsubmission.IDIn(ids...)).
			CollectFields(ctx, "CodingSubmission").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case codingsubmissionstaffdata.Table:
		nodes, err := c.CodingSubmissionStaffData.Query().
			Where(codingsubmissionstaffdata.IDIn(ids...)).
			CollectFields(ctx, "CodingSubmissionStaffData").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case codingtestcase.Table:
		nodes, err := c.CodingTestCase.Query().
			Where(codingtestcase.IDIn(ids...)).
			CollectFields(ctx, "CodingTestCase").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case codingtestcasedata.Table:
		nodes, err := c.CodingTestCaseData.Query().
			Where(codingtestcasedata.IDIn(ids...)).
			CollectFields(ctx, "CodingTestCaseData").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		nodes, err := c.User.Query().
			Where(user.IDIn(ids...)).
			CollectFields(ctx, "User").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
