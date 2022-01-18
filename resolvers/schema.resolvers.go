package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/user"
	resolvers "170-ag/resolvers/generated"
	model "170-ag/schema/generated"
	"170-ag/site"
	"context"
	"fmt"
)

func (r *codingProblemResolver) MyDraft(ctx context.Context, obj *ent.CodingProblem) (*ent.CodingDraft, error) {
	viewer, ok := site.ViewerFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("viewer not found")
	}
	return obj.QueryDrafts().Where(codingdraft.HasAuthorWith(user.ID(viewer.ID))).Only(ctx)
}

func (r *codingProblemResolver) MySubmissions(ctx context.Context, obj *ent.CodingProblem, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.CodingSubmissionConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

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

func (r *mutationResolver) SaveDraft(ctx context.Context, input *model.CodingDraftInput) (*ent.CodingDraft, error) {
	viewer, ok := site.ViewerFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("viewer not found")
	}
	coding_draft_id, err := r.client.CodingDraft.Create().
		SetAuthor(viewer).
		SetCode(input.Code).
		SetCodingProblemID(input.ProblemID).
		OnConflictColumns(codingdraft.CodingProblemColumn, codingdraft.AuthorColumn).
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.CodingDraft.Get(ctx, coding_draft_id)
}

func (r *mutationResolver) CreateSubmission(ctx context.Context, input *model.CodingSubmissionInput) (*ent.CodingSubmission, error) {
	viewer, ok := site.ViewerFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("viewer not found")
	}
	submission, err := r.client.CodingSubmission.Create().
		SetAuthor(viewer).
		SetCode(input.Code).
		SetCodingProblemID(input.ProblemID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	err = site.EnqueueSubmission(ctx, submission)
	if err != nil {
		return nil, err
	}
	return submission, nil
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

// CodingProblem returns resolvers.CodingProblemResolver implementation.
func (r *Resolver) CodingProblem() resolvers.CodingProblemResolver { return &codingProblemResolver{r} }

// Mutation returns resolvers.MutationResolver implementation.
func (r *Resolver) Mutation() resolvers.MutationResolver { return &mutationResolver{r} }

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type codingProblemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
