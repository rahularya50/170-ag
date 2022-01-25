package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"170-ag/site/policy"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CodingTestCaseData holds the schema definition for the CodingTestCaseData entity.
type CodingTestCaseData struct {
	ent.Schema
}

// Fields of the CodingTestCaseData.
func (CodingTestCaseData) Fields() []ent.Field {
	return []ent.Field{
		field.Text("input").Default("0\n"),
		field.Text("output").Default(""),
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

func (CodingTestCaseData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
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
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			// the associated test cases must be visible to the viewer
			privacyrules.DenyQueryIfSubPolicyFails(
				privacyrules.AllowQueryIfDelegatesVisible(func(allow_ctx context.Context, q generated.Query) ([]int, error) {
					return q.(*generated.CodingTestCaseDataQuery).Clone().QueryTestCase().IDs(allow_ctx)
				}, func(client *generated.Client, ctx context.Context, ids []int) ([]int, error) {
					return client.CodingTestCase.Query().Where(codingtestcase.IDIn(ids...)).IDs(ctx)
				}),
				privacy.AlwaysDenyRule(),
			),
			// The judge service account can always read test case data
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			policy.DenyIfNoViewer(),
			// staff can see private test data
			policy.AllowIfViewerIsStaff(),
			// if the case is public, we can show its data to anyone who can see the case
			allowIfTestCaseIsPublic(),
			// we also need to read its contents when sending a submission to the judge
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.SubmissionEnqueuingAccessToken),
			privacy.AlwaysDenyRule(),
		},
	}
}
