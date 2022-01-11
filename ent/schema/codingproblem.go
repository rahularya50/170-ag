package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CodingProblem holds the schema definition for the CodingProblem entity.
type CodingProblem struct {
	ent.Schema
}

// Fields of the CodingProblem.
func (CodingProblem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(128),
		field.Text("statement").Default("This is the problem statement").NotEmpty(),
		field.Bool("released").Default(false),
	}
}

// Edges of the CodingProblem.
func (CodingProblem) Edges() []ent.Edge {
	return nil
}

func allowCodingProblemQueryIfReleased() privacy.CodingProblemQueryRuleFunc {
	return privacy.CodingProblemQueryRuleFunc(func(c context.Context, cpq *generated.CodingProblemQuery) error {
		allow_ctx := privacy.DecisionContext(c, privacy.Allow)
		problems, err := cpq.Clone().All(allow_ctx)
		if err != nil {
			return privacy.Skip
		}
		for _, problem := range problems {
			if !problem.Released {
				return privacy.Deny
			}
		}
		return privacy.Allow
	})
}

func (CodingProblem) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.DenyIfNoViewer(),
			privacyrules.AllowIfViewerIsStaff(),
			allowCodingProblemQueryIfReleased(),
			privacy.AlwaysDenyRule(),
		},
	}
}
