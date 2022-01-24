// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/codingtestcasedata"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CodingTestCaseDataCreate is the builder for creating a CodingTestCaseData entity.
type CodingTestCaseDataCreate struct {
	config
	mutation *CodingTestCaseDataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetInput sets the "input" field.
func (ctcdc *CodingTestCaseDataCreate) SetInput(s string) *CodingTestCaseDataCreate {
	ctcdc.mutation.SetInput(s)
	return ctcdc
}

// SetOutput sets the "output" field.
func (ctcdc *CodingTestCaseDataCreate) SetOutput(s string) *CodingTestCaseDataCreate {
	ctcdc.mutation.SetOutput(s)
	return ctcdc
}

// SetTestCaseID sets the "test_case" edge to the CodingTestCase entity by ID.
func (ctcdc *CodingTestCaseDataCreate) SetTestCaseID(id int) *CodingTestCaseDataCreate {
	ctcdc.mutation.SetTestCaseID(id)
	return ctcdc
}

// SetTestCase sets the "test_case" edge to the CodingTestCase entity.
func (ctcdc *CodingTestCaseDataCreate) SetTestCase(c *CodingTestCase) *CodingTestCaseDataCreate {
	return ctcdc.SetTestCaseID(c.ID)
}

// Mutation returns the CodingTestCaseDataMutation object of the builder.
func (ctcdc *CodingTestCaseDataCreate) Mutation() *CodingTestCaseDataMutation {
	return ctcdc.mutation
}

// Save creates the CodingTestCaseData in the database.
func (ctcdc *CodingTestCaseDataCreate) Save(ctx context.Context) (*CodingTestCaseData, error) {
	var (
		err  error
		node *CodingTestCaseData
	)
	if len(ctcdc.hooks) == 0 {
		if err = ctcdc.check(); err != nil {
			return nil, err
		}
		node, err = ctcdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodingTestCaseDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctcdc.check(); err != nil {
				return nil, err
			}
			ctcdc.mutation = mutation
			if node, err = ctcdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctcdc.hooks) - 1; i >= 0; i-- {
			if ctcdc.hooks[i] == nil {
				return nil, fmt.Errorf("generated: uninitialized hook (forgotten import generated/runtime?)")
			}
			mut = ctcdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctcdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctcdc *CodingTestCaseDataCreate) SaveX(ctx context.Context) *CodingTestCaseData {
	v, err := ctcdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcdc *CodingTestCaseDataCreate) Exec(ctx context.Context) error {
	_, err := ctcdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcdc *CodingTestCaseDataCreate) ExecX(ctx context.Context) {
	if err := ctcdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctcdc *CodingTestCaseDataCreate) check() error {
	if _, ok := ctcdc.mutation.Input(); !ok {
		return &ValidationError{Name: "input", err: errors.New(`generated: missing required field "CodingTestCaseData.input"`)}
	}
	if _, ok := ctcdc.mutation.Output(); !ok {
		return &ValidationError{Name: "output", err: errors.New(`generated: missing required field "CodingTestCaseData.output"`)}
	}
	if _, ok := ctcdc.mutation.TestCaseID(); !ok {
		return &ValidationError{Name: "test_case", err: errors.New(`generated: missing required edge "CodingTestCaseData.test_case"`)}
	}
	return nil
}

func (ctcdc *CodingTestCaseDataCreate) sqlSave(ctx context.Context) (*CodingTestCaseData, error) {
	_node, _spec := ctcdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctcdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctcdc *CodingTestCaseDataCreate) createSpec() (*CodingTestCaseData, *sqlgraph.CreateSpec) {
	var (
		_node = &CodingTestCaseData{config: ctcdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: codingtestcasedata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingtestcasedata.FieldID,
			},
		}
	)
	_spec.OnConflict = ctcdc.conflict
	if value, ok := ctcdc.mutation.Input(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingtestcasedata.FieldInput,
		})
		_node.Input = value
	}
	if value, ok := ctcdc.mutation.Output(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: codingtestcasedata.FieldOutput,
		})
		_node.Output = value
	}
	if nodes := ctcdc.mutation.TestCaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   codingtestcasedata.TestCaseTable,
			Columns: []string{codingtestcasedata.TestCaseColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CodingTestCaseData.Create().
//		SetInput(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CodingTestCaseDataUpsert) {
//			SetInput(v+v).
//		}).
//		Exec(ctx)
//
func (ctcdc *CodingTestCaseDataCreate) OnConflict(opts ...sql.ConflictOption) *CodingTestCaseDataUpsertOne {
	ctcdc.conflict = opts
	return &CodingTestCaseDataUpsertOne{
		create: ctcdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CodingTestCaseData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ctcdc *CodingTestCaseDataCreate) OnConflictColumns(columns ...string) *CodingTestCaseDataUpsertOne {
	ctcdc.conflict = append(ctcdc.conflict, sql.ConflictColumns(columns...))
	return &CodingTestCaseDataUpsertOne{
		create: ctcdc,
	}
}

type (
	// CodingTestCaseDataUpsertOne is the builder for "upsert"-ing
	//  one CodingTestCaseData node.
	CodingTestCaseDataUpsertOne struct {
		create *CodingTestCaseDataCreate
	}

	// CodingTestCaseDataUpsert is the "OnConflict" setter.
	CodingTestCaseDataUpsert struct {
		*sql.UpdateSet
	}
)

// SetInput sets the "input" field.
func (u *CodingTestCaseDataUpsert) SetInput(v string) *CodingTestCaseDataUpsert {
	u.Set(codingtestcasedata.FieldInput, v)
	return u
}

// UpdateInput sets the "input" field to the value that was provided on create.
func (u *CodingTestCaseDataUpsert) UpdateInput() *CodingTestCaseDataUpsert {
	u.SetExcluded(codingtestcasedata.FieldInput)
	return u
}

// SetOutput sets the "output" field.
func (u *CodingTestCaseDataUpsert) SetOutput(v string) *CodingTestCaseDataUpsert {
	u.Set(codingtestcasedata.FieldOutput, v)
	return u
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *CodingTestCaseDataUpsert) UpdateOutput() *CodingTestCaseDataUpsert {
	u.SetExcluded(codingtestcasedata.FieldOutput)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.CodingTestCaseData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CodingTestCaseDataUpsertOne) UpdateNewValues() *CodingTestCaseDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CodingTestCaseData.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CodingTestCaseDataUpsertOne) Ignore() *CodingTestCaseDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CodingTestCaseDataUpsertOne) DoNothing() *CodingTestCaseDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CodingTestCaseDataCreate.OnConflict
// documentation for more info.
func (u *CodingTestCaseDataUpsertOne) Update(set func(*CodingTestCaseDataUpsert)) *CodingTestCaseDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CodingTestCaseDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetInput sets the "input" field.
func (u *CodingTestCaseDataUpsertOne) SetInput(v string) *CodingTestCaseDataUpsertOne {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.SetInput(v)
	})
}

// UpdateInput sets the "input" field to the value that was provided on create.
func (u *CodingTestCaseDataUpsertOne) UpdateInput() *CodingTestCaseDataUpsertOne {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.UpdateInput()
	})
}

// SetOutput sets the "output" field.
func (u *CodingTestCaseDataUpsertOne) SetOutput(v string) *CodingTestCaseDataUpsertOne {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.SetOutput(v)
	})
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *CodingTestCaseDataUpsertOne) UpdateOutput() *CodingTestCaseDataUpsertOne {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.UpdateOutput()
	})
}

// Exec executes the query.
func (u *CodingTestCaseDataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for CodingTestCaseDataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CodingTestCaseDataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CodingTestCaseDataUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CodingTestCaseDataUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CodingTestCaseDataCreateBulk is the builder for creating many CodingTestCaseData entities in bulk.
type CodingTestCaseDataCreateBulk struct {
	config
	builders []*CodingTestCaseDataCreate
	conflict []sql.ConflictOption
}

// Save creates the CodingTestCaseData entities in the database.
func (ctcdcb *CodingTestCaseDataCreateBulk) Save(ctx context.Context) ([]*CodingTestCaseData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcdcb.builders))
	nodes := make([]*CodingTestCaseData, len(ctcdcb.builders))
	mutators := make([]Mutator, len(ctcdcb.builders))
	for i := range ctcdcb.builders {
		func(i int, root context.Context) {
			builder := ctcdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CodingTestCaseDataMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ctcdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcdcb *CodingTestCaseDataCreateBulk) SaveX(ctx context.Context) []*CodingTestCaseData {
	v, err := ctcdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcdcb *CodingTestCaseDataCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcdcb *CodingTestCaseDataCreateBulk) ExecX(ctx context.Context) {
	if err := ctcdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CodingTestCaseData.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CodingTestCaseDataUpsert) {
//			SetInput(v+v).
//		}).
//		Exec(ctx)
//
func (ctcdcb *CodingTestCaseDataCreateBulk) OnConflict(opts ...sql.ConflictOption) *CodingTestCaseDataUpsertBulk {
	ctcdcb.conflict = opts
	return &CodingTestCaseDataUpsertBulk{
		create: ctcdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CodingTestCaseData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ctcdcb *CodingTestCaseDataCreateBulk) OnConflictColumns(columns ...string) *CodingTestCaseDataUpsertBulk {
	ctcdcb.conflict = append(ctcdcb.conflict, sql.ConflictColumns(columns...))
	return &CodingTestCaseDataUpsertBulk{
		create: ctcdcb,
	}
}

// CodingTestCaseDataUpsertBulk is the builder for "upsert"-ing
// a bulk of CodingTestCaseData nodes.
type CodingTestCaseDataUpsertBulk struct {
	create *CodingTestCaseDataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CodingTestCaseData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CodingTestCaseDataUpsertBulk) UpdateNewValues() *CodingTestCaseDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CodingTestCaseData.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CodingTestCaseDataUpsertBulk) Ignore() *CodingTestCaseDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CodingTestCaseDataUpsertBulk) DoNothing() *CodingTestCaseDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CodingTestCaseDataCreateBulk.OnConflict
// documentation for more info.
func (u *CodingTestCaseDataUpsertBulk) Update(set func(*CodingTestCaseDataUpsert)) *CodingTestCaseDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CodingTestCaseDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetInput sets the "input" field.
func (u *CodingTestCaseDataUpsertBulk) SetInput(v string) *CodingTestCaseDataUpsertBulk {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.SetInput(v)
	})
}

// UpdateInput sets the "input" field to the value that was provided on create.
func (u *CodingTestCaseDataUpsertBulk) UpdateInput() *CodingTestCaseDataUpsertBulk {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.UpdateInput()
	})
}

// SetOutput sets the "output" field.
func (u *CodingTestCaseDataUpsertBulk) SetOutput(v string) *CodingTestCaseDataUpsertBulk {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.SetOutput(v)
	})
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *CodingTestCaseDataUpsertBulk) UpdateOutput() *CodingTestCaseDataUpsertBulk {
	return u.Update(func(s *CodingTestCaseDataUpsert) {
		s.UpdateOutput()
	})
}

// Exec executes the query.
func (u *CodingTestCaseDataUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the CodingTestCaseDataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for CodingTestCaseDataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CodingTestCaseDataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}