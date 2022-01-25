package judge

import (
	"170-ag/proto/schemas"
	"context"
	"os"
)

func PopFromJudgingQueue(ctx context.Context) error {
	conn, err := NewConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := schemas.NewJudgingServerClient(conn)
	judgingRequest, err := client.GetJudgingRequest(ctx, &schemas.GetJudgingRequestParams{
		Token: os.Getenv("SCALER_TOKEN"),
	})
	if err != nil {
		return err
	}

	// store the proto locally for the next container in the chain
	err = storeJudgingProto(judgingRequest)
	if err != nil {
		return err
	}
	return nil
}
