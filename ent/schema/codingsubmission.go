package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/privacy"
	"170-ag/ent/generated/user"
	"170-ag/ent/models"
	"170-ag/privacyrules"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// CodingSubmission holds the schema definition for the CodingSubmission entity.
type CodingSubmission struct {
	ent.Schema
}

// Fields of the CodingSubmission.
func (CodingSubmission) Fields() []ent.Field {
	return []ent.Field{
		field.Text("code").Immutable().MaxLen(65535),
		field.Enum("status").
			NamedValues(
				"Queued", "QUEUED",
				"Running", "RUNNING",
				"Completed", "COMPLETED",
			).
			Default("QUEUED"),
		field.JSON("results", models.CodingSubmissionResults{}).Optional(),
	}
}

// Edges of the CodingSubmission.
func (CodingSubmission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Required().Unique().Annotations(entgql.Bind()),
		edge.To("coding_problem", CodingProblem.Type).Required().Unique().Annotations(entgql.Bind()),
		edge.From("staff_data", CodingSubmissionStaffData.Type).Ref("coding_submission").Unique().Annotations(entgql.Bind()),
	}
}

func (CodingSubmission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (CodingSubmission) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("author", "coding_problem"),
		index.Edges("coding_problem"),
		index.Fields("status"),
	}
}

func (CodingSubmission) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacy.CodingSubmissionMutationRuleFunc(func(c context.Context, csm *generated.CodingSubmissionMutation) error {
				status, statusSet := csm.Status()

				if csm.Op().Is(ent.OpCreate) && status == codingsubmission.DefaultStatus {
					// this is OK
				} else if statusSet {
					return privacy.Denyf("Cannot explicitly set status of submission, except to initialize to default")
				}
				return privacy.Skip
			}),
			privacyrules.FilterToViewerID(func(c context.Context, f privacy.Filter, user_id int) error {
				f.(*generated.CodingSubmissionFilter).WhereHasAuthorWith(user.ID(user_id))
				return privacy.Allow
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacyrules.AllowQueryIfIDsMatchViewer(func(c context.Context, q ent.Query) ([]int, error) {
				return q.(*generated.CodingSubmissionQuery).Clone().QueryAuthor().IDs(c)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
