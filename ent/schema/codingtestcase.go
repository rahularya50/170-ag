package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"170-ag/site/policy"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CodingTestCase holds the schema definition for the CodingTestCase entity.
type CodingTestCase struct {
	ent.Schema
}

// Fields of the CodingTestCase.
func (CodingTestCase) Fields() []ent.Field {
	return []ent.Field{
		field.Int("points").NonNegative().Default(0),
		field.Bool("public").Default(false),
	}
}

// Edges of the CodingTestCase.
func (CodingTestCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coding_problem", CodingProblem.Type).
			Ref("test_cases").
			Unique().
			Required().
			Annotations(entgql.Bind()),
		edge.From("data", CodingTestCaseData.Type).
			Unique().
			Ref("test_case").
			Annotations(
				entgql.Bind(),
				entsql.Annotation{OnDelete: entsql.Cascade},
			),
	}
}

func (CodingTestCase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (CodingTestCase) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			// test cases are visible iff their associated problem is visible (the test *data* is in a different ent)
			privacyrules.AllowQueryIfDelegatesVisible(func(allow_ctx context.Context, q generated.Query) ([]int, error) {
				return q.(*generated.CodingTestCaseQuery).Clone().QueryCodingProblem().IDs(allow_ctx)
			}, func(client *generated.Client, ctx context.Context, ids []int) ([]int, error) {
				return client.CodingProblem.Query().Where(codingproblem.IDIn(ids...)).IDs(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
