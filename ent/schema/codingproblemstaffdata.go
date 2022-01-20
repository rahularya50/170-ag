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

// CodingProblemStaffData holds the schema definition for the CodingProblemStaffData entity.
type CodingProblemStaffData struct {
	ent.Schema
}

// Fields of the CodingProblemStaffData.
func (CodingProblemStaffData) Fields() []ent.Field {
	return []ent.Field{
		field.Text("input"),
	}
}

// Edges of the CodingProblemStaffData.
func (CodingProblemStaffData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coding_problem", CodingProblem.Type).Required().Unique().Annotations(entgql.Bind()),
	}
}

func denyQueryIfCodingProblemsUnreleased() privacy.CodingProblemStaffDataQueryRuleFunc {
	return privacy.CodingProblemStaffDataQueryRuleFunc(func(c context.Context, q *generated.CodingProblemStaffDataQuery) error {
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

func (CodingProblemStaffData) Policy() ent.Policy {
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
