package schema

import (
	"170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/privacyrules"
	"170-ag/site/policy"
	"context"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CodingProblem holds the schema definition for the CodingProblem entity.
type CodingProblem struct {
	ent.Schema
}

// Fields of the CodingProblem.
func (CodingProblem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("Untitled Problem").NotEmpty().MaxLen(128),
		field.Text("statement").Default("This is the problem statement").NotEmpty(),
		field.Text("skeleton").Default(""),
		field.Bool("released").Default(false),
		field.Time("deadline").Default(time.Now),
	}
}

// Edges of the CodingProblem.
func (CodingProblem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("drafts", CodingDraft.Type).
			Ref("coding_problem").
			Annotations(entgql.Bind()),
		edge.To("test_cases", CodingTestCase.Type).
			Annotations(
				entgql.Bind(),
				entsql.Annotation{OnDelete: entsql.Cascade},
			),
		edge.From("submissions", CodingSubmission.Type).
			Ref("coding_problem").
			Annotations(entgql.Bind()),
	}
}

func (CodingProblem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
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
				return privacy.Skip
			}
		}
		return privacy.Allow
	})
}

func (CodingProblem) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacyrules.AllowWithPrivacyAccessToken(privacyrules.JudgeScalingServerAccessToken),
			policy.DenyIfNoViewer(),
			policy.AllowIfViewerIsStaff(),
			allowCodingProblemQueryIfReleased(),
			privacy.AlwaysDenyRule(),
		},
	}
}
