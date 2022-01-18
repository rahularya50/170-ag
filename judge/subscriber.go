package judge

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
)

func PopFromJudgingQueue(ctx context.Context) error {
	client, err := NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	sub := client.Subscription("judge")
	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxOutstandingMessages = 1
	sub.ReceiveSettings.NumGoroutines = 1

	// if we're started but for some reason no messages are received, exit without output
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	errorChan := make(chan error, 1)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// there's only ever one unacked message in flight at a time, and we have it

		cancel() // stop the context so more handlers aren't started

		// in the event of downstream failure, a CronJob will
		// retransmit pending submissions on a timer, so it's safe to Ack here
		msg.Ack()

		judgingRequest, err := decodeJudgingProto(msg.Data)
		if err != nil {
			errorChan <- err
			return
		}

		// store the proto locally for the next container in the chain
		err = storeJudgingProto(judgingRequest)
		if err != nil {
			errorChan <- err
			return
		}

		errorChan <- nil
	})

	if err != nil {
		return err
	}

	err = <-errorChan
	return err
}
