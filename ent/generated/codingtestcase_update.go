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
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CodingTestCaseUpdate is the builder for updating CodingTestCase entities.
type CodingTestCaseUpdate struct {
	config
	hooks    []Hook
	mutation *CodingTestCaseMutation
}

// Where appends a list predicates to the CodingTestCaseUpdate builder.
func (ctcu *CodingTestCaseUpdate) Where(ps ...predicate.CodingTestCase) *CodingTestCaseUpdate {
	ctcu.mutation.Where(ps...)
	return ctcu
}

// SetUpdateTime sets the "update_time" field.
func (ctcu *CodingTestCaseUpdate) SetUpdateTime(t time.Time) *CodingTestCaseUpdate {
	ctcu.mutation.SetUpdateTime(t)
	return ctcu
}

// SetPoints sets the "points" field.
func (ctcu *CodingTestCaseUpdate) SetPoints(i int) *CodingTestCaseUpdate {
	ctcu.mutation.ResetPoints()
	ctcu.mutation.SetPoints(i)
	return ctcu
}

// SetNillablePoints sets the "points" field if the given value is not nil.
func (ctcu *CodingTestCaseUpdate) SetNillablePoints(i *int) *CodingTestCaseUpdate {
	if i != nil {
		ctcu.SetPoints(*i)
	}
	return ctcu
}

// AddPoints adds i to the "points" field.
func (ctcu *CodingTestCaseUpdate) AddPoints(i int) *CodingTestCaseUpdate {
	ctcu.mutation.AddPoints(i)
	return ctcu
}

// SetVisibility sets the "visibility" field.
func (ctcu *CodingTestCaseUpdate) SetVisibility(c codingtestcase.Visibility) *CodingTestCaseUpdate {
	ctcu.mutation.SetVisibility(c)
	return ctcu
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (ctcu *CodingTestCaseUpdate) SetNillableVisibility(c *codingtestcase.Visibility) *CodingTestCaseUpdate {
	if c != nil {
		ctcu.SetVisibility(*c)
	}
	return ctcu
}

// SetCodingProblemID sets the "coding_problem" edge to the CodingProblem entity by ID.
func (ctcu *CodingTestCaseUpdate) SetCodingProblemID(id int) *CodingTestCaseUpdate {
	ctcu.mutation.SetCodingProblemID(id)
	return ctcu
}

// SetCodingProblem sets the "coding_problem" edge to the CodingProblem entity.
func (ctcu *CodingTestCaseUpdate) SetCodingProblem(c *CodingProblem) *CodingTestCaseUpdate {
	return ctcu.SetCodingProblemID(c.ID)
}

// SetDataID sets the "data" edge to the CodingTestCaseData entity by ID.
func (ctcu *CodingTestCaseUpdate) SetDataID(id int) *CodingTestCaseUpdate {
	ctcu.mutation.SetDataID(id)
	return ctcu
}

// SetNillableDataID sets the "data" edge to the CodingTestCaseData entity by ID if the given value is not nil.
func (ctcu *CodingTestCaseUpdate) SetNillableDataID(id *int) *CodingTestCaseUpdate {
	if id != nil {
		ctcu = ctcu.SetDataID(*id)
	}
	return ctcu
}

// SetData sets the "data" edge to the CodingTestCaseData entity.
func (ctcu *CodingTestCaseUpdate) SetData(c *CodingTestCaseData) *CodingTestCaseUpdate {
	return ctcu.SetDataID(c.ID)
}

// Mutation returns the CodingTestCaseMutation object of the builder.
func (ctcu *CodingTestCaseUpdate) Mutation() *CodingTestCaseMutation {
	return ctcu.mutation
}

// ClearCodingProblem clears the "coding_problem" edge to the CodingProblem entity.
func (ctcu *CodingTestCaseUpdate) ClearCodingProblem() *CodingTestCaseUpdate {
	ctcu.mutation.ClearCodingProblem()
	return ctcu
}

// ClearData clears the "data" edge to the CodingTestCaseData entity.
func (ctcu *CodingTestCaseUpdate) ClearData() *CodingTestCaseUpdate {
	ctcu.mutation.ClearData()
	return ctcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctcu *CodingTestCaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ctcu.defaults(); err != nil {
		return 0, err
	}
	if len(ctcu.hooks) == 0 {
		if err = ctcu.check(); err != nil {
			return 0, err
		}
		affected, err = ctcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctcu.check(); err != nil {
				return 0, err
			}
			ctcu.mutation = mutation
			affected, err = ctcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctcu.hooks) - 1; i >= 0; i-- {
			if ctcu.hooks[i] == nil {
				return 0, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = ctcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctcu *CodingTestCaseUpdate) SaveX(ctx context.Context) int {
	affected, err := ctcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctcu *CodingTestCaseUpdate) Exec(ctx context.Context) error {
	_, err := ctcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcu *CodingTestCaseUpdate) ExecX(ctx context.Context) {
	if err := ctcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctcu *CodingTestCaseUpdate) defaults() error {
	if _, ok := ctcu.mutation.UpdateTime(); !ok {
		if codingtestcase.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized codingtestcase.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := codingtestcase.UpdateDefaultUpdateTime()
		ctcu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ctcu *CodingTestCaseUpdate) check() error {
	if v, ok := ctcu.mutation.Points(); ok {
		if err := codingtestcase.PointsValidator(v); err != nil {
			return &ValidationError{Name: "points", err: fmt.Errorf(`generated: validator failed for field "CodingTestCase.points": %w`, err)}
		}
	}
	if v, ok := ctcu.mutation.Visibility(); ok {
		if err := codingtestcase.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "CodingTestCase.visibility": %w`, err)}
		}
	}
	if _, ok := ctcu.mutation.CodingProblemID(); ctcu.mutation.CodingProblemCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CodingTestCase.coding_problem"`)
	}
	return nil
}

func (ctcu *CodingTestCaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingtestcase.Table,
			Columns: codingtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingtestcase.FieldID,
			},
		},
	}
	if ps := ctcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctcu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: codingtestcase.FieldUpdateTime,
		})
	}
	if value, ok := ctcu.mutation.Points(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: codingtestcase.FieldPoints,
		})
	}
	if value, ok := ctcu.mutation.AddedPoints(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: codingtestcase.FieldPoints,
		})
	}
	if value, ok := ctcu.mutation.Visibility(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: codingtestcase.FieldVisibility,
		})
	}
	if ctcu.mutation.CodingProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codingtestcase.CodingProblemTable,
			Columns: []string{codingtestcase.CodingProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingproblem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctcu.mutation.CodingProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codingtestcase.CodingProblemTable,
			Columns: []string{codingtestcase.CodingProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingproblem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctcu.mutation.DataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingtestcase.DataTable,
			Columns: []string{codingtestcase.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcasedata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctcu.mutation.DataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingtestcase.DataTable,
			Columns: []string{codingtestcase.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcasedata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codingtestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CodingTestCaseUpdateOne is the builder for updating a single CodingTestCase entity.
type CodingTestCaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CodingTestCaseMutation
}

// SetUpdateTime sets the "update_time" field.
func (ctcuo *CodingTestCaseUpdateOne) SetUpdateTime(t time.Time) *CodingTestCaseUpdateOne {
	ctcuo.mutation.SetUpdateTime(t)
	return ctcuo
}

// SetPoints sets the "points" field.
func (ctcuo *CodingTestCaseUpdateOne) SetPoints(i int) *CodingTestCaseUpdateOne {
	ctcuo.mutation.ResetPoints()
	ctcuo.mutation.SetPoints(i)
	return ctcuo
}

// SetNillablePoints sets the "points" field if the given value is not nil.
func (ctcuo *CodingTestCaseUpdateOne) SetNillablePoints(i *int) *CodingTestCaseUpdateOne {
	if i != nil {
		ctcuo.SetPoints(*i)
	}
	return ctcuo
}

// AddPoints adds i to the "points" field.
func (ctcuo *CodingTestCaseUpdateOne) AddPoints(i int) *CodingTestCaseUpdateOne {
	ctcuo.mutation.AddPoints(i)
	return ctcuo
}

// SetVisibility sets the "visibility" field.
func (ctcuo *CodingTestCaseUpdateOne) SetVisibility(c codingtestcase.Visibility) *CodingTestCaseUpdateOne {
	ctcuo.mutation.SetVisibility(c)
	return ctcuo
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (ctcuo *CodingTestCaseUpdateOne) SetNillableVisibility(c *codingtestcase.Visibility) *CodingTestCaseUpdateOne {
	if c != nil {
		ctcuo.SetVisibility(*c)
	}
	return ctcuo
}

// SetCodingProblemID sets the "coding_problem" edge to the CodingProblem entity by ID.
func (ctcuo *CodingTestCaseUpdateOne) SetCodingProblemID(id int) *CodingTestCaseUpdateOne {
	ctcuo.mutation.SetCodingProblemID(id)
	return ctcuo
}

// SetCodingProblem sets the "coding_problem" edge to the CodingProblem entity.
func (ctcuo *CodingTestCaseUpdateOne) SetCodingProblem(c *CodingProblem) *CodingTestCaseUpdateOne {
	return ctcuo.SetCodingProblemID(c.ID)
}

// SetDataID sets the "data" edge to the CodingTestCaseData entity by ID.
func (ctcuo *CodingTestCaseUpdateOne) SetDataID(id int) *CodingTestCaseUpdateOne {
	ctcuo.mutation.SetDataID(id)
	return ctcuo
}

// SetNillableDataID sets the "data" edge to the CodingTestCaseData entity by ID if the given value is not nil.
func (ctcuo *CodingTestCaseUpdateOne) SetNillableDataID(id *int) *CodingTestCaseUpdateOne {
	if id != nil {
		ctcuo = ctcuo.SetDataID(*id)
	}
	return ctcuo
}

// SetData sets the "data" edge to the CodingTestCaseData entity.
func (ctcuo *CodingTestCaseUpdateOne) SetData(c *CodingTestCaseData) *CodingTestCaseUpdateOne {
	return ctcuo.SetDataID(c.ID)
}

// Mutation returns the CodingTestCaseMutation object of the builder.
func (ctcuo *CodingTestCaseUpdateOne) Mutation() *CodingTestCaseMutation {
	return ctcuo.mutation
}

// ClearCodingProblem clears the "coding_problem" edge to the CodingProblem entity.
func (ctcuo *CodingTestCaseUpdateOne) ClearCodingProblem() *CodingTestCaseUpdateOne {
	ctcuo.mutation.ClearCodingProblem()
	return ctcuo
}

// ClearData clears the "data" edge to the CodingTestCaseData entity.
func (ctcuo *CodingTestCaseUpdateOne) ClearData() *CodingTestCaseUpdateOne {
	ctcuo.mutation.ClearData()
	return ctcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctcuo *CodingTestCaseUpdateOne) Select(field string, fields ...string) *CodingTestCaseUpdateOne {
	ctcuo.fields = append([]string{field}, fields...)
	return ctcuo
}

// Save executes the query and returns the updated CodingTestCase entity.
func (ctcuo *CodingTestCaseUpdateOne) Save(ctx context.Context) (*CodingTestCase, error) {
	var (
		err  error
		node *CodingTestCase
	)
	if err := ctcuo.defaults(); err != nil {
		return nil, err
	}
	if len(ctcuo.hooks) == 0 {
		if err = ctcuo.check(); err != nil {
			return nil, err
		}
		node, err = ctcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctcuo.check(); err != nil {
				return nil, err
			}
			ctcuo.mutation = mutation
			node, err = ctcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctcuo.hooks) - 1; i >= 0; i-- {
			if ctcuo.hooks[i] == nil {
				return nil, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = ctcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctcuo *CodingTestCaseUpdateOne) SaveX(ctx context.Context) *CodingTestCase {
	node, err := ctcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctcuo *CodingTestCaseUpdateOne) Exec(ctx context.Context) error {
	_, err := ctcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcuo *CodingTestCaseUpdateOne) ExecX(ctx context.Context) {
	if err := ctcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctcuo *CodingTestCaseUpdateOne) defaults() error {
	if _, ok := ctcuo.mutation.UpdateTime(); !ok {
		if codingtestcase.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized codingtestcase.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := codingtestcase.UpdateDefaultUpdateTime()
		ctcuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ctcuo *CodingTestCaseUpdateOne) check() error {
	if v, ok := ctcuo.mutation.Points(); ok {
		if err := codingtestcase.PointsValidator(v); err != nil {
			return &ValidationError{Name: "points", err: fmt.Errorf(`generated: validator failed for field "CodingTestCase.points": %w`, err)}
		}
	}
	if v, ok := ctcuo.mutation.Visibility(); ok {
		if err := codingtestcase.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "CodingTestCase.visibility": %w`, err)}
		}
	}
	if _, ok := ctcuo.mutation.CodingProblemID(); ctcuo.mutation.CodingProblemCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CodingTestCase.coding_problem"`)
	}
	return nil
}

func (ctcuo *CodingTestCaseUpdateOne) sqlSave(ctx context.Context) (_node *CodingTestCase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingtestcase.Table,
			Columns: codingtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingtestcase.FieldID,
			},
		},
	}
	id, ok := ctcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "CodingTestCase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codingtestcase.FieldID)
		for _, f := range fields {
			if !codingtestcase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != codingtestcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctcuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: codingtestcase.FieldUpdateTime,
		})
	}
	if value, ok := ctcuo.mutation.Points(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: codingtestcase.FieldPoints,
		})
	}
	if value, ok := ctcuo.mutation.AddedPoints(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: codingtestcase.FieldPoints,
		})
	}
	if value, ok := ctcuo.mutation.Visibility(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: codingtestcase.FieldVisibility,
		})
	}
	if ctcuo.mutation.CodingProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codingtestcase.CodingProblemTable,
			Columns: []string{codingtestcase.CodingProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingproblem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctcuo.mutation.CodingProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codingtestcase.CodingProblemTable,
			Columns: []string{codingtestcase.CodingProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingproblem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctcuo.mutation.DataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingtestcase.DataTable,
			Columns: []string{codingtestcase.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcasedata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctcuo.mutation.DataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingtestcase.DataTable,
			Columns: []string{codingtestcase.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcasedata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CodingTestCase{config: ctcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codingtestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
