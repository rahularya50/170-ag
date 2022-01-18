package judge

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func NewClient(ctx context.Context) (*pubsub.Client, error) {
	return pubsub.NewClient(ctx, "formidable-gate-337712")
}
