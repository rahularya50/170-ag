package main

import (
	_ "170-ag/ent/generated/runtime"
	"170-ag/privacyrules"
	"170-ag/proto/schemas"
	"170-ag/site"
	"context"
	"log"
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

func countEnqueuedSubmissions(c context.Context, client *ent.Client) (int, error) {
	c = privacyrules.NewContextWithAccessToken(c, privacyrules.JudgeScalingServerAccessToken)
	return client.CodingSubmission.
		Query().
		Where(codingsubmission.StatusEQ(codingsubmission.StatusQueued)).
		Count(c)
}

func dequeueSubmission(c context.Context, client *ent.Client) (*ent.CodingSubmission, error) {
	c = privacyrules.NewContextWithAccessToken(c, privacyrules.JudgeScalingServerAccessToken)
	submission, err := client.CodingSubmission.
		Query().
		Where(codingsubmission.StatusEQ(codingsubmission.StatusQueued)).
		First(c)
	if err != nil {
		return nil, err
	}
	submission, err = submission.Update().SetStatus(codingsubmission.StatusRunning).Save(c)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (s *ScalerServer) IsActive(c context.Context, _ *schemas.ScaledObjectRef) (*schemas.IsActiveResponse, error) {
	count, err := countEnqueuedSubmissions(c, s.client)
	if err != nil {
		return nil, err
	}
	return &schemas.IsActiveResponse{Result: count > 0}, nil
}

func (*ScalerServer) GetMetricSpec(context.Context, *schemas.ScaledObjectRef) (*schemas.GetMetricSpecResponse, error) {
	return &schemas.GetMetricSpecResponse{
		MetricSpecs: []*schemas.MetricSpec{{
			MetricName: "JobsPerTask",
			TargetSize: 1,
		}},
	}, nil
}

func (s *ScalerServer) GetMetrics(c context.Context, _ *schemas.GetMetricsRequest) (*schemas.GetMetricsResponse, error) {
	count, err := countEnqueuedSubmissions(c, s.client)
	if err != nil {
		return nil, err
	}
	return &schemas.GetMetricsResponse{
		MetricValues: []*schemas.MetricValue{{
			MetricName:  "JobsPerTask",
			MetricValue: int64(count),
		}},
	}, nil
}

func (s *ScalerServer) GetJudgingRequest(ctx context.Context, _ *schemas.GetJudgingRequestParams) (*schemas.JudgingRequest, error) {
	submission, err := dequeueSubmission(ctx, s.client)
	if err != nil {
		return nil, err
	}
	return &schemas.JudgingRequest{
		Code:    submission.Code,
		Input:   "",
		IdNonce: 1, // TODO FIXME
	}, nil
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
