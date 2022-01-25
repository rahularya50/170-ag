package judge

import (
	"170-ag/proto/schemas"
	"context"
	"os"
)

func PushToGradingQueue(ctx context.Context) error {
	conn, err := NewConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	response, err := loadGradingProto()
	if err != nil {
		return err
	}
	response.Token = os.Getenv("SCALER_TOKEN")

	client := schemas.NewJudgingServerClient(conn)

	_, err = client.SubmitGradingResponse(ctx, response)
	if err != nil {
		return err
	}

	return nil
}
