// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingtestcase"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CodingTestCaseCreate is the builder for creating a CodingTestCase entity.
type CodingTestCaseCreate struct {
	config
	mutation *CodingTestCaseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetInput sets the "input" field.
func (ctcc *CodingTestCaseCreate) SetInput(s string) *CodingTestCaseCreate {
	ctcc.mutation.SetInput(s)
	return ctcc
}

// SetOutput sets the "output" field.
func (ctcc *CodingTestCaseCreate) SetOutput(s string) *CodingTestCaseCreate {
	ctcc.mutation.SetOutput(s)
	return ctcc
}

// SetPoints sets the "points" field.
func (ctcc *CodingTestCaseCreate) SetPoints(i int) *CodingTestCaseCreate {
	ctcc.mutation.SetPoints(i)
	return ctcc
}

// SetVisible sets the "visible" field.
func (ctcc *CodingTestCaseCreate) SetVisible(b bool) *CodingTestCaseCreate {
	ctcc.mutation.SetVisible(b)
	return ctcc
}

// AddCodingProblemIDs adds the "coding_problem" edge to the CodingProblem entity by IDs.
func (ctcc *CodingTestCaseCreate) AddCodingProblemIDs(ids ...int) *CodingTestCaseCreate {
	ctcc.mutation.AddCodingProblemIDs(ids...)
	return ctcc
}

// AddCodingProblem adds the "coding_problem" edges to the CodingProblem entity.
func (ctcc *CodingTestCaseCreate) AddCodingProblem(c ...*CodingProblem) *CodingTestCaseCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctcc.AddCodingProblemIDs(ids...)
}

// Mutation returns the CodingTestCaseMutation object of the builder.
func (ctcc *CodingTestCaseCreate) Mutation() *CodingTestCaseMutation {
	return ctcc.mutation
}

// Save creates the CodingTestCase in the database.
func (ctcc *CodingTestCaseCreate) Save(ctx context.Context) (*CodingTestCase, error) {
	var (
		err  error
		node *CodingTestCase
	)
	if len(ctcc.hooks) == 0 {
		if err = ctcc.check(); err != nil {
			return nil, err
		}
		node, err = ctcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctcc.check(); err != nil {
				return nil, err
			}
			ctcc.mutation = mutation
			if node, err = ctcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctcc.hooks) - 1; i >= 0; i-- {
			if ctcc.hooks[i] == nil {
				return nil, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = ctcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctcc *CodingTestCaseCreate) SaveX(ctx context.Context) *CodingTestCase {
	v, err := ctcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcc *CodingTestCaseCreate) Exec(ctx context.Context) error {
	_, err := ctcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcc *CodingTestCaseCreate) ExecX(ctx context.Context) {
	if err := ctcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctcc *CodingTestCaseCreate) check() error {
	if _, ok := ctcc.mutation.Input(); !ok {
		return &ValidationError{Name: "input", err: errors.New(`generated: missing required field "input"`)}
	}
	if _, ok := ctcc.mutation.Output(); !ok {
		return &ValidationError{Name: "output", err: errors.New(`generated: missing required field "output"`)}
	}
	if _, ok := ctcc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`generated: missing required field "points"`)}
	}
	if v, ok := ctcc.mutation.Points(); ok {
		if err := codingtestcase.PointsValidator(v); err != nil {
			return &ValidationError{Name: "points", err: fmt.Errorf(`generated: validator failed for field "points": %w`, err)}
		}
	}
	if _, ok := ctcc.mutation.Visible(); !ok {
		return &ValidationError{Name: "visible", err: errors.New(`generated: missing required field "visible"`)}
	}
	return nil
}

func (ctcc *CodingTestCaseCreate) sqlSave(ctx context.Context) (*CodingTestCase, error) {
	_node, _spec := ctcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctcc *CodingTestCaseCreate) createSpec() (*CodingTestCase, *sqlgraph.CreateSpec) {
	var (
		_node = &CodingTestCase{config: ctcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: codingtestcase.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingtestcase.FieldID,
			},
		}
	)
	_spec.OnConflict = ctcc.conflict
	if value, ok := ctcc.mutation.Input(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingtestcase.FieldInput,
		})
		_node.Input = value
	}
	if value, ok := ctcc.mutation.Output(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingtestcase.FieldOutput,
		})
		_node.Output = value
	}
	if value, ok := ctcc.mutation.Points(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: codingtestcase.FieldPoints,
		})
		_node.Points = value
	}
	if value, ok := ctcc.mutation.Visible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: codingtestcase.FieldVisible,
		})
		_node.Visible = value
	}
	if nodes := ctcc.mutation.CodingProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   codingtestcase.CodingProblemTable,
			Columns: codingtestcase.CodingProblemPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CodingTestCase.Create().
//		SetInput(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CodingTestCaseUpsert) {
//			SetInput(v+v).
//		}).
//		Exec(ctx)
//
func (ctcc *CodingTestCaseCreate) OnConflict(opts ...sql.ConflictOption) *CodingTestCaseUpsertOne {
	ctcc.conflict = opts
	return &CodingTestCaseUpsertOne{
		create: ctcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CodingTestCase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ctcc *CodingTestCaseCreate) OnConflictColumns(columns ...string) *CodingTestCaseUpsertOne {
	ctcc.conflict = append(ctcc.conflict, sql.ConflictColumns(columns...))
	return &CodingTestCaseUpsertOne{
		create: ctcc,
	}
}

type (
	// CodingTestCaseUpsertOne is the builder for "upsert"-ing
	//  one CodingTestCase node.
	CodingTestCaseUpsertOne struct {
		create *CodingTestCaseCreate
	}

	// CodingTestCaseUpsert is the "OnConflict" setter.
	CodingTestCaseUpsert struct {
		*sql.UpdateSet
	}
)

// SetInput sets the "input" field.
func (u *CodingTestCaseUpsert) SetInput(v string) *CodingTestCaseUpsert {
	u.Set(codingtestcase.FieldInput, v)
	return u
}

// UpdateInput sets the "input" field to the value that was provided on create.
func (u *CodingTestCaseUpsert) UpdateInput() *CodingTestCaseUpsert {
	u.SetExcluded(codingtestcase.FieldInput)
	return u
}

// SetOutput sets the "output" field.
func (u *CodingTestCaseUpsert) SetOutput(v string) *CodingTestCaseUpsert {
	u.Set(codingtestcase.FieldOutput, v)
	return u
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *CodingTestCaseUpsert) UpdateOutput() *CodingTestCaseUpsert {
	u.SetExcluded(codingtestcase.FieldOutput)
	return u
}

// SetPoints sets the "points" field.
func (u *CodingTestCaseUpsert) SetPoints(v int) *CodingTestCaseUpsert {
	u.Set(codingtestcase.FieldPoints, v)
	return u
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *CodingTestCaseUpsert) UpdatePoints() *CodingTestCaseUpsert {
	u.SetExcluded(codingtestcase.FieldPoints)
	return u
}

// SetVisible sets the "visible" field.
func (u *CodingTestCaseUpsert) SetVisible(v bool) *CodingTestCaseUpsert {
	u.Set(codingtestcase.FieldVisible, v)
	return u
}

// UpdateVisible sets the "visible" field to the value that was provided on create.
func (u *CodingTestCaseUpsert) UpdateVisible() *CodingTestCaseUpsert {
	u.SetExcluded(codingtestcase.FieldVisible)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.CodingTestCase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CodingTestCaseUpsertOne) UpdateNewValues() *CodingTestCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CodingTestCase.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CodingTestCaseUpsertOne) Ignore() *CodingTestCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CodingTestCaseUpsertOne) DoNothing() *CodingTestCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CodingTestCaseCreate.OnConflict
// documentation for more info.
func (u *CodingTestCaseUpsertOne) Update(set func(*CodingTestCaseUpsert)) *CodingTestCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CodingTestCaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetInput sets the "input" field.
func (u *CodingTestCaseUpsertOne) SetInput(v string) *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetInput(v)
	})
}

// UpdateInput sets the "input" field to the value that was provided on create.
func (u *CodingTestCaseUpsertOne) UpdateInput() *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdateInput()
	})
}

// SetOutput sets the "output" field.
func (u *CodingTestCaseUpsertOne) SetOutput(v string) *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetOutput(v)
	})
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *CodingTestCaseUpsertOne) UpdateOutput() *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdateOutput()
	})
}

// SetPoints sets the "points" field.
func (u *CodingTestCaseUpsertOne) SetPoints(v int) *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *CodingTestCaseUpsertOne) UpdatePoints() *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdatePoints()
	})
}

// SetVisible sets the "visible" field.
func (u *CodingTestCaseUpsertOne) SetVisible(v bool) *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetVisible(v)
	})
}

// UpdateVisible sets the "visible" field to the value that was provided on create.
func (u *CodingTestCaseUpsertOne) UpdateVisible() *CodingTestCaseUpsertOne {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdateVisible()
	})
}

// Exec executes the query.
func (u *CodingTestCaseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for CodingTestCaseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CodingTestCaseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CodingTestCaseUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CodingTestCaseUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CodingTestCaseCreateBulk is the builder for creating many CodingTestCase entities in bulk.
type CodingTestCaseCreateBulk struct {
	config
	builders []*CodingTestCaseCreate
	conflict []sql.ConflictOption
}

// Save creates the CodingTestCase entities in the database.
func (ctccb *CodingTestCaseCreateBulk) Save(ctx context.Context) ([]*CodingTestCase, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctccb.builders))
	nodes := make([]*CodingTestCase, len(ctccb.builders))
	mutators := make([]Mutator, len(ctccb.builders))
	for i := range ctccb.builders {
		func(i int, root context.Context) {
			builder := ctccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CodingTestCaseMutation)
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
					_, err = mutators[i+1].Mutate(root, ctccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ctccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctccb *CodingTestCaseCreateBulk) SaveX(ctx context.Context) []*CodingTestCase {
	v, err := ctccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctccb *CodingTestCaseCreateBulk) Exec(ctx context.Context) error {
	_, err := ctccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctccb *CodingTestCaseCreateBulk) ExecX(ctx context.Context) {
	if err := ctccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CodingTestCase.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CodingTestCaseUpsert) {
//			SetInput(v+v).
//		}).
//		Exec(ctx)
//
func (ctccb *CodingTestCaseCreateBulk) OnConflict(opts ...sql.ConflictOption) *CodingTestCaseUpsertBulk {
	ctccb.conflict = opts
	return &CodingTestCaseUpsertBulk{
		create: ctccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CodingTestCase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ctccb *CodingTestCaseCreateBulk) OnConflictColumns(columns ...string) *CodingTestCaseUpsertBulk {
	ctccb.conflict = append(ctccb.conflict, sql.ConflictColumns(columns...))
	return &CodingTestCaseUpsertBulk{
		create: ctccb,
	}
}

// CodingTestCaseUpsertBulk is the builder for "upsert"-ing
// a bulk of CodingTestCase nodes.
type CodingTestCaseUpsertBulk struct {
	create *CodingTestCaseCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CodingTestCase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CodingTestCaseUpsertBulk) UpdateNewValues() *CodingTestCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CodingTestCase.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CodingTestCaseUpsertBulk) Ignore() *CodingTestCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CodingTestCaseUpsertBulk) DoNothing() *CodingTestCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CodingTestCaseCreateBulk.OnConflict
// documentation for more info.
func (u *CodingTestCaseUpsertBulk) Update(set func(*CodingTestCaseUpsert)) *CodingTestCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CodingTestCaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetInput sets the "input" field.
func (u *CodingTestCaseUpsertBulk) SetInput(v string) *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetInput(v)
	})
}

// UpdateInput sets the "input" field to the value that was provided on create.
func (u *CodingTestCaseUpsertBulk) UpdateInput() *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdateInput()
	})
}

// SetOutput sets the "output" field.
func (u *CodingTestCaseUpsertBulk) SetOutput(v string) *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetOutput(v)
	})
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *CodingTestCaseUpsertBulk) UpdateOutput() *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdateOutput()
	})
}

// SetPoints sets the "points" field.
func (u *CodingTestCaseUpsertBulk) SetPoints(v int) *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *CodingTestCaseUpsertBulk) UpdatePoints() *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdatePoints()
	})
}

// SetVisible sets the "visible" field.
func (u *CodingTestCaseUpsertBulk) SetVisible(v bool) *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.SetVisible(v)
	})
}

// UpdateVisible sets the "visible" field to the value that was provided on create.
func (u *CodingTestCaseUpsertBulk) UpdateVisible() *CodingTestCaseUpsertBulk {
	return u.Update(func(s *CodingTestCaseUpsert) {
		s.UpdateVisible()
	})
}

// Exec executes the query.
func (u *CodingTestCaseUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the CodingTestCaseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for CodingTestCaseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CodingTestCaseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}