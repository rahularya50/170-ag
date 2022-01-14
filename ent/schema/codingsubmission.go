package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/ent/generated/user"
	"170-ag/privacyrules"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CodingSubmission holds the schema definition for the CodingSubmission entity.
type CodingSubmission struct {
	ent.Schema
}

// Fields of the CodingSubmission.
func (CodingSubmission) Fields() []ent.Field {
	return []ent.Field{
		field.Text("code"),
		field.Enum("status").
			NamedValues(
				"Queued", "QUEUED",
				"Running", "RUNNING",
				"Completed", "COMPLETED",
			).
			Default("QUEUED"),
	}
}

// Edges of the CodingSubmission.
func (CodingSubmission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Required().Unique().Annotations(entgql.Bind()),
		edge.To("coding_problem", CodingProblem.Type).Required().Unique().Annotations(entgql.Bind()),
	}
}

func (CodingSubmission) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("author", "coding_problem"),
		index.Edges("coding_problem"),
	}
}

func (CodingSubmission) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.DenyIfNoViewer(),
			privacy.CodingSubmissionMutationRuleFunc(func(c context.Context, csm *generated.CodingSubmissionMutation) error {
				_, statusSet := csm.Status()
				if statusSet {
					return privacy.Deny
				}
				return privacy.Skip
			}),
			privacyrules.AllowIfViewerIsStaff(),
			privacyrules.FilterToViewerID(func(c context.Context, f privacy.Filter, user_id int) error {
				f.(*generated.CodingSubmissionFilter).WhereHasAuthorWith(user.ID(user_id))
				return privacy.Allow
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacyrules.AllowQueryIfIDsMatchViewer(func(c context.Context, q ent.Query) ([]int, error) {
				return q.(*generated.CodingSubmissionQuery).Clone().QueryAuthor().IDs(c)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
