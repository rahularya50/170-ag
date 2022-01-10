package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
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

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.UserMutationRuleFunc(func(ctx context.Context, m *generated.UserMutation) error {
				if !m.Op().Is(ent.OpCreate) {
					return privacy.Skip
				}
				email, ok := site.EmailFromContext(ctx)
				if !ok {
					return privacy.Deny
				}
				new_email, email_exists := m.Email()
				if !email_exists || email != new_email {
					return privacy.Deny
				}
				return privacy.Allow
			}),
			// TODO: add rule for allowing user modification by themselves
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.QueryPolicy{privacy.UserQueryRuleFunc(func(c context.Context, uq *generated.UserQuery) error {
				viewer, ok := site.ViewerFromContext(c)
				if !ok {
					return privacy.Deny
				}
				// check what the query resolves to
				allow_ctx := privacy.DecisionContext(c, privacy.Allow)
				id, err := uq.Clone().OnlyID(allow_ctx)
				if err == nil && id == viewer.ID {
					return privacy.Allow
				}
				return privacy.Skip
			})},
			privacy.AlwaysDenyRule(),
		},
	}
}
