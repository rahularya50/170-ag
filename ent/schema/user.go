package schema

import (
	"170-ag/ent/generated/privacy"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Deny if not set otherwise.
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			// Allow any viewer to read anything.
			privacy.AlwaysAllowRule(),
		},
	}
}
