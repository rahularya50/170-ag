// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingsubmissionstaffdata"
	"170-ag/ent/generated/predicate"
	"170-ag/ent/generated/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CodingSubmissionUpdate is the builder for updating CodingSubmission entities.
type CodingSubmissionUpdate struct {
	config
	hooks    []Hook
	mutation *CodingSubmissionMutation
}

// Where appends a list predicates to the CodingSubmissionUpdate builder.
func (csu *CodingSubmissionUpdate) Where(ps ...predicate.CodingSubmission) *CodingSubmissionUpdate {
	csu.mutation.Where(ps...)
	return csu
}

// SetStatus sets the "status" field.
func (csu *CodingSubmissionUpdate) SetStatus(c codingsubmission.Status) *CodingSubmissionUpdate {
	csu.mutation.SetStatus(c)
	return csu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (csu *CodingSubmissionUpdate) SetNillableStatus(c *codingsubmission.Status) *CodingSubmissionUpdate {
	if c != nil {
		csu.SetStatus(*c)
	}
	return csu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (csu *CodingSubmissionUpdate) SetAuthorID(id int) *CodingSubmissionUpdate {
	csu.mutation.SetAuthorID(id)
	return csu
}

// SetAuthor sets the "author" edge to the User entity.
func (csu *CodingSubmissionUpdate) SetAuthor(u *User) *CodingSubmissionUpdate {
	return csu.SetAuthorID(u.ID)
}

// SetCodingProblemID sets the "coding_problem" edge to the CodingProblem entity by ID.
func (csu *CodingSubmissionUpdate) SetCodingProblemID(id int) *CodingSubmissionUpdate {
	csu.mutation.SetCodingProblemID(id)
	return csu
}

// SetCodingProblem sets the "coding_problem" edge to the CodingProblem entity.
func (csu *CodingSubmissionUpdate) SetCodingProblem(c *CodingProblem) *CodingSubmissionUpdate {
	return csu.SetCodingProblemID(c.ID)
}

// SetStaffDataID sets the "staff_data" edge to the CodingSubmissionStaffData entity by ID.
func (csu *CodingSubmissionUpdate) SetStaffDataID(id int) *CodingSubmissionUpdate {
	csu.mutation.SetStaffDataID(id)
	return csu
}

// SetNillableStaffDataID sets the "staff_data" edge to the CodingSubmissionStaffData entity by ID if the given value is not nil.
func (csu *CodingSubmissionUpdate) SetNillableStaffDataID(id *int) *CodingSubmissionUpdate {
	if id != nil {
		csu = csu.SetStaffDataID(*id)
	}
	return csu
}

// SetStaffData sets the "staff_data" edge to the CodingSubmissionStaffData entity.
func (csu *CodingSubmissionUpdate) SetStaffData(c *CodingSubmissionStaffData) *CodingSubmissionUpdate {
	return csu.SetStaffDataID(c.ID)
}

// Mutation returns the CodingSubmissionMutation object of the builder.
func (csu *CodingSubmissionUpdate) Mutation() *CodingSubmissionMutation {
	return csu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (csu *CodingSubmissionUpdate) ClearAuthor() *CodingSubmissionUpdate {
	csu.mutation.ClearAuthor()
	return csu
}

// ClearCodingProblem clears the "coding_problem" edge to the CodingProblem entity.
func (csu *CodingSubmissionUpdate) ClearCodingProblem() *CodingSubmissionUpdate {
	csu.mutation.ClearCodingProblem()
	return csu
}

// ClearStaffData clears the "staff_data" edge to the CodingSubmissionStaffData entity.
func (csu *CodingSubmissionUpdate) ClearStaffData() *CodingSubmissionUpdate {
	csu.mutation.ClearStaffData()
	return csu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csu *CodingSubmissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(csu.hooks) == 0 {
		if err = csu.check(); err != nil {
			return 0, err
		}
		affected, err = csu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingSubmissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csu.check(); err != nil {
				return 0, err
			}
			csu.mutation = mutation
			affected, err = csu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csu.hooks) - 1; i >= 0; i-- {
			if csu.hooks[i] == nil {
				return 0, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = csu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (csu *CodingSubmissionUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *CodingSubmissionUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *CodingSubmissionUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csu *CodingSubmissionUpdate) check() error {
	if v, ok := csu.mutation.Status(); ok {
		if err := codingsubmission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "CodingSubmission.status": %w`, err)}
		}
	}
	if _, ok := csu.mutation.AuthorID(); csu.mutation.AuthorCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CodingSubmission.author"`)
	}
	if _, ok := csu.mutation.CodingProblemID(); csu.mutation.CodingProblemCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CodingSubmission.coding_problem"`)
	}
	return nil
}

func (csu *CodingSubmissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingsubmission.Table,
			Columns: codingsubmission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingsubmission.FieldID,
			},
		},
	}
	if ps := csu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: codingsubmission.FieldStatus,
		})
	}
	if csu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.AuthorTable,
			Columns: []string{codingsubmission.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.AuthorTable,
			Columns: []string{codingsubmission.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if csu.mutation.CodingProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.CodingProblemTable,
			Columns: []string{codingsubmission.CodingProblemColumn},
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
	if nodes := csu.mutation.CodingProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.CodingProblemTable,
			Columns: []string{codingsubmission.CodingProblemColumn},
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
	if csu.mutation.StaffDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingsubmission.StaffDataTable,
			Columns: []string{codingsubmission.StaffDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmissionstaffdata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.StaffDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingsubmission.StaffDataTable,
			Columns: []string{codingsubmission.StaffDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmissionstaffdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codingsubmission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CodingSubmissionUpdateOne is the builder for updating a single CodingSubmission entity.
type CodingSubmissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CodingSubmissionMutation
}

// SetStatus sets the "status" field.
func (csuo *CodingSubmissionUpdateOne) SetStatus(c codingsubmission.Status) *CodingSubmissionUpdateOne {
	csuo.mutation.SetStatus(c)
	return csuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (csuo *CodingSubmissionUpdateOne) SetNillableStatus(c *codingsubmission.Status) *CodingSubmissionUpdateOne {
	if c != nil {
		csuo.SetStatus(*c)
	}
	return csuo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (csuo *CodingSubmissionUpdateOne) SetAuthorID(id int) *CodingSubmissionUpdateOne {
	csuo.mutation.SetAuthorID(id)
	return csuo
}

// SetAuthor sets the "author" edge to the User entity.
func (csuo *CodingSubmissionUpdateOne) SetAuthor(u *User) *CodingSubmissionUpdateOne {
	return csuo.SetAuthorID(u.ID)
}

// SetCodingProblemID sets the "coding_problem" edge to the CodingProblem entity by ID.
func (csuo *CodingSubmissionUpdateOne) SetCodingProblemID(id int) *CodingSubmissionUpdateOne {
	csuo.mutation.SetCodingProblemID(id)
	return csuo
}

// SetCodingProblem sets the "coding_problem" edge to the CodingProblem entity.
func (csuo *CodingSubmissionUpdateOne) SetCodingProblem(c *CodingProblem) *CodingSubmissionUpdateOne {
	return csuo.SetCodingProblemID(c.ID)
}

// SetStaffDataID sets the "staff_data" edge to the CodingSubmissionStaffData entity by ID.
func (csuo *CodingSubmissionUpdateOne) SetStaffDataID(id int) *CodingSubmissionUpdateOne {
	csuo.mutation.SetStaffDataID(id)
	return csuo
}

// SetNillableStaffDataID sets the "staff_data" edge to the CodingSubmissionStaffData entity by ID if the given value is not nil.
func (csuo *CodingSubmissionUpdateOne) SetNillableStaffDataID(id *int) *CodingSubmissionUpdateOne {
	if id != nil {
		csuo = csuo.SetStaffDataID(*id)
	}
	return csuo
}

// SetStaffData sets the "staff_data" edge to the CodingSubmissionStaffData entity.
func (csuo *CodingSubmissionUpdateOne) SetStaffData(c *CodingSubmissionStaffData) *CodingSubmissionUpdateOne {
	return csuo.SetStaffDataID(c.ID)
}

// Mutation returns the CodingSubmissionMutation object of the builder.
func (csuo *CodingSubmissionUpdateOne) Mutation() *CodingSubmissionMutation {
	return csuo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (csuo *CodingSubmissionUpdateOne) ClearAuthor() *CodingSubmissionUpdateOne {
	csuo.mutation.ClearAuthor()
	return csuo
}

// ClearCodingProblem clears the "coding_problem" edge to the CodingProblem entity.
func (csuo *CodingSubmissionUpdateOne) ClearCodingProblem() *CodingSubmissionUpdateOne {
	csuo.mutation.ClearCodingProblem()
	return csuo
}

// ClearStaffData clears the "staff_data" edge to the CodingSubmissionStaffData entity.
func (csuo *CodingSubmissionUpdateOne) ClearStaffData() *CodingSubmissionUpdateOne {
	csuo.mutation.ClearStaffData()
	return csuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (csuo *CodingSubmissionUpdateOne) Select(field string, fields ...string) *CodingSubmissionUpdateOne {
	csuo.fields = append([]string{field}, fields...)
	return csuo
}

// Save executes the query and returns the updated CodingSubmission entity.
func (csuo *CodingSubmissionUpdateOne) Save(ctx context.Context) (*CodingSubmission, error) {
	var (
		err  error
		node *CodingSubmission
	)
	if len(csuo.hooks) == 0 {
		if err = csuo.check(); err != nil {
			return nil, err
		}
		node, err = csuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingSubmissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csuo.check(); err != nil {
				return nil, err
			}
			csuo.mutation = mutation
			node, err = csuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csuo.hooks) - 1; i >= 0; i-- {
			if csuo.hooks[i] == nil {
				return nil, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = csuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *CodingSubmissionUpdateOne) SaveX(ctx context.Context) *CodingSubmission {
	node, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csuo *CodingSubmissionUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *CodingSubmissionUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csuo *CodingSubmissionUpdateOne) check() error {
	if v, ok := csuo.mutation.Status(); ok {
		if err := codingsubmission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "CodingSubmission.status": %w`, err)}
		}
	}
	if _, ok := csuo.mutation.AuthorID(); csuo.mutation.AuthorCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CodingSubmission.author"`)
	}
	if _, ok := csuo.mutation.CodingProblemID(); csuo.mutation.CodingProblemCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CodingSubmission.coding_problem"`)
	}
	return nil
}

func (csuo *CodingSubmissionUpdateOne) sqlSave(ctx context.Context) (_node *CodingSubmission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingsubmission.Table,
			Columns: codingsubmission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingsubmission.FieldID,
			},
		},
	}
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "CodingSubmission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := csuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codingsubmission.FieldID)
		for _, f := range fields {
			if !codingsubmission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != codingsubmission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := csuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: codingsubmission.FieldStatus,
		})
	}
	if csuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.AuthorTable,
			Columns: []string{codingsubmission.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.AuthorTable,
			Columns: []string{codingsubmission.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if csuo.mutation.CodingProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.CodingProblemTable,
			Columns: []string{codingsubmission.CodingProblemColumn},
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
	if nodes := csuo.mutation.CodingProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.CodingProblemTable,
			Columns: []string{codingsubmission.CodingProblemColumn},
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
	if csuo.mutation.StaffDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingsubmission.StaffDataTable,
			Columns: []string{codingsubmission.StaffDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmissionstaffdata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.StaffDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingsubmission.StaffDataTable,
			Columns: []string{codingsubmission.StaffDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmissionstaffdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CodingSubmission{config: csuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codingsubmission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
