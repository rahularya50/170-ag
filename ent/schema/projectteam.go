package schema

import (
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"170-ag/site/policy"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// ProjectTeam holds the schema definition for the ProjectTeam entity.
type ProjectTeam struct {
	ent.Schema
}

// Fields of the ProjectScore.
func (ProjectTeam) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("team_id").Unique(),
		field.String("name").NotEmpty().MaxLen(100).Unique(),
	}
}

// Edges of the ProjectTeam.
func (ProjectTeam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("scores", ProjectScore.Type).Required().Annotations(
			entgql.Bind(),
			entsql.Annotation{OnDelete: entsql.Cascade},
		),
	}
}

func (ProjectTeam) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (ProjectTeam) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("team_id").Unique(),
	}
}

func (ProjectTeam) Policy() ent.Policy {
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
