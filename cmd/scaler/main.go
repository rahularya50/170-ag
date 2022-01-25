package main

import (
	"170-ag/ent/generated/codingsubmissionstaffdata"
	"170-ag/ent/generated/codingtestcase"
	_ "170-ag/ent/generated/runtime"
	"170-ag/privacyrules"
	"170-ag/proto/schemas"
	"170-ag/site"
	"context"
	"crypto/rand"
	"log"
	"math"
	"math/big"
	"net"

	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingsubmission"

	"google.golang.org/grpc"
)

type ScalerServer struct {
	schemas.UnimplementedExternalScalerServer
	schemas.UnimplementedJudgingServerServer
	client *ent.Client
}

func withJudgingContext(ctx context.Context, client *ent.Client) context.Context {
	ctx = privacyrules.NewContextWithAccessToken(ctx, privacyrules.JudgeScalingServerAccessToken)
	ctx = site.ContextWithEntClient(ctx, client)
	return ctx
}

func countEnqueuedSubmissions(c context.Context, client *ent.Client) (int, error) {
	return client.CodingSubmission.
		Query().
		Where(codingsubmission.StatusEQ(codingsubmission.StatusQueued)).
		Count(c)
}

func dequeueSubmission(c context.Context, client *ent.Client) (*ent.CodingSubmission, *ent.CodingSubmissionStaffData, error) {
	tx, err := client.Tx(c)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	submission, err := tx.CodingSubmission.
		Query().
		Where(codingsubmission.StatusEQ(codingsubmission.StatusQueued)).
		First(c)
	if err != nil {
		return nil, nil, err
	}
	nonce, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return nil, nil, err
	}
	staff_data, err := submission.StaffData(c)
	if err != nil {
		return nil, nil, err
	}
	staff_data, err = staff_data.Update().SetExecutionID(nonce.Int64()).Save(c)
	if err != nil {
		return nil, nil, err
	}
	submission, err = submission.Update().SetStatus(codingsubmission.StatusRunning).Save(c)
	if err != nil {
		return nil, nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, nil, err
	}
	return submission.Unwrap(), staff_data.Unwrap(), nil
}

func (s *ScalerServer) IsActive(ctx context.Context, _ *schemas.ScaledObjectRef) (*schemas.IsActiveResponse, error) {
	ctx = withJudgingContext(ctx, s.client)
	count, err := countEnqueuedSubmissions(ctx, s.client)
	if err != nil {
		return nil, err
	}
	return &schemas.IsActiveResponse{Result: count > 0}, nil
}

func (*ScalerServer) GetMetricSpec(context.Context, *schemas.ScaledObjectRef) (*schemas.GetMetricSpecResponse, error) {
	return &schemas.GetMetricSpecResponse{
		MetricSpecs: []*schemas.MetricSpec{{
			MetricName: "queueLength",
			TargetSize: 1,
		}},
	}, nil
}

func (s *ScalerServer) GetMetrics(ctx context.Context, req *schemas.GetMetricsRequest) (*schemas.GetMetricsResponse, error) {
	ctx = withJudgingContext(ctx, s.client)
	count, err := countEnqueuedSubmissions(ctx, s.client)
	if err != nil {
		return nil, err
	}
	return &schemas.GetMetricsResponse{
		MetricValues: []*schemas.MetricValue{{
			MetricName:  "queueLength",
			MetricValue: int64(count),
		}},
	}, nil
}

func (s *ScalerServer) GetJudgingRequest(ctx context.Context, _ *schemas.GetJudgingRequestParams) (*schemas.JudgingRequest, error) {
	ctx = withJudgingContext(ctx, s.client)
	submission, staff_data, err := dequeueSubmission(ctx, s.client)
	if err != nil {
		return nil, err
	}
	return &schemas.JudgingRequest{
		Code:    submission.Code,
		Input:   staff_data.Input,
		IdNonce: uint64(*staff_data.ExecutionID),
	}, nil
}

func (s *ScalerServer) SubmitGradingResponse(ctx context.Context, response *schemas.GradingResponse) (*schemas.GradingReply, error) {
	ctx = withJudgingContext(ctx, s.client)
	nonce := int64(response.IdNonce)
	tx, err := s.client.Tx(ctx)
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

	results := site.ScoreOutput(test_cases, response.Stdout)
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

func main() {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		panic(err)
	}
	client, err := site.GetEntClient()
	if err != nil {
		panic(err)
	}
	schemas.RegisterExternalScalerServer(grpcServer, &ScalerServer{client: client})
	schemas.RegisterJudgingServerServer(grpcServer, &ScalerServer{client: client})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
