package judge

import (
	"170-ag/proto/schemas"
	"context"
	"os/exec"
)

func JudgeLoadedRequest(ctx context.Context) error {
	request, err := loadJudgingProto()
	if err != nil {
		return err
	}
	cmd := exec.Command("python", "-c", request.GetCode())
	stdout, _ := cmd.Output() // TODO: handle err, stderr, inputs
	response := &schemas.GradingResponse{
		IdNonce: request.IdNonce, Stdout: string(stdout),
	}
	err = storeGradingProto(response)
	if err != nil {
		return err
	}
	return nil
}
