package grading

import (
	"context"
	"errors"

	"google.golang.org/grpc"
)

func InterceptErrors(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	out, err := handler(ctx, req)
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	return out, nil
}
