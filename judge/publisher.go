package judge

import (
	"170-ag/proto/schemas"
	"context"
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

	client := schemas.NewJudgingServerClient(conn)

	_, err = client.SubmitGradingResponse(ctx, response)
	if err != nil {
		return err
	}

	return nil
}
