package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CodingTestCase holds the schema definition for the CodingTestCase entity.
type CodingTestCase struct {
	ent.Schema
}

// Fields of the CodingTestCase.
func (CodingTestCase) Fields() []ent.Field {
	return []ent.Field{
		field.Text("input"),
		field.Text("output"),
		field.Int("points").NonNegative(),
		field.Bool("visible"),
	}
}

// Edges of the CodingTestCase.
func (CodingTestCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coding_problem", CodingProblem.Type).
			Unique().
			Ref("test_cases").
			Required().
			Annotations(entgql.Bind()),
	}
}

func denyQueryIfCodingProblemsUnreleased() privacy.CodingTestCaseQueryRuleFunc {
	return privacy.CodingTestCaseQueryRuleFunc(func(c context.Context, q *generated.CodingTestCaseQuery) error {
		problems, err := q.Clone().QueryCodingProblem().All(privacy.DecisionContext(c, privacy.Allow))
		if err != nil {
			return privacy.Deny
		}
		for _, problem := range problems {
			if !problem.Released {
				return privacy.Deny
			}
		}
		return privacy.Skip
	})
}

func (CodingTestCase) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.AllowQueryIfSubPolicyPasses(
				denyQueryIfCodingProblemsUnreleased(),
				privacyrules.AllowWithPrivacyAccessToken(privacyrules.SubmissionEnqueuingAccessToken),
			),
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
	}
}
