package grading

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingsubmissionstaffdata"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/models"
	"170-ag/proto/schemas"
	"170-ag/site"
	"context"
	"time"
)

type JudgingServer struct {
	schemas.UnimplementedJudgingServerServer
	Client *ent.Client
}

func (s *JudgingServer) GetJudgingRequest(ctx context.Context, req *schemas.GetJudgingRequestParams) (*schemas.JudgingRequest, error) {
	ctx, err := withAuthorizedContext(ctx, s.Client, req.Token)
	if err != nil {
		return nil, err
	}
	err = site.CleanupStalled(ctx, s.Client, time.Minute*5)
	if err != nil {
		return nil, err
	}
	submission, staff_data, err := site.DequeueSubmission(ctx, s.Client)
	if err != nil {
		return nil, err
	}
	return &schemas.JudgingRequest{
		Code:    submission.Code,
		Input:   staff_data.Input,
		IdNonce: uint64(*staff_data.ExecutionID),
	}, nil
}

func (s *JudgingServer) SubmitGradingResponse(ctx context.Context, response *schemas.GradingResponse) (*schemas.GradingReply, error) {
	ctx, err := withAuthorizedContext(ctx, s.Client, response.Token)
	if err != nil {
		return nil, err
	}
	nonce := int64(response.IdNonce)
	tx, err := s.Client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	staff_data, err := tx.CodingSubmissionStaffData.Query().
		Where(codingsubmissionstaffdata.ExecutionID(nonce)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	submission, err := staff_data.CodingSubmission(ctx)
	if err != nil {
		return nil, err
	}
	err = staff_data.Update().
		SetOutput(response.Stdout).
		SetStderr(response.Stderr).
		SetExitError(response.ErrorCode).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	test_cases, err := submission.
		QueryCodingProblem().
		QueryTestCases().
		Order(ent.Asc(codingtestcase.FieldCreateTime)).
		WithData().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var defaultFailResult = models.ResultPresentationError
	switch response.Result {
	case schemas.ExecutionResult_OK:
		if response.ErrorCode != "" {
			defaultFailResult = models.ResultRuntimeError
		}
	case schemas.ExecutionResult_TIME_LIMIT_EXCEEDED:
		defaultFailResult = models.ResultTimeLimitExceeded
	case schemas.ExecutionResult_COMPILE_ERROR:
		defaultFailResult = models.ResultRuntimeError
	case schemas.ExecutionResult_MEMORY_ERROR:
		defaultFailResult = models.ResultMemoryExhausted
	}

	results := site.ScoreOutput(test_cases, response.Stdout, defaultFailResult)
	points := 0
	for _, result := range results.CaseResults {
		points += result.Points
	}
	err = submission.Update().
		SetStatus(codingsubmission.StatusCompleted).
		SetResults(results).
		SetPoints(points).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return &schemas.GradingReply{}, nil
}
