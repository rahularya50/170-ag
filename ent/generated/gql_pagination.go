// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/user"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// CodingDraftEdge is the edge representation of CodingDraft.
type CodingDraftEdge struct {
	Node   *CodingDraft `json:"node"`
	Cursor Cursor       `json:"cursor"`
}

// CodingDraftConnection is the connection containing edges to CodingDraft.
type CodingDraftConnection struct {
	Edges      []*CodingDraftEdge `json:"edges"`
	PageInfo   PageInfo           `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

// CodingDraftPaginateOption enables pagination customization.
type CodingDraftPaginateOption func(*codingDraftPager) error

// WithCodingDraftOrder configures pagination ordering.
func WithCodingDraftOrder(order *CodingDraftOrder) CodingDraftPaginateOption {
	if order == nil {
		order = DefaultCodingDraftOrder
	}
	o := *order
	return func(pager *codingDraftPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCodingDraftOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCodingDraftFilter configures pagination filter.
func WithCodingDraftFilter(filter func(*CodingDraftQuery) (*CodingDraftQuery, error)) CodingDraftPaginateOption {
	return func(pager *codingDraftPager) error {
		if filter == nil {
			return errors.New("CodingDraftQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type codingDraftPager struct {
	order  *CodingDraftOrder
	filter func(*CodingDraftQuery) (*CodingDraftQuery, error)
}

func newCodingDraftPager(opts []CodingDraftPaginateOption) (*codingDraftPager, error) {
	pager := &codingDraftPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCodingDraftOrder
	}
	return pager, nil
}

func (p *codingDraftPager) applyFilter(query *CodingDraftQuery) (*CodingDraftQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *codingDraftPager) toCursor(cd *CodingDraft) Cursor {
	return p.order.Field.toCursor(cd)
}

func (p *codingDraftPager) applyCursors(query *CodingDraftQuery, after, before *Cursor) *CodingDraftQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultCodingDraftOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *codingDraftPager) applyOrder(query *CodingDraftQuery, reverse bool) *CodingDraftQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultCodingDraftOrder.Field {
		query = query.Order(direction.orderFunc(DefaultCodingDraftOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to CodingDraft.
func (cd *CodingDraftQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CodingDraftPaginateOption,
) (*CodingDraftConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCodingDraftPager(opts)
	if err != nil {
		return nil, err
	}

	if cd, err = pager.applyFilter(cd); err != nil {
		return nil, err
	}

	conn := &CodingDraftConnection{Edges: []*CodingDraftEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := cd.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := cd.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	cd = pager.applyCursors(cd, after, before)
	cd = pager.applyOrder(cd, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		cd = cd.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		cd = cd.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := cd.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *CodingDraft
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *CodingDraft {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *CodingDraft {
			return nodes[i]
		}
	}

	conn.Edges = make([]*CodingDraftEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &CodingDraftEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// CodingDraftOrderField defines the ordering field of CodingDraft.
type CodingDraftOrderField struct {
	field    string
	toCursor func(*CodingDraft) Cursor
}

// CodingDraftOrder defines the ordering of CodingDraft.
type CodingDraftOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     *CodingDraftOrderField `json:"field"`
}

// DefaultCodingDraftOrder is the default ordering of CodingDraft.
var DefaultCodingDraftOrder = &CodingDraftOrder{
	Direction: OrderDirectionAsc,
	Field: &CodingDraftOrderField{
		field: codingdraft.FieldID,
		toCursor: func(cd *CodingDraft) Cursor {
			return Cursor{ID: cd.ID}
		},
	},
}

// ToEdge converts CodingDraft into CodingDraftEdge.
func (cd *CodingDraft) ToEdge(order *CodingDraftOrder) *CodingDraftEdge {
	if order == nil {
		order = DefaultCodingDraftOrder
	}
	return &CodingDraftEdge{
		Node:   cd,
		Cursor: order.Field.toCursor(cd),
	}
}

// CodingProblemEdge is the edge representation of CodingProblem.
type CodingProblemEdge struct {
	Node   *CodingProblem `json:"node"`
	Cursor Cursor         `json:"cursor"`
}

// CodingProblemConnection is the connection containing edges to CodingProblem.
type CodingProblemConnection struct {
	Edges      []*CodingProblemEdge `json:"edges"`
	PageInfo   PageInfo             `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

// CodingProblemPaginateOption enables pagination customization.
type CodingProblemPaginateOption func(*codingProblemPager) error

// WithCodingProblemOrder configures pagination ordering.
func WithCodingProblemOrder(order *CodingProblemOrder) CodingProblemPaginateOption {
	if order == nil {
		order = DefaultCodingProblemOrder
	}
	o := *order
	return func(pager *codingProblemPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCodingProblemOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCodingProblemFilter configures pagination filter.
func WithCodingProblemFilter(filter func(*CodingProblemQuery) (*CodingProblemQuery, error)) CodingProblemPaginateOption {
	return func(pager *codingProblemPager) error {
		if filter == nil {
			return errors.New("CodingProblemQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type codingProblemPager struct {
	order  *CodingProblemOrder
	filter func(*CodingProblemQuery) (*CodingProblemQuery, error)
}

func newCodingProblemPager(opts []CodingProblemPaginateOption) (*codingProblemPager, error) {
	pager := &codingProblemPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCodingProblemOrder
	}
	return pager, nil
}

func (p *codingProblemPager) applyFilter(query *CodingProblemQuery) (*CodingProblemQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *codingProblemPager) toCursor(cp *CodingProblem) Cursor {
	return p.order.Field.toCursor(cp)
}

func (p *codingProblemPager) applyCursors(query *CodingProblemQuery, after, before *Cursor) *CodingProblemQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultCodingProblemOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *codingProblemPager) applyOrder(query *CodingProblemQuery, reverse bool) *CodingProblemQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultCodingProblemOrder.Field {
		query = query.Order(direction.orderFunc(DefaultCodingProblemOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to CodingProblem.
func (cp *CodingProblemQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CodingProblemPaginateOption,
) (*CodingProblemConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCodingProblemPager(opts)
	if err != nil {
		return nil, err
	}

	if cp, err = pager.applyFilter(cp); err != nil {
		return nil, err
	}

	conn := &CodingProblemConnection{Edges: []*CodingProblemEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := cp.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := cp.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	cp = pager.applyCursors(cp, after, before)
	cp = pager.applyOrder(cp, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		cp = cp.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		cp = cp.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := cp.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *CodingProblem
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *CodingProblem {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *CodingProblem {
			return nodes[i]
		}
	}

	conn.Edges = make([]*CodingProblemEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &CodingProblemEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// CodingProblemOrderField defines the ordering field of CodingProblem.
type CodingProblemOrderField struct {
	field    string
	toCursor func(*CodingProblem) Cursor
}

// CodingProblemOrder defines the ordering of CodingProblem.
type CodingProblemOrder struct {
	Direction OrderDirection           `json:"direction"`
	Field     *CodingProblemOrderField `json:"field"`
}

// DefaultCodingProblemOrder is the default ordering of CodingProblem.
var DefaultCodingProblemOrder = &CodingProblemOrder{
	Direction: OrderDirectionAsc,
	Field: &CodingProblemOrderField{
		field: codingproblem.FieldID,
		toCursor: func(cp *CodingProblem) Cursor {
			return Cursor{ID: cp.ID}
		},
	},
}

// ToEdge converts CodingProblem into CodingProblemEdge.
func (cp *CodingProblem) ToEdge(order *CodingProblemOrder) *CodingProblemEdge {
	if order == nil {
		order = DefaultCodingProblemOrder
	}
	return &CodingProblemEdge{
		Node:   cp,
		Cursor: order.Field.toCursor(cp),
	}
}

// CodingSubmissionEdge is the edge representation of CodingSubmission.
type CodingSubmissionEdge struct {
	Node   *CodingSubmission `json:"node"`
	Cursor Cursor            `json:"cursor"`
}

// CodingSubmissionConnection is the connection containing edges to CodingSubmission.
type CodingSubmissionConnection struct {
	Edges      []*CodingSubmissionEdge `json:"edges"`
	PageInfo   PageInfo                `json:"pageInfo"`
	TotalCount int                     `json:"totalCount"`
}

// CodingSubmissionPaginateOption enables pagination customization.
type CodingSubmissionPaginateOption func(*codingSubmissionPager) error

// WithCodingSubmissionOrder configures pagination ordering.
func WithCodingSubmissionOrder(order *CodingSubmissionOrder) CodingSubmissionPaginateOption {
	if order == nil {
		order = DefaultCodingSubmissionOrder
	}
	o := *order
	return func(pager *codingSubmissionPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCodingSubmissionOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCodingSubmissionFilter configures pagination filter.
func WithCodingSubmissionFilter(filter func(*CodingSubmissionQuery) (*CodingSubmissionQuery, error)) CodingSubmissionPaginateOption {
	return func(pager *codingSubmissionPager) error {
		if filter == nil {
			return errors.New("CodingSubmissionQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type codingSubmissionPager struct {
	order  *CodingSubmissionOrder
	filter func(*CodingSubmissionQuery) (*CodingSubmissionQuery, error)
}

func newCodingSubmissionPager(opts []CodingSubmissionPaginateOption) (*codingSubmissionPager, error) {
	pager := &codingSubmissionPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCodingSubmissionOrder
	}
	return pager, nil
}

func (p *codingSubmissionPager) applyFilter(query *CodingSubmissionQuery) (*CodingSubmissionQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *codingSubmissionPager) toCursor(cs *CodingSubmission) Cursor {
	return p.order.Field.toCursor(cs)
}

func (p *codingSubmissionPager) applyCursors(query *CodingSubmissionQuery, after, before *Cursor) *CodingSubmissionQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultCodingSubmissionOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *codingSubmissionPager) applyOrder(query *CodingSubmissionQuery, reverse bool) *CodingSubmissionQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultCodingSubmissionOrder.Field {
		query = query.Order(direction.orderFunc(DefaultCodingSubmissionOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to CodingSubmission.
func (cs *CodingSubmissionQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CodingSubmissionPaginateOption,
) (*CodingSubmissionConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCodingSubmissionPager(opts)
	if err != nil {
		return nil, err
	}

	if cs, err = pager.applyFilter(cs); err != nil {
		return nil, err
	}

	conn := &CodingSubmissionConnection{Edges: []*CodingSubmissionEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := cs.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := cs.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	cs = pager.applyCursors(cs, after, before)
	cs = pager.applyOrder(cs, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		cs = cs.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		cs = cs.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := cs.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *CodingSubmission
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *CodingSubmission {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *CodingSubmission {
			return nodes[i]
		}
	}

	conn.Edges = make([]*CodingSubmissionEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &CodingSubmissionEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// CodingSubmissionOrderField defines the ordering field of CodingSubmission.
type CodingSubmissionOrderField struct {
	field    string
	toCursor func(*CodingSubmission) Cursor
}

// CodingSubmissionOrder defines the ordering of CodingSubmission.
type CodingSubmissionOrder struct {
	Direction OrderDirection              `json:"direction"`
	Field     *CodingSubmissionOrderField `json:"field"`
}

// DefaultCodingSubmissionOrder is the default ordering of CodingSubmission.
var DefaultCodingSubmissionOrder = &CodingSubmissionOrder{
	Direction: OrderDirectionAsc,
	Field: &CodingSubmissionOrderField{
		field: codingsubmission.FieldID,
		toCursor: func(cs *CodingSubmission) Cursor {
			return Cursor{ID: cs.ID}
		},
	},
}

// ToEdge converts CodingSubmission into CodingSubmissionEdge.
func (cs *CodingSubmission) ToEdge(order *CodingSubmissionOrder) *CodingSubmissionEdge {
	if order == nil {
		order = DefaultCodingSubmissionOrder
	}
	return &CodingSubmissionEdge{
		Node:   cs,
		Cursor: order.Field.toCursor(cs),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}

	conn := &UserConnection{Edges: []*UserEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := u.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := u.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		u = u.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}

	conn.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
