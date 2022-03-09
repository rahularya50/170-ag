package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/privacy"
	"170-ag/ent/generated/user"
	"170-ag/ent/models"
	"170-ag/privacyrules"
	"170-ag/site"
	"170-ag/site/policy"
	"170-ag/site/web"
	"context"
	"time"

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
		field.Bool("is_validation").Immutable().Default(false),
		field.Enum("status").
			NamedValues(
				"Queued", "QUEUED",
				"Running", "RUNNING",
				"Completed", "COMPLETED",
				"Internal Error", "INTERNAL_ERROR",
			).
			Default("QUEUED"),
		field.Int("points").Optional().Nillable(),
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

func denyIfRecentViewerSubmissionExistsSinceThreshold(threshold time.Duration) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(c context.Context) error {
		viewer, ok := web.ViewerFromContext(c)
		if !ok {
			return privacy.Deny
		}
		recent_exists, err := site.EntClientFromContext(c).
			CodingSubmission.
			Query().
			Where(
				codingsubmission.HasAuthorWith(user.IDEQ(viewer.ID)),
				codingsubmission.CreateTimeGT(time.Now().Add(-threshold)),
			).Exist(c)
		if recent_exists || err != nil {
			return privacy.Deny
		}
		return privacy.Skip
	})
}

func (CodingSubmission) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			// students can only create submissions
			privacy.CodingSubmissionMutationRuleFunc(func(c context.Context, cdm *generated.CodingSubmissionMutation) error {
				if !cdm.Op().Is(ent.OpCreate) {
					return privacy.Deny
				}
				return privacy.Skip
			}),
			privacy.CodingSubmissionMutationRuleFunc(func(c context.Context, csm *generated.CodingSubmissionMutation) error {
				status, statusSet := csm.Status()
				if status == codingsubmission.DefaultStatus {
					// students can enqueue submissions on creation
				} else if statusSet {
					return privacy.Deny
				}

				_, pointsSet := csm.Points()
				if pointsSet {
					return privacy.Deny
				}

				_, resultsSet := csm.Results()
				if resultsSet {
					return privacy.Deny
				}

				return privacy.Skip
			}),
			denyIfRecentViewerSubmissionExistsSinceThreshold(time.Minute),
			policy.FilterToViewerID(func(c context.Context, f privacy.Filter, user_id int) error {
				f.(*generated.CodingSubmissionFilter).WhereHasAuthorWith(user.ID(user_id))
				return privacy.Allow
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			policy.AllowQueryIfIDsMatchViewer(func(c context.Context, q ent.Query) ([]int, error) {
				return q.(*generated.CodingSubmissionQuery).Clone().QueryAuthor().IDs(c)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
