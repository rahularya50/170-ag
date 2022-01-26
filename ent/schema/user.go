package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/site/policy"
	"170-ag/site/web"
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().NotEmpty().MaxLen(128).Immutable(),
		field.String("name").Optional().NotEmpty().MaxLen(128),
		field.Bool("is_staff").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("drafts", CodingDraft.Type).Ref("author"),
		edge.From("submissions", CodingSubmission.Type).Ref("author"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}

func allowUserMutateIfEmailMatchesContext() privacy.UserMutationRuleFunc {
	return privacy.UserMutationRuleFunc(func(ctx context.Context, m *generated.UserMutation) error {
		if !m.Op().Is(ent.OpCreate) {
			return privacy.Skip
		}
		email, ok := web.EmailFromContext(ctx)
		if !ok {
			return privacy.Skip
		}
		new_email, email_exists := m.Email()
		if !email_exists || email != new_email {
			return privacy.Skip
		}
		return privacy.Allow
	})
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				allowUserMutateIfEmailMatchesContext(),
				ent.OpCreate,
			),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			policy.AllowQueryIfIDsMatchViewer(func(c context.Context, q generated.Query) ([]int, error) {
				return q.(*generated.UserQuery).Clone().IDs(c)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
