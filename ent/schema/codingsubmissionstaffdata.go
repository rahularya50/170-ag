package schema

import (
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"170-ag/site/policy"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// CodingSubmissionStaffData holds the schema definition for the CodingSubmissionStaffData entity.
type CodingSubmissionStaffData struct {
	ent.Schema
}

// Fields of the CodingSubmissionStaffData.
func (CodingSubmissionStaffData) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("execution_id").Nillable().Optional(),
		field.Text("input"),
		field.Text("output").MaxLen(8388607).Nillable().Optional(),
		field.Text("stderr").MaxLen(262143).Nillable().Optional(),
		field.Text("exit_error").MaxLen(262143).Nillable().Optional(),
	}
}

// Edges of the CodingSubmissionStaffData.
func (CodingSubmissionStaffData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coding_submission", CodingSubmission.Type).Required().Unique().Annotations(entgql.Bind()),
	}
}

func (CodingSubmissionStaffData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (CodingSubmissionStaffData) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("execution_id"),
	}
}

func (CodingSubmissionStaffData) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.SubmissionEnqueuingAccessToken),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.SubmissionEnqueuingAccessToken),
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
	}
}
