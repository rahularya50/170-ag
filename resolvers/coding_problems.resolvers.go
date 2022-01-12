package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingproblem"
	model "170-ag/schema/generated"
	"context"
)

func (r *mutationResolver) NewProblem(ctx context.Context, input *model.CodingProblemInput) (*ent.CodingProblem, error) {
	return r.client.CodingProblem.
		Create().
		SetName(input.Name).
		SetStatement(input.Statement).
		SetReleased(input.Released).
		Save(ctx)
}

func (r *queryResolver) CodingProblem(ctx context.Context, id int) (*ent.CodingProblem, error) {
	return r.client.CodingProblem.Get(ctx, id)
}

func (r *queryResolver) CodingProblems(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, includeUnreleased bool) (*ent.CodingProblemConnection, error) {
	query := r.client.CodingProblem.Query()

	if !includeUnreleased {
		query = query.Where(codingproblem.Released(true))
	}

	return query.Paginate(ctx, after, first, before, last)
}
