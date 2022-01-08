package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	ent "170-ag/ent/generated"
	resolvers "170-ag/resolvers/generated"
	"context"
)

func (r *mutationResolver) NewUser(ctx context.Context, name *string) (*ent.User, error) {
	return r.client.User.Create().SetAge(10).SetName(*name).Save(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Mutation returns resolvers.MutationResolver implementation.
func (r *Resolver) Mutation() resolvers.MutationResolver { return &mutationResolver{r} }

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
