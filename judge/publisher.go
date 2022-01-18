package judge

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/proto"
)

func PushToGradingQueue(ctx context.Context) error {
	client, err := NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	response, err := loadGradingProto()
	if err != nil {
		return err
	}

	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	result := client.Topic("grading").Publish(ctx, &pubsub.Message{
		Data: []byte(data),
	})

	_, err = result.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
