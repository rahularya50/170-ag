package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/codingtestcasedata"
	"170-ag/ent/generated/user"
	"170-ag/privacyrules"
	resolvers "170-ag/resolvers/generated"
	model "170-ag/schema/generated"
	"170-ag/site"
	"170-ag/site/web"
	"context"
	"fmt"
	"time"
)

func (r *codingProblemResolver) MyDraft(ctx context.Context, obj *ent.CodingProblem) (*ent.CodingDraft, error) {
	viewer, ok := web.ViewerFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("viewer not found")
	}
	return obj.QueryDrafts().Where(codingdraft.HasAuthorWith(user.ID(viewer.ID))).Only(ctx)
}

func (r *codingProblemResolver) MySubmissions(ctx context.Context, obj *ent.CodingProblem, after *ent.Cursor, first *int, before *ent.Cursor, last *int, isValidation *bool) (*ent.CodingSubmissionConnection, error) {
	viewer, ok := web.ViewerFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("viewer not found")
	}
	query := obj.QuerySubmissions().
		Where(codingsubmission.HasAuthorWith(user.ID(viewer.ID))).
		Order(ent.Desc(codingsubmission.FieldCreateTime))

	if isValidation != nil {
		query = query.Where(codingsubmission.IsValidation(*isValidation))
	}

	return query.Paginate(ctx, after, first, before, last)
}

func (r *codingProblemResolver) AllSubmissions(ctx context.Context, obj *ent.CodingProblem, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.CodingSubmissionConnection, error) {
	return obj.QuerySubmissions().Paginate(ctx, after, first, before, last)
}

func (r *codingProblemResolver) MyDeadline(ctx context.Context, obj *ent.CodingProblem) (*time.Time, error) {
	viewer, ok := web.ViewerFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("viewer not found")
	}
	return site.GetSubmissionDeadlineForStudent(ctx, obj, viewer)
}

func (r *codingProblemResolver) ExtensionRoster(ctx context.Context, obj *ent.CodingProblem) (string, error) {
	return site.ExportProblemExtensions(ctx, obj)
}

func (r *codingProblemResolver) ScoreRoster(ctx context.Context, obj *ent.CodingProblem) (string, error) {
	return site.ExportStudentScores(ctx, obj)
}

func (r *codingSubmissionStaffDataResolver) ExecutionID(ctx context.Context, obj *ent.CodingSubmissionStaffData) (*string, error) {
	if obj.ExecutionID == nil {
		return nil, nil
	}
	s := fmt.Sprint(*obj.ExecutionID)
	return &s, nil
}

func (r *codingTestCaseResolver) ExpandedData(ctx context.Context, obj *ent.CodingTestCase) (*ent.CodingTestCaseData, error) {
	if obj.Visibility == codingtestcase.VisibilityExpanded {
		return obj.QueryData().Only(ctx)
	}
	return nil, nil
}

func (r *mutationResolver) NewUser(ctx context.Context, name string) (*ent.User, error) {
	return r.client.User.Create().SetName(name).Save(ctx)
}

func (r *mutationResolver) NewProblem(ctx context.Context, input model.CodingProblemInput) (*ent.CodingProblem, error) {
	return r.client.CodingProblem.
		Create().
		SetNillableName(input.Name).
		SetNillableStatement(input.Statement).
		SetNillableSkeleton(input.Skeleton).
		SetNillableReleased(input.Released).
		Save(ctx)
}

func (r *mutationResolver) SaveDraft(ctx context.Context, input model.CodingDraftInput) (*ent.CodingDraft, error) {
	viewer, ok := web.ViewerFromContext(ctx)
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

func (r *mutationResolver) CreateSubmission(ctx context.Context, input model.CodingSubmissionInput) (*ent.CodingSubmission, error) {
	viewer, ok := web.ViewerFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("viewer not found")
	}

	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	submission, err := tx.CodingSubmission.Create().
		SetAuthor(viewer).
		SetCode(input.Code).
		SetCodingProblemID(input.ProblemID).
		SetIsValidation(input.IsValidation).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	submission_ctx := privacyrules.NewContextWithAccessToken(ctx, privacyrules.SubmissionEnqueuingAccessToken)
	if !input.IsValidation {
		// authorize viewer to see hidden test cases
		submission_ctx = privacyrules.NewContextWithAccessToken(submission_ctx, privacyrules.FullSubmissionTestCaseAccessToken)
	}
	test_case_data, err := site.QueryTestCases(submission).QueryData().Order(ent.Asc(codingtestcasedata.FieldCreateTime)).All(submission_ctx)

	if err != nil {
		return nil, err
	}

	inputData, err := site.GenerateInput(test_case_data)
	if err != nil {
		return nil, err
	}

	err = tx.CodingSubmissionStaffData.Create().
		SetCodingSubmission(submission).
		SetInput(inputData).
		Exec(submission_ctx)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return submission.Unwrap(), nil
}

func (r *mutationResolver) UpdateProblem(ctx context.Context, input model.UpdateCodingProblemInput) (*ent.CodingProblem, error) {
	problem, err := r.client.CodingProblem.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return problem.Update().
		SetNillableName(input.Problem.Name).
		SetNillableStatement(input.Problem.Statement).
		SetNillableSkeleton(input.Problem.Skeleton).
		SetNillableReleased(input.Problem.Released).
		SetNillableDeadline(input.Problem.Deadline).
		Save(ctx)
}

func (r *mutationResolver) AddTestCase(ctx context.Context, input model.CreateTestCaseInput) (*ent.CodingTestCase, error) {
	tx, err := r.client.Tx(ctx)
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	test_case, err := tx.CodingTestCase.Create().SetCodingProblemID(input.ProblemID).Save(ctx)
	if err != nil {
		return nil, err
	}
	err = tx.CodingTestCaseData.Create().SetTestCase(test_case).Exec(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return test_case.Unwrap(), nil
}

func (r *mutationResolver) UpdateTestCase(ctx context.Context, input model.UpdateTestCaseInput) (*ent.CodingTestCase, error) {
	tx, err := r.client.Tx(ctx)
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	test_case, err := tx.CodingTestCase.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	data, err := test_case.QueryData().Only(ctx)
	if err != nil {
		return nil, err
	}
	err = data.Update().
		SetNillableInput(input.Input).
		SetNillableOutput(input.Output).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	test_case, err = test_case.Update().
		SetNillablePoints(input.Points).
		SetNillableVisibility(input.Visibility).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return test_case.Unwrap(), nil
}

func (r *mutationResolver) DeleteTestCase(ctx context.Context, input model.DeleteTestCaseInput) (*ent.CodingProblem, error) {
	problem, err := r.client.CodingTestCase.Query().Where(codingtestcase.ID(input.ID)).QueryCodingProblem().Only(ctx)
	if err != nil {
		return nil, err
	}
	err = r.client.CodingTestCase.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func (r *mutationResolver) SetProblemExtensions(ctx context.Context, input model.ExtensionsInput) (*ent.CodingProblem, error) {
	problem, err := r.client.CodingProblem.Query().Where(codingproblem.ID(input.ProblemID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	err = site.SetProblemExtensions(ctx, r.client, problem, input.Roster)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*ent.User, error) {
	viewer, ok := web.ViewerFromContext(ctx)
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

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last)
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

func (r *queryResolver) TestCase(ctx context.Context, id int) (*ent.CodingTestCase, error) {
	return r.client.CodingTestCase.Get(ctx, id)
}

func (r *queryResolver) CodingSubmission(ctx context.Context, id int) (*ent.CodingSubmission, error) {
	return r.client.CodingSubmission.Get(ctx, id)
}

// CodingProblem returns resolvers.CodingProblemResolver implementation.
func (r *Resolver) CodingProblem() resolvers.CodingProblemResolver { return &codingProblemResolver{r} }

// CodingSubmissionStaffData returns resolvers.CodingSubmissionStaffDataResolver implementation.
func (r *Resolver) CodingSubmissionStaffData() resolvers.CodingSubmissionStaffDataResolver {
	return &codingSubmissionStaffDataResolver{r}
}

// CodingTestCase returns resolvers.CodingTestCaseResolver implementation.
func (r *Resolver) CodingTestCase() resolvers.CodingTestCaseResolver {
	return &codingTestCaseResolver{r}
}

// Mutation returns resolvers.MutationResolver implementation.
func (r *Resolver) Mutation() resolvers.MutationResolver { return &mutationResolver{r} }

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type codingProblemResolver struct{ *Resolver }
type codingSubmissionStaffDataResolver struct{ *Resolver }
type codingTestCaseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
