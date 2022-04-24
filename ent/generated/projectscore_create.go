// Code generated by entc, DO NOT EDIT.

package generated

import (
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

// ProjectScoreCreate is the builder for creating a ProjectScore entity.
type ProjectScoreCreate struct {
	config
	mutation *ProjectScoreMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (psc *ProjectScoreCreate) SetCreateTime(t time.Time) *ProjectScoreCreate {
	psc.mutation.SetCreateTime(t)
	return psc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (psc *ProjectScoreCreate) SetNillableCreateTime(t *time.Time) *ProjectScoreCreate {
	if t != nil {
		psc.SetCreateTime(*t)
	}
	return psc
}

// SetUpdateTime sets the "update_time" field.
func (psc *ProjectScoreCreate) SetUpdateTime(t time.Time) *ProjectScoreCreate {
	psc.mutation.SetUpdateTime(t)
	return psc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (psc *ProjectScoreCreate) SetNillableUpdateTime(t *time.Time) *ProjectScoreCreate {
	if t != nil {
		psc.SetUpdateTime(*t)
	}
	return psc
}

// SetCaseID sets the "case_id" field.
func (psc *ProjectScoreCreate) SetCaseID(i int32) *ProjectScoreCreate {
	psc.mutation.SetCaseID(i)
	return psc
}

// SetScore sets the "score" field.
func (psc *ProjectScoreCreate) SetScore(f float64) *ProjectScoreCreate {
	psc.mutation.SetScore(f)
	return psc
}

// SetType sets the "type" field.
func (psc *ProjectScoreCreate) SetType(pr projectscore.Type) *ProjectScoreCreate {
	psc.mutation.SetType(pr)
	return psc
}

// SetTeamID sets the "team" edge to the ProjectTeam entity by ID.
func (psc *ProjectScoreCreate) SetTeamID(id int) *ProjectScoreCreate {
	psc.mutation.SetTeamID(id)
	return psc
}

// SetNillableTeamID sets the "team" edge to the ProjectTeam entity by ID if the given value is not nil.
func (psc *ProjectScoreCreate) SetNillableTeamID(id *int) *ProjectScoreCreate {
	if id != nil {
		psc = psc.SetTeamID(*id)
	}
	return psc
}

// SetTeam sets the "team" edge to the ProjectTeam entity.
func (psc *ProjectScoreCreate) SetTeam(p *ProjectTeam) *ProjectScoreCreate {
	return psc.SetTeamID(p.ID)
}

// Mutation returns the ProjectScoreMutation object of the builder.
func (psc *ProjectScoreCreate) Mutation() *ProjectScoreMutation {
	return psc.mutation
}

// Save creates the ProjectScore in the database.
func (psc *ProjectScoreCreate) Save(ctx context.Context) (*ProjectScore, error) {
	var (
		err  error
		node *ProjectScore
	)
	if err := psc.defaults(); err != nil {
		return nil, err
	}
	if len(psc.hooks) == 0 {
		if err = psc.check(); err != nil {
			return nil, err
		}
		node, err = psc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectScoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psc.check(); err != nil {
				return nil, err
			}
			psc.mutation = mutation
			if node, err = psc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(psc.hooks) - 1; i >= 0; i-- {
			if psc.hooks[i] == nil {
				return nil, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = psc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (psc *ProjectScoreCreate) SaveX(ctx context.Context) *ProjectScore {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *ProjectScoreCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *ProjectScoreCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *ProjectScoreCreate) defaults() error {
	if _, ok := psc.mutation.CreateTime(); !ok {
		if projectscore.DefaultCreateTime == nil {
			return fmt.Errorf("generated: uninitialized projectscore.DefaultCreateTime (forgotten import generated/runtime?)")
		}
		v := projectscore.DefaultCreateTime()
		psc.mutation.SetCreateTime(v)
	}
	if _, ok := psc.mutation.UpdateTime(); !ok {
		if projectscore.DefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized projectscore.DefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := projectscore.DefaultUpdateTime()
		psc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (psc *ProjectScoreCreate) check() error {
	if _, ok := psc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`generated: missing required field "ProjectScore.create_time"`)}
	}
	if _, ok := psc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`generated: missing required field "ProjectScore.update_time"`)}
	}
	if _, ok := psc.mutation.CaseID(); !ok {
		return &ValidationError{Name: "case_id", err: errors.New(`generated: missing required field "ProjectScore.case_id"`)}
	}
	if _, ok := psc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`generated: missing required field "ProjectScore.score"`)}
	}
	if _, ok := psc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`generated: missing required field "ProjectScore.type"`)}
	}
	if v, ok := psc.mutation.GetType(); ok {
		if err := projectscore.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`generated: validator failed for field "ProjectScore.type": %w`, err)}
		}
	}
	return nil
}

func (psc *ProjectScoreCreate) sqlSave(ctx context.Context) (*ProjectScore, error) {
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (psc *ProjectScoreCreate) createSpec() (*ProjectScore, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectScore{config: psc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: projectscore.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectscore.FieldID,
			},
		}
	)
	_spec.OnConflict = psc.conflict
	if value, ok := psc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projectscore.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := psc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projectscore.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := psc.mutation.CaseID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: projectscore.FieldCaseID,
		})
		_node.CaseID = value
	}
	if value, ok := psc.mutation.Score(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: projectscore.FieldScore,
		})
		_node.Score = value
	}
	if value, ok := psc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projectscore.FieldType,
		})
		_node.Type = value
	}
	if nodes := psc.mutation.TeamIDs(); len(nodes) > 0 {
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
		_node.project_team_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProjectScore.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectScoreUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (psc *ProjectScoreCreate) OnConflict(opts ...sql.ConflictOption) *ProjectScoreUpsertOne {
	psc.conflict = opts
	return &ProjectScoreUpsertOne{
		create: psc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProjectScore.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (psc *ProjectScoreCreate) OnConflictColumns(columns ...string) *ProjectScoreUpsertOne {
	psc.conflict = append(psc.conflict, sql.ConflictColumns(columns...))
	return &ProjectScoreUpsertOne{
		create: psc,
	}
}

type (
	// ProjectScoreUpsertOne is the builder for "upsert"-ing
	//  one ProjectScore node.
	ProjectScoreUpsertOne struct {
		create *ProjectScoreCreate
	}

	// ProjectScoreUpsert is the "OnConflict" setter.
	ProjectScoreUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *ProjectScoreUpsert) SetCreateTime(v time.Time) *ProjectScoreUpsert {
	u.Set(projectscore.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ProjectScoreUpsert) UpdateCreateTime() *ProjectScoreUpsert {
	u.SetExcluded(projectscore.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ProjectScoreUpsert) SetUpdateTime(v time.Time) *ProjectScoreUpsert {
	u.Set(projectscore.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ProjectScoreUpsert) UpdateUpdateTime() *ProjectScoreUpsert {
	u.SetExcluded(projectscore.FieldUpdateTime)
	return u
}

// SetCaseID sets the "case_id" field.
func (u *ProjectScoreUpsert) SetCaseID(v int32) *ProjectScoreUpsert {
	u.Set(projectscore.FieldCaseID, v)
	return u
}

// UpdateCaseID sets the "case_id" field to the value that was provided on create.
func (u *ProjectScoreUpsert) UpdateCaseID() *ProjectScoreUpsert {
	u.SetExcluded(projectscore.FieldCaseID)
	return u
}

// AddCaseID adds v to the "case_id" field.
func (u *ProjectScoreUpsert) AddCaseID(v int32) *ProjectScoreUpsert {
	u.Add(projectscore.FieldCaseID, v)
	return u
}

// SetScore sets the "score" field.
func (u *ProjectScoreUpsert) SetScore(v float64) *ProjectScoreUpsert {
	u.Set(projectscore.FieldScore, v)
	return u
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *ProjectScoreUpsert) UpdateScore() *ProjectScoreUpsert {
	u.SetExcluded(projectscore.FieldScore)
	return u
}

// AddScore adds v to the "score" field.
func (u *ProjectScoreUpsert) AddScore(v float64) *ProjectScoreUpsert {
	u.Add(projectscore.FieldScore, v)
	return u
}

// SetType sets the "type" field.
func (u *ProjectScoreUpsert) SetType(v projectscore.Type) *ProjectScoreUpsert {
	u.Set(projectscore.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ProjectScoreUpsert) UpdateType() *ProjectScoreUpsert {
	u.SetExcluded(projectscore.FieldType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.ProjectScore.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ProjectScoreUpsertOne) UpdateNewValues() *ProjectScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(projectscore.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ProjectScore.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ProjectScoreUpsertOne) Ignore() *ProjectScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectScoreUpsertOne) DoNothing() *ProjectScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectScoreCreate.OnConflict
// documentation for more info.
func (u *ProjectScoreUpsertOne) Update(set func(*ProjectScoreUpsert)) *ProjectScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectScoreUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ProjectScoreUpsertOne) SetCreateTime(v time.Time) *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ProjectScoreUpsertOne) UpdateCreateTime() *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ProjectScoreUpsertOne) SetUpdateTime(v time.Time) *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ProjectScoreUpsertOne) UpdateUpdateTime() *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCaseID sets the "case_id" field.
func (u *ProjectScoreUpsertOne) SetCaseID(v int32) *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetCaseID(v)
	})
}

// AddCaseID adds v to the "case_id" field.
func (u *ProjectScoreUpsertOne) AddCaseID(v int32) *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.AddCaseID(v)
	})
}

// UpdateCaseID sets the "case_id" field to the value that was provided on create.
func (u *ProjectScoreUpsertOne) UpdateCaseID() *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateCaseID()
	})
}

// SetScore sets the "score" field.
func (u *ProjectScoreUpsertOne) SetScore(v float64) *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetScore(v)
	})
}

// AddScore adds v to the "score" field.
func (u *ProjectScoreUpsertOne) AddScore(v float64) *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.AddScore(v)
	})
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *ProjectScoreUpsertOne) UpdateScore() *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateScore()
	})
}

// SetType sets the "type" field.
func (u *ProjectScoreUpsertOne) SetType(v projectscore.Type) *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ProjectScoreUpsertOne) UpdateType() *ProjectScoreUpsertOne {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *ProjectScoreUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for ProjectScoreCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectScoreUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProjectScoreUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProjectScoreUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProjectScoreCreateBulk is the builder for creating many ProjectScore entities in bulk.
type ProjectScoreCreateBulk struct {
	config
	builders []*ProjectScoreCreate
	conflict []sql.ConflictOption
}

// Save creates the ProjectScore entities in the database.
func (pscb *ProjectScoreCreateBulk) Save(ctx context.Context) ([]*ProjectScore, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*ProjectScore, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectScoreMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *ProjectScoreCreateBulk) SaveX(ctx context.Context) []*ProjectScore {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *ProjectScoreCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *ProjectScoreCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProjectScore.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectScoreUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (pscb *ProjectScoreCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProjectScoreUpsertBulk {
	pscb.conflict = opts
	return &ProjectScoreUpsertBulk{
		create: pscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProjectScore.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pscb *ProjectScoreCreateBulk) OnConflictColumns(columns ...string) *ProjectScoreUpsertBulk {
	pscb.conflict = append(pscb.conflict, sql.ConflictColumns(columns...))
	return &ProjectScoreUpsertBulk{
		create: pscb,
	}
}

// ProjectScoreUpsertBulk is the builder for "upsert"-ing
// a bulk of ProjectScore nodes.
type ProjectScoreUpsertBulk struct {
	create *ProjectScoreCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ProjectScore.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ProjectScoreUpsertBulk) UpdateNewValues() *ProjectScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(projectscore.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProjectScore.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ProjectScoreUpsertBulk) Ignore() *ProjectScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectScoreUpsertBulk) DoNothing() *ProjectScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectScoreCreateBulk.OnConflict
// documentation for more info.
func (u *ProjectScoreUpsertBulk) Update(set func(*ProjectScoreUpsert)) *ProjectScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectScoreUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ProjectScoreUpsertBulk) SetCreateTime(v time.Time) *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ProjectScoreUpsertBulk) UpdateCreateTime() *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ProjectScoreUpsertBulk) SetUpdateTime(v time.Time) *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ProjectScoreUpsertBulk) UpdateUpdateTime() *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCaseID sets the "case_id" field.
func (u *ProjectScoreUpsertBulk) SetCaseID(v int32) *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetCaseID(v)
	})
}

// AddCaseID adds v to the "case_id" field.
func (u *ProjectScoreUpsertBulk) AddCaseID(v int32) *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.AddCaseID(v)
	})
}

// UpdateCaseID sets the "case_id" field to the value that was provided on create.
func (u *ProjectScoreUpsertBulk) UpdateCaseID() *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateCaseID()
	})
}

// SetScore sets the "score" field.
func (u *ProjectScoreUpsertBulk) SetScore(v float64) *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetScore(v)
	})
}

// AddScore adds v to the "score" field.
func (u *ProjectScoreUpsertBulk) AddScore(v float64) *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.AddScore(v)
	})
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *ProjectScoreUpsertBulk) UpdateScore() *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateScore()
	})
}

// SetType sets the "type" field.
func (u *ProjectScoreUpsertBulk) SetType(v projectscore.Type) *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ProjectScoreUpsertBulk) UpdateType() *ProjectScoreUpsertBulk {
	return u.Update(func(s *ProjectScoreUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *ProjectScoreUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the ProjectScoreCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for ProjectScoreCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectScoreUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}