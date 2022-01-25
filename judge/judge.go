package judge

import (
	"170-ag/proto/schemas"
	"context"
	"os"
	"os/exec"
	"strings"
	"time"
)

func JudgeLoadedRequest(ctx context.Context) error {
	request, err := loadJudgingProto()
	if err != nil {
		return err
	}
	cmd_ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*5))
	defer cancel()
	err = os.WriteFile("tmp.py", []byte(request.Code), 0644)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(cmd_ctx, "/bin/sh", "-c", "ulimit -v 500000; python3 tmp.py")
	cmd.Stdin = strings.NewReader(request.Input)
	stdout, err := cmd.Output()
	result := schemas.ExecutionResult_OK
	if cmd_ctx.Err() == context.DeadlineExceeded {
		result = schemas.ExecutionResult_TIME_LIMIT_EXCEEDED
	}
	exit_err, _ := err.(*exec.ExitError)
	var stderr string
	var errorCode string
	if exit_err != nil {
		stderr = string(exit_err.Stderr)
		errorCode = err.Error()
		if strings.Contains(stderr, "memory") {
			result = schemas.ExecutionResult_MEMORY_ERROR
		}
	}
	response := &schemas.GradingResponse{
		IdNonce:   request.IdNonce,
		Stdout:    string(stdout),
		Stderr:    stderr,
		ErrorCode: errorCode,
		Result:    result,
	}
	err = storeGradingProto(response)
	if err != nil {
		return err
	}
	return nil
}
