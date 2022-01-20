package schema

import (
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.Text("output").MaxLen(65535).Nillable().Optional(),
		field.Text("stderr").MaxLen(65535).Nillable().Optional(),
		field.Text("exit_error").MaxLen(65535).Nillable().Optional(),
	}
}

// Edges of the CodingSubmissionStaffData.
func (CodingSubmissionStaffData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coding_submission", CodingSubmission.Type).Required().Unique().Annotations(entgql.Bind()),
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
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
	}
}
