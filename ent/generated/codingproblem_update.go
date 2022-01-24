// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CodingProblemUpdate is the builder for updating CodingProblem entities.
type CodingProblemUpdate struct {
	config
	hooks    []Hook
	mutation *CodingProblemMutation
}

// Where appends a list predicates to the CodingProblemUpdate builder.
func (cpu *CodingProblemUpdate) Where(ps ...predicate.CodingProblem) *CodingProblemUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetUpdateTime sets the "update_time" field.
func (cpu *CodingProblemUpdate) SetUpdateTime(t time.Time) *CodingProblemUpdate {
	cpu.mutation.SetUpdateTime(t)
	return cpu
}

// SetName sets the "name" field.
func (cpu *CodingProblemUpdate) SetName(s string) *CodingProblemUpdate {
	cpu.mutation.SetName(s)
	return cpu
}

// SetStatement sets the "statement" field.
func (cpu *CodingProblemUpdate) SetStatement(s string) *CodingProblemUpdate {
	cpu.mutation.SetStatement(s)
	return cpu
}

// SetNillableStatement sets the "statement" field if the given value is not nil.
func (cpu *CodingProblemUpdate) SetNillableStatement(s *string) *CodingProblemUpdate {
	if s != nil {
		cpu.SetStatement(*s)
	}
	return cpu
}

// SetReleased sets the "released" field.
func (cpu *CodingProblemUpdate) SetReleased(b bool) *CodingProblemUpdate {
	cpu.mutation.SetReleased(b)
	return cpu
}

// SetNillableReleased sets the "released" field if the given value is not nil.
func (cpu *CodingProblemUpdate) SetNillableReleased(b *bool) *CodingProblemUpdate {
	if b != nil {
		cpu.SetReleased(*b)
	}
	return cpu
}

// AddDraftIDs adds the "drafts" edge to the CodingDraft entity by IDs.
func (cpu *CodingProblemUpdate) AddDraftIDs(ids ...int) *CodingProblemUpdate {
	cpu.mutation.AddDraftIDs(ids...)
	return cpu
}

// AddDrafts adds the "drafts" edges to the CodingDraft entity.
func (cpu *CodingProblemUpdate) AddDrafts(c ...*CodingDraft) *CodingProblemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpu.AddDraftIDs(ids...)
}

// AddTestCaseIDs adds the "test_cases" edge to the CodingTestCase entity by IDs.
func (cpu *CodingProblemUpdate) AddTestCaseIDs(ids ...int) *CodingProblemUpdate {
	cpu.mutation.AddTestCaseIDs(ids...)
	return cpu
}

// AddTestCases adds the "test_cases" edges to the CodingTestCase entity.
func (cpu *CodingProblemUpdate) AddTestCases(c ...*CodingTestCase) *CodingProblemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpu.AddTestCaseIDs(ids...)
}

// AddSubmissionIDs adds the "submissions" edge to the CodingSubmission entity by IDs.
func (cpu *CodingProblemUpdate) AddSubmissionIDs(ids ...int) *CodingProblemUpdate {
	cpu.mutation.AddSubmissionIDs(ids...)
	return cpu
}

// AddSubmissions adds the "submissions" edges to the CodingSubmission entity.
func (cpu *CodingProblemUpdate) AddSubmissions(c ...*CodingSubmission) *CodingProblemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpu.AddSubmissionIDs(ids...)
}

// Mutation returns the CodingProblemMutation object of the builder.
func (cpu *CodingProblemUpdate) Mutation() *CodingProblemMutation {
	return cpu.mutation
}

// ClearDrafts clears all "drafts" edges to the CodingDraft entity.
func (cpu *CodingProblemUpdate) ClearDrafts() *CodingProblemUpdate {
	cpu.mutation.ClearDrafts()
	return cpu
}

// RemoveDraftIDs removes the "drafts" edge to CodingDraft entities by IDs.
func (cpu *CodingProblemUpdate) RemoveDraftIDs(ids ...int) *CodingProblemUpdate {
	cpu.mutation.RemoveDraftIDs(ids...)
	return cpu
}

// RemoveDrafts removes "drafts" edges to CodingDraft entities.
func (cpu *CodingProblemUpdate) RemoveDrafts(c ...*CodingDraft) *CodingProblemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpu.RemoveDraftIDs(ids...)
}

// ClearTestCases clears all "test_cases" edges to the CodingTestCase entity.
func (cpu *CodingProblemUpdate) ClearTestCases() *CodingProblemUpdate {
	cpu.mutation.ClearTestCases()
	return cpu
}

// RemoveTestCaseIDs removes the "test_cases" edge to CodingTestCase entities by IDs.
func (cpu *CodingProblemUpdate) RemoveTestCaseIDs(ids ...int) *CodingProblemUpdate {
	cpu.mutation.RemoveTestCaseIDs(ids...)
	return cpu
}

// RemoveTestCases removes "test_cases" edges to CodingTestCase entities.
func (cpu *CodingProblemUpdate) RemoveTestCases(c ...*CodingTestCase) *CodingProblemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpu.RemoveTestCaseIDs(ids...)
}

// ClearSubmissions clears all "submissions" edges to the CodingSubmission entity.
func (cpu *CodingProblemUpdate) ClearSubmissions() *CodingProblemUpdate {
	cpu.mutation.ClearSubmissions()
	return cpu
}

// RemoveSubmissionIDs removes the "submissions" edge to CodingSubmission entities by IDs.
func (cpu *CodingProblemUpdate) RemoveSubmissionIDs(ids ...int) *CodingProblemUpdate {
	cpu.mutation.RemoveSubmissionIDs(ids...)
	return cpu
}

// RemoveSubmissions removes "submissions" edges to CodingSubmission entities.
func (cpu *CodingProblemUpdate) RemoveSubmissions(c ...*CodingSubmission) *CodingProblemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpu.RemoveSubmissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *CodingProblemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cpu.defaults(); err != nil {
		return 0, err
	}
	if len(cpu.hooks) == 0 {
		if err = cpu.check(); err != nil {
			return 0, err
		}
		affected, err = cpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingProblemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpu.check(); err != nil {
				return 0, err
			}
			cpu.mutation = mutation
			affected, err = cpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpu.hooks) - 1; i >= 0; i-- {
			if cpu.hooks[i] == nil {
				return 0, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = cpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *CodingProblemUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *CodingProblemUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *CodingProblemUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpu *CodingProblemUpdate) defaults() error {
	if _, ok := cpu.mutation.UpdateTime(); !ok {
		if codingproblem.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized codingproblem.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := codingproblem.UpdateDefaultUpdateTime()
		cpu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cpu *CodingProblemUpdate) check() error {
	if v, ok := cpu.mutation.Name(); ok {
		if err := codingproblem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "CodingProblem.name": %w`, err)}
		}
	}
	if v, ok := cpu.mutation.Statement(); ok {
		if err := codingproblem.StatementValidator(v); err != nil {
			return &ValidationError{Name: "statement", err: fmt.Errorf(`generated: validator failed for field "CodingProblem.statement": %w`, err)}
		}
	}
	return nil
}

func (cpu *CodingProblemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingproblem.Table,
			Columns: codingproblem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingproblem.FieldID,
			},
		},
	}
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: codingproblem.FieldUpdateTime,
		})
	}
	if value, ok := cpu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingproblem.FieldName,
		})
	}
	if value, ok := cpu.mutation.Statement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingproblem.FieldStatement,
		})
	}
	if value, ok := cpu.mutation.Released(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: codingproblem.FieldReleased,
		})
	}
	if cpu.mutation.DraftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.DraftsTable,
			Columns: []string{codingproblem.DraftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingdraft.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.RemovedDraftsIDs(); len(nodes) > 0 && !cpu.mutation.DraftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.DraftsTable,
			Columns: []string{codingproblem.DraftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingdraft.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.DraftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.DraftsTable,
			Columns: []string{codingproblem.DraftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingdraft.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpu.mutation.TestCasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   codingproblem.TestCasesTable,
			Columns: []string{codingproblem.TestCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcase.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.RemovedTestCasesIDs(); len(nodes) > 0 && !cpu.mutation.TestCasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   codingproblem.TestCasesTable,
			Columns: []string{codingproblem.TestCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcase.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.TestCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   codingproblem.TestCasesTable,
			Columns: []string{codingproblem.TestCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcase.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpu.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.SubmissionsTable,
			Columns: []string{codingproblem.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.RemovedSubmissionsIDs(); len(nodes) > 0 && !cpu.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.SubmissionsTable,
			Columns: []string{codingproblem.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.SubmissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.SubmissionsTable,
			Columns: []string{codingproblem.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codingproblem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CodingProblemUpdateOne is the builder for updating a single CodingProblem entity.
type CodingProblemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CodingProblemMutation
}

// SetUpdateTime sets the "update_time" field.
func (cpuo *CodingProblemUpdateOne) SetUpdateTime(t time.Time) *CodingProblemUpdateOne {
	cpuo.mutation.SetUpdateTime(t)
	return cpuo
}

// SetName sets the "name" field.
func (cpuo *CodingProblemUpdateOne) SetName(s string) *CodingProblemUpdateOne {
	cpuo.mutation.SetName(s)
	return cpuo
}

// SetStatement sets the "statement" field.
func (cpuo *CodingProblemUpdateOne) SetStatement(s string) *CodingProblemUpdateOne {
	cpuo.mutation.SetStatement(s)
	return cpuo
}

// SetNillableStatement sets the "statement" field if the given value is not nil.
func (cpuo *CodingProblemUpdateOne) SetNillableStatement(s *string) *CodingProblemUpdateOne {
	if s != nil {
		cpuo.SetStatement(*s)
	}
	return cpuo
}

// SetReleased sets the "released" field.
func (cpuo *CodingProblemUpdateOne) SetReleased(b bool) *CodingProblemUpdateOne {
	cpuo.mutation.SetReleased(b)
	return cpuo
}

// SetNillableReleased sets the "released" field if the given value is not nil.
func (cpuo *CodingProblemUpdateOne) SetNillableReleased(b *bool) *CodingProblemUpdateOne {
	if b != nil {
		cpuo.SetReleased(*b)
	}
	return cpuo
}

// AddDraftIDs adds the "drafts" edge to the CodingDraft entity by IDs.
func (cpuo *CodingProblemUpdateOne) AddDraftIDs(ids ...int) *CodingProblemUpdateOne {
	cpuo.mutation.AddDraftIDs(ids...)
	return cpuo
}

// AddDrafts adds the "drafts" edges to the CodingDraft entity.
func (cpuo *CodingProblemUpdateOne) AddDrafts(c ...*CodingDraft) *CodingProblemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpuo.AddDraftIDs(ids...)
}

// AddTestCaseIDs adds the "test_cases" edge to the CodingTestCase entity by IDs.
func (cpuo *CodingProblemUpdateOne) AddTestCaseIDs(ids ...int) *CodingProblemUpdateOne {
	cpuo.mutation.AddTestCaseIDs(ids...)
	return cpuo
}

// AddTestCases adds the "test_cases" edges to the CodingTestCase entity.
func (cpuo *CodingProblemUpdateOne) AddTestCases(c ...*CodingTestCase) *CodingProblemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpuo.AddTestCaseIDs(ids...)
}

// AddSubmissionIDs adds the "submissions" edge to the CodingSubmission entity by IDs.
func (cpuo *CodingProblemUpdateOne) AddSubmissionIDs(ids ...int) *CodingProblemUpdateOne {
	cpuo.mutation.AddSubmissionIDs(ids...)
	return cpuo
}

// AddSubmissions adds the "submissions" edges to the CodingSubmission entity.
func (cpuo *CodingProblemUpdateOne) AddSubmissions(c ...*CodingSubmission) *CodingProblemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpuo.AddSubmissionIDs(ids...)
}

// Mutation returns the CodingProblemMutation object of the builder.
func (cpuo *CodingProblemUpdateOne) Mutation() *CodingProblemMutation {
	return cpuo.mutation
}

// ClearDrafts clears all "drafts" edges to the CodingDraft entity.
func (cpuo *CodingProblemUpdateOne) ClearDrafts() *CodingProblemUpdateOne {
	cpuo.mutation.ClearDrafts()
	return cpuo
}

// RemoveDraftIDs removes the "drafts" edge to CodingDraft entities by IDs.
func (cpuo *CodingProblemUpdateOne) RemoveDraftIDs(ids ...int) *CodingProblemUpdateOne {
	cpuo.mutation.RemoveDraftIDs(ids...)
	return cpuo
}

// RemoveDrafts removes "drafts" edges to CodingDraft entities.
func (cpuo *CodingProblemUpdateOne) RemoveDrafts(c ...*CodingDraft) *CodingProblemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpuo.RemoveDraftIDs(ids...)
}

// ClearTestCases clears all "test_cases" edges to the CodingTestCase entity.
func (cpuo *CodingProblemUpdateOne) ClearTestCases() *CodingProblemUpdateOne {
	cpuo.mutation.ClearTestCases()
	return cpuo
}

// RemoveTestCaseIDs removes the "test_cases" edge to CodingTestCase entities by IDs.
func (cpuo *CodingProblemUpdateOne) RemoveTestCaseIDs(ids ...int) *CodingProblemUpdateOne {
	cpuo.mutation.RemoveTestCaseIDs(ids...)
	return cpuo
}

// RemoveTestCases removes "test_cases" edges to CodingTestCase entities.
func (cpuo *CodingProblemUpdateOne) RemoveTestCases(c ...*CodingTestCase) *CodingProblemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpuo.RemoveTestCaseIDs(ids...)
}

// ClearSubmissions clears all "submissions" edges to the CodingSubmission entity.
func (cpuo *CodingProblemUpdateOne) ClearSubmissions() *CodingProblemUpdateOne {
	cpuo.mutation.ClearSubmissions()
	return cpuo
}

// RemoveSubmissionIDs removes the "submissions" edge to CodingSubmission entities by IDs.
func (cpuo *CodingProblemUpdateOne) RemoveSubmissionIDs(ids ...int) *CodingProblemUpdateOne {
	cpuo.mutation.RemoveSubmissionIDs(ids...)
	return cpuo
}

// RemoveSubmissions removes "submissions" edges to CodingSubmission entities.
func (cpuo *CodingProblemUpdateOne) RemoveSubmissions(c ...*CodingSubmission) *CodingProblemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cpuo.RemoveSubmissionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *CodingProblemUpdateOne) Select(field string, fields ...string) *CodingProblemUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated CodingProblem entity.
func (cpuo *CodingProblemUpdateOne) Save(ctx context.Context) (*CodingProblem, error) {
	var (
		err  error
		node *CodingProblem
	)
	if err := cpuo.defaults(); err != nil {
		return nil, err
	}
	if len(cpuo.hooks) == 0 {
		if err = cpuo.check(); err != nil {
			return nil, err
		}
		node, err = cpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingProblemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpuo.check(); err != nil {
				return nil, err
			}
			cpuo.mutation = mutation
			node, err = cpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cpuo.hooks) - 1; i >= 0; i-- {
			if cpuo.hooks[i] == nil {
				return nil, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = cpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *CodingProblemUpdateOne) SaveX(ctx context.Context) *CodingProblem {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *CodingProblemUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *CodingProblemUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpuo *CodingProblemUpdateOne) defaults() error {
	if _, ok := cpuo.mutation.UpdateTime(); !ok {
		if codingproblem.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized codingproblem.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := codingproblem.UpdateDefaultUpdateTime()
		cpuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cpuo *CodingProblemUpdateOne) check() error {
	if v, ok := cpuo.mutation.Name(); ok {
		if err := codingproblem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "CodingProblem.name": %w`, err)}
		}
	}
	if v, ok := cpuo.mutation.Statement(); ok {
		if err := codingproblem.StatementValidator(v); err != nil {
			return &ValidationError{Name: "statement", err: fmt.Errorf(`generated: validator failed for field "CodingProblem.statement": %w`, err)}
		}
	}
	return nil
}

func (cpuo *CodingProblemUpdateOne) sqlSave(ctx context.Context) (_node *CodingProblem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codingproblem.Table,
			Columns: codingproblem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingproblem.FieldID,
			},
		},
	}
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "CodingProblem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codingproblem.FieldID)
		for _, f := range fields {
			if !codingproblem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != codingproblem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: codingproblem.FieldUpdateTime,
		})
	}
	if value, ok := cpuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingproblem.FieldName,
		})
	}
	if value, ok := cpuo.mutation.Statement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingproblem.FieldStatement,
		})
	}
	if value, ok := cpuo.mutation.Released(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: codingproblem.FieldReleased,
		})
	}
	if cpuo.mutation.DraftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.DraftsTable,
			Columns: []string{codingproblem.DraftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingdraft.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.RemovedDraftsIDs(); len(nodes) > 0 && !cpuo.mutation.DraftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.DraftsTable,
			Columns: []string{codingproblem.DraftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingdraft.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.DraftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.DraftsTable,
			Columns: []string{codingproblem.DraftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingdraft.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpuo.mutation.TestCasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   codingproblem.TestCasesTable,
			Columns: []string{codingproblem.TestCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcase.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.RemovedTestCasesIDs(); len(nodes) > 0 && !cpuo.mutation.TestCasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   codingproblem.TestCasesTable,
			Columns: []string{codingproblem.TestCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcase.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.TestCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   codingproblem.TestCasesTable,
			Columns: []string{codingproblem.TestCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingtestcase.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpuo.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.SubmissionsTable,
			Columns: []string{codingproblem.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.RemovedSubmissionsIDs(); len(nodes) > 0 && !cpuo.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.SubmissionsTable,
			Columns: []string{codingproblem.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.SubmissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.SubmissionsTable,
			Columns: []string{codingproblem.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: codingsubmission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CodingProblem{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codingproblem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
