package judge

import (
	"170-ag/proto/schemas"
	"context"
	"os/exec"
	"strings"
)

func JudgeLoadedRequest(ctx context.Context) error {
	request, err := loadJudgingProto()
	if err != nil {
		return err
	}
	cmd := exec.Command("python3", "-c", request.GetCode())
	cmd.Stdin = strings.NewReader(request.Input)
	stdout, err := cmd.Output()
	exit_err, _ := err.(*exec.ExitError)
	var stderr string
	var errorCode string
	if exit_err != nil {
		stderr = string(exit_err.Stderr)
		errorCode = err.Error()
	}
	response := &schemas.GradingResponse{
		IdNonce:   request.IdNonce,
		Stdout:    string(stdout),
		Stderr:    stderr,
		ErrorCode: errorCode,
	}
	err = storeGradingProto(response)
	if err != nil {
		return err
	}
	return nil
}
