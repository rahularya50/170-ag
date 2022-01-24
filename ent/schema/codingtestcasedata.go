package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CodingTestCaseData holds the schema definition for the CodingTestCaseData entity.
type CodingTestCaseData struct {
	ent.Schema
}

// Fields of the CodingTestCaseData.
func (CodingTestCaseData) Fields() []ent.Field {
	return []ent.Field{
		field.Text("input"),
		field.Text("output"),
	}
}

// Edges of the CodingTestCaseData.
func (CodingTestCaseData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("test_case", CodingTestCase.Type).
			Unique().
			Required().
			Annotations(entgql.Bind()),
	}
}

func allowIfTestCaseIsPublic() privacy.QueryRule {
	return privacy.CodingTestCaseDataQueryRuleFunc(func(ctx context.Context, q *generated.CodingTestCaseDataQuery) error {
		allow_ctx := privacy.DecisionContext(ctx, privacy.Allow)
		test_cases, err := q.Clone().QueryTestCase().All(allow_ctx)
		if err != nil {
			return err
		}
		for _, test_case := range test_cases {
			if !test_case.Public {
				return privacy.Skip
			}
		}
		return privacy.Allow
	})
}

func (CodingTestCaseData) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.DenyIfNoViewer(),
			// the associated test cases must be visible
			privacyrules.DenyQueryIfSubPolicyFails(
				privacyrules.AllowQueryIfDelegatesVisible(func(allow_ctx context.Context, q generated.Query) ([]int, error) {
					return q.(*generated.CodingTestCaseDataQuery).Clone().QueryTestCase().IDs(allow_ctx)
				}, func(client *generated.Client, ctx context.Context, ids []int) ([]int, error) {
					return client.CodingTestCase.Query().Where(codingtestcase.IDIn(ids...)).IDs(ctx)
				}),
				privacy.AlwaysDenyRule(),
			),
			// if the case is public, we can show its contents
			allowIfTestCaseIsPublic(),
			// we also need to read its contents when sending a submision to the judge
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.SubmissionEnqueuingAccessToken),
			privacy.AlwaysDenyRule(),
		},
	}
}
