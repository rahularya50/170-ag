// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/predicate"
	"170-ag/ent/generated/projectscore"
	"170-ag/ent/generated/projectteam"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectScoreUpdate is the builder for updating ProjectScore entities.
type ProjectScoreUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectScoreMutation
}

// Where appends a list predicates to the ProjectScoreUpdate builder.
func (psu *ProjectScoreUpdate) Where(ps ...predicate.ProjectScore) *ProjectScoreUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetUpdateTime sets the "update_time" field.
func (psu *ProjectScoreUpdate) SetUpdateTime(t time.Time) *ProjectScoreUpdate {
	psu.mutation.SetUpdateTime(t)
	return psu
}

// SetCaseID sets the "case_id" field.
func (psu *ProjectScoreUpdate) SetCaseID(i int32) *ProjectScoreUpdate {
	psu.mutation.ResetCaseID()
	psu.mutation.SetCaseID(i)
	return psu
}

// AddCaseID adds i to the "case_id" field.
func (psu *ProjectScoreUpdate) AddCaseID(i int32) *ProjectScoreUpdate {
	psu.mutation.AddCaseID(i)
	return psu
}

// SetScore sets the "score" field.
func (psu *ProjectScoreUpdate) SetScore(f float64) *ProjectScoreUpdate {
	psu.mutation.ResetScore()
	psu.mutation.SetScore(f)
	return psu
}

// AddScore adds f to the "score" field.
func (psu *ProjectScoreUpdate) AddScore(f float64) *ProjectScoreUpdate {
	psu.mutation.AddScore(f)
	return psu
}

// SetType sets the "type" field.
func (psu *ProjectScoreUpdate) SetType(pr projectscore.Type) *ProjectScoreUpdate {
	psu.mutation.SetType(pr)
	return psu
}

// SetTeamID sets the "team" edge to the ProjectTeam entity by ID.
func (psu *ProjectScoreUpdate) SetTeamID(id int) *ProjectScoreUpdate {
	psu.mutation.SetTeamID(id)
	return psu
}

// SetNillableTeamID sets the "team" edge to the ProjectTeam entity by ID if the given value is not nil.
func (psu *ProjectScoreUpdate) SetNillableTeamID(id *int) *ProjectScoreUpdate {
	if id != nil {
		psu = psu.SetTeamID(*id)
	}
	return psu
}

// SetTeam sets the "team" edge to the ProjectTeam entity.
func (psu *ProjectScoreUpdate) SetTeam(p *ProjectTeam) *ProjectScoreUpdate {
	return psu.SetTeamID(p.ID)
}

// Mutation returns the ProjectScoreMutation object of the builder.
func (psu *ProjectScoreUpdate) Mutation() *ProjectScoreMutation {
	return psu.mutation
}

// ClearTeam clears the "team" edge to the ProjectTeam entity.
func (psu *ProjectScoreUpdate) ClearTeam() *ProjectScoreUpdate {
	psu.mutation.ClearTeam()
	return psu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *ProjectScoreUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := psu.defaults(); err != nil {
		return 0, err
	}
	if len(psu.hooks) == 0 {
		if err = psu.check(); err != nil {
			return 0, err
		}
		affected, err = psu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectScoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psu.check(); err != nil {
				return 0, err
			}
			psu.mutation = mutation
			affected, err = psu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psu.hooks) - 1; i >= 0; i-- {
			if psu.hooks[i] == nil {
				return 0, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = psu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (psu *ProjectScoreUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *ProjectScoreUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *ProjectScoreUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psu *ProjectScoreUpdate) defaults() error {
	if _, ok := psu.mutation.UpdateTime(); !ok {
		if projectscore.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized projectscore.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := projectscore.UpdateDefaultUpdateTime()
		psu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (psu *ProjectScoreUpdate) check() error {
	if v, ok := psu.mutation.GetType(); ok {
		if err := projectscore.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`generated: validator failed for field "ProjectScore.type": %w`, err)}
		}
	}
	return nil
}

func (psu *ProjectScoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectscore.Table,
			Columns: projectscore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectscore.FieldID,
			},
		},
	}
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projectscore.FieldUpdateTime,
		})
	}
	if value, ok := psu.mutation.CaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: projectscore.FieldCaseID,
		})
	}
	if value, ok := psu.mutation.AddedCaseID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: projectscore.FieldCaseID,
		})
	}
	if value, ok := psu.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: projectscore.FieldScore,
		})
	}
	if value, ok := psu.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: projectscore.FieldScore,
		})
	}
	if value, ok := psu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projectscore.FieldType,
		})
	}
	if psu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectscore.TeamTable,
			Columns: []string{projectscore.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectteam.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectscore.TeamTable,
			Columns: []string{projectscore.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectScoreUpdateOne is the builder for updating a single ProjectScore entity.
type ProjectScoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectScoreMutation
}

// SetUpdateTime sets the "update_time" field.
func (psuo *ProjectScoreUpdateOne) SetUpdateTime(t time.Time) *ProjectScoreUpdateOne {
	psuo.mutation.SetUpdateTime(t)
	return psuo
}

// SetCaseID sets the "case_id" field.
func (psuo *ProjectScoreUpdateOne) SetCaseID(i int32) *ProjectScoreUpdateOne {
	psuo.mutation.ResetCaseID()
	psuo.mutation.SetCaseID(i)
	return psuo
}

// AddCaseID adds i to the "case_id" field.
func (psuo *ProjectScoreUpdateOne) AddCaseID(i int32) *ProjectScoreUpdateOne {
	psuo.mutation.AddCaseID(i)
	return psuo
}

// SetScore sets the "score" field.
func (psuo *ProjectScoreUpdateOne) SetScore(f float64) *ProjectScoreUpdateOne {
	psuo.mutation.ResetScore()
	psuo.mutation.SetScore(f)
	return psuo
}

// AddScore adds f to the "score" field.
func (psuo *ProjectScoreUpdateOne) AddScore(f float64) *ProjectScoreUpdateOne {
	psuo.mutation.AddScore(f)
	return psuo
}

// SetType sets the "type" field.
func (psuo *ProjectScoreUpdateOne) SetType(pr projectscore.Type) *ProjectScoreUpdateOne {
	psuo.mutation.SetType(pr)
	return psuo
}

// SetTeamID sets the "team" edge to the ProjectTeam entity by ID.
func (psuo *ProjectScoreUpdateOne) SetTeamID(id int) *ProjectScoreUpdateOne {
	psuo.mutation.SetTeamID(id)
	return psuo
}

// SetNillableTeamID sets the "team" edge to the ProjectTeam entity by ID if the given value is not nil.
func (psuo *ProjectScoreUpdateOne) SetNillableTeamID(id *int) *ProjectScoreUpdateOne {
	if id != nil {
		psuo = psuo.SetTeamID(*id)
	}
	return psuo
}

// SetTeam sets the "team" edge to the ProjectTeam entity.
func (psuo *ProjectScoreUpdateOne) SetTeam(p *ProjectTeam) *ProjectScoreUpdateOne {
	return psuo.SetTeamID(p.ID)
}

// Mutation returns the ProjectScoreMutation object of the builder.
func (psuo *ProjectScoreUpdateOne) Mutation() *ProjectScoreMutation {
	return psuo.mutation
}

// ClearTeam clears the "team" edge to the ProjectTeam entity.
func (psuo *ProjectScoreUpdateOne) ClearTeam() *ProjectScoreUpdateOne {
	psuo.mutation.ClearTeam()
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *ProjectScoreUpdateOne) Select(field string, fields ...string) *ProjectScoreUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated ProjectScore entity.
func (psuo *ProjectScoreUpdateOne) Save(ctx context.Context) (*ProjectScore, error) {
	var (
		err  error
		node *ProjectScore
	)
	if err := psuo.defaults(); err != nil {
		return nil, err
	}
	if len(psuo.hooks) == 0 {
		if err = psuo.check(); err != nil {
			return nil, err
		}
		node, err = psuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectScoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psuo.check(); err != nil {
				return nil, err
			}
			psuo.mutation = mutation
			node, err = psuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psuo.hooks) - 1; i >= 0; i-- {
			if psuo.hooks[i] == nil {
				return nil, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = psuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *ProjectScoreUpdateOne) SaveX(ctx context.Context) *ProjectScore {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *ProjectScoreUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *ProjectScoreUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psuo *ProjectScoreUpdateOne) defaults() error {
	if _, ok := psuo.mutation.UpdateTime(); !ok {
		if projectscore.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized projectscore.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := projectscore.UpdateDefaultUpdateTime()
		psuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (psuo *ProjectScoreUpdateOne) check() error {
	if v, ok := psuo.mutation.GetType(); ok {
		if err := projectscore.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`generated: validator failed for field "ProjectScore.type": %w`, err)}
		}
	}
	return nil
}

func (psuo *ProjectScoreUpdateOne) sqlSave(ctx context.Context) (_node *ProjectScore, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectscore.Table,
			Columns: projectscore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectscore.FieldID,
			},
		},
	}
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "ProjectScore.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectscore.FieldID)
		for _, f := range fields {
			if !projectscore.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != projectscore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projectscore.FieldUpdateTime,
		})
	}
	if value, ok := psuo.mutation.CaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: projectscore.FieldCaseID,
		})
	}
	if value, ok := psuo.mutation.AddedCaseID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: projectscore.FieldCaseID,
		})
	}
	if value, ok := psuo.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: projectscore.FieldScore,
		})
	}
	if value, ok := psuo.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: projectscore.FieldScore,
		})
	}
	if value, ok := psuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projectscore.FieldType,
		})
	}
	if psuo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectscore.TeamTable,
			Columns: []string{projectscore.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectteam.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectscore.TeamTable,
			Columns: []string{projectscore.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectScore{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
