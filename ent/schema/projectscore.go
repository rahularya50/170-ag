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

// ProjectScore holds the schema definition for the ProjectScore entity.
type ProjectScore struct {
	ent.Schema
}

// Fields of the ProjectScore.
func (ProjectScore) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("case_id"),
		field.Float("score"),
	}
}

// Edges of the ProjectScore.
func (ProjectScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", ProjectTeam.Type).Unique().Ref("scores").Annotations(entgql.Bind()),
	}
}

func (ProjectScore) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (ProjectScore) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("team"),
		index.Fields("case_id").Edges("team").Unique(),
	}
}

func (ProjectScore) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.GradescopeProjectSubmissionAccessToken),
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.CloudflareCacheAccessToken),
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
	}
}
