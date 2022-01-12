package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingproblem"
	resolvers "170-ag/resolvers/generated"
	model "170-ag/schema/generated"
	"170-ag/site"
	"context"
	"fmt"
)

func (r *mutationResolver) NewUser(ctx context.Context, name *string) (*ent.User, error) {
	return r.client.User.Create().SetName(*name).Save(ctx)
}

func (r *mutationResolver) NewProblem(ctx context.Context, input *model.CodingProblemInput) (*ent.CodingProblem, error) {
	return r.client.CodingProblem.
		Create().
		SetName(input.Name).
		SetStatement(input.Statement).
		SetReleased(input.Released).
		Save(ctx)
}

func (r *mutationResolver) SaveDraft(ctx context.Context, input *model.CodingDraftInput) (*model.CodingDraft, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Viewer(ctx context.Context) (*ent.User, error) {
	viewer, ok := site.ViewerFromContext(ctx)
	if ok {
		return viewer, nil
	} else {
		return nil, nil
	}
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
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

// Mutation returns resolvers.MutationResolver implementation.
func (r *Resolver) Mutation() resolvers.MutationResolver { return &mutationResolver{r} }

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
