package roles

import (
	ent "170-ag/ent/generated"
	"170-ag/privacyrules"
	"170-ag/site"
	"context"

	"google.golang.org/grpc"
)

func WithGradingContext(ctx context.Context, client *ent.Client) context.Context {
	ctx = privacyrules.NewContextWithAccessToken(ctx, privacyrules.JudgeScalingServerAccessToken)
	ctx = site.ContextWithEntClient(ctx, client)
	return ctx
}

func AddGradingContext(client *ent.Client) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx = privacyrules.NewContextWithAccessToken(ctx, privacyrules.JudgeScalingServerAccessToken)
		ctx = site.ContextWithEntClient(ctx, client)
		return handler(ctx, req)
	}
}
