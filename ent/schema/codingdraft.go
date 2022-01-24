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
	"entgo.io/ent/schema/mixin"
)

// CodingDraft holds the schema definition for the CodingDraft entity.
type CodingDraft struct {
	ent.Schema
}

// Fields of the CodingDraft.
func (CodingDraft) Fields() []ent.Field {
	return []ent.Field{
		field.Text("code"),
	}
}

// Edges of the CodingDraft.
func (CodingDraft) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Required().Unique().Annotations(entgql.Bind()),
		edge.To("coding_problem", CodingProblem.Type).Required().Unique().Annotations(entgql.Bind()),
	}
}

func (CodingDraft) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (CodingDraft) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("author", "coding_problem").Unique(),
	}
}

func (CodingDraft) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacyrules.FilterToViewerID(func(c context.Context, f privacy.Filter, user_id int) error {
				f.(*generated.CodingDraftFilter).WhereHasAuthorWith(user.ID(user_id))
				return privacy.Allow
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacyrules.AllowQueryIfIDsMatchViewer(func(c context.Context, q ent.Query) ([]int, error) {
				return q.(*generated.CodingDraftQuery).Clone().QueryAuthor().IDs(c)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
