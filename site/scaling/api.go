package scaling

import (
	ent "170-ag/ent/generated"
	"170-ag/proto/schemas"
	"170-ag/site"
	"context"
)

type ExternalScalerServer struct {
	schemas.UnimplementedExternalScalerServer
	Client *ent.Client
}

func (s *ExternalScalerServer) IsActive(ctx context.Context, _ *schemas.ScaledObjectRef) (*schemas.IsActiveResponse, error) {
	count, err := site.NumEnqueuedSubmissions(ctx, s.Client)
	if err != nil {
		return nil, err
	}
	return &schemas.IsActiveResponse{Result: count > 0}, nil
}

func (*ExternalScalerServer) GetMetricSpec(context.Context, *schemas.ScaledObjectRef) (*schemas.GetMetricSpecResponse, error) {
	return &schemas.GetMetricSpecResponse{
		MetricSpecs: []*schemas.MetricSpec{{
			MetricName: "queueLength",
			TargetSize: 1,
		}},
	}, nil
}

func (s *ExternalScalerServer) GetMetrics(ctx context.Context, req *schemas.GetMetricsRequest) (*schemas.GetMetricsResponse, error) {
	count, err := site.NumEnqueuedSubmissions(ctx, s.Client)
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
