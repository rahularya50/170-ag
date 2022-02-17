package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/site/policy"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// CodingExtension holds the schema definition for the CodingExtension entity.
type CodingExtension struct {
	ent.Schema
}

// Fields of the CodingExtension.
func (CodingExtension) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deadline"),
	}
}

// Edges of the CodingExtension.
func (CodingExtension) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", User.Type).Required().Unique().Annotations(entgql.Bind()),
		edge.To("coding_problem", CodingProblem.Type).Required().Unique().Annotations(entgql.Bind()),
	}
}

func (CodingExtension) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (CodingExtension) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("student", "coding_problem").Unique(),
		index.Edges("student"),
		index.Edges("coding_problem"),
	}
}

func (CodingExtension) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			policy.AllowQueryIfIDsMatchViewer(func(c context.Context, q generated.Query) ([]int, error) {
				return q.(*generated.CodingExtensionQuery).Clone().QueryStudent().IDs(c)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
