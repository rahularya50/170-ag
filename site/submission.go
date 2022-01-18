package site

import (
	ent "170-ag/ent/generated"
	"170-ag/judge"
	"170-ag/proto/schemas"
	"context"
	"crypto/rand"
	"log"
	"math"
	"math/big"

	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/proto"
)

func EnqueueSubmission(ctx context.Context, submission *ent.CodingSubmission) error {
	log.Printf("Enqueuing submission")
	client, err := judge.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	nonce, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return err
	}

	request := &schemas.JudgingRequest{
		Code:    submission.Code,
		Input:   "TODO",
		IdNonce: nonce.Uint64(),
	}

	data, err := proto.Marshal(request)
	if err != nil {
		return err
	}

	topic := client.Topic("judging")
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(data),
	})

	_, err = result.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
