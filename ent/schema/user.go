package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"170-ag/site"
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}

func allowUserQueryIfIdMatchesViewer() privacy.UserQueryRuleFunc {
	return privacy.UserQueryRuleFunc(func(c context.Context, uq *generated.UserQuery) error {
		viewer, _ := site.ViewerFromContext(c)
		// check what the query resolves to
		allow_ctx := privacy.DecisionContext(c, privacy.Allow)
		id, err := uq.Clone().OnlyID(allow_ctx)
		if err == nil && id == viewer.ID {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

func allowUserMutateIfEmailMatchesContext() privacy.UserMutationRuleFunc {
	return privacy.UserMutationRuleFunc(func(ctx context.Context, m *generated.UserMutation) error {
		if !m.Op().Is(ent.OpCreate) {
			return privacy.Skip
		}
		email, ok := site.EmailFromContext(ctx)
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
			privacyrules.DenyIfNoViewer(),
			allowUserQueryIfIdMatchesViewer(),
			privacy.AlwaysDenyRule(),
		},
	}
}
