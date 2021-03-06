package judge

import (
	"170-ag/proto/schemas"
	"context"
	"os"
	"os/exec"
	"strings"
	"time"
)

const OUTPUT_MAXLEN = 7000000
const ERR_MAXLEN = 10000

func truncate(s string, maxLen int) string {
	if len(s) < maxLen {
		return s
	}
	return s[:maxLen]
}

func JudgeLoadedRequest(ctx context.Context) error {
	request, err := loadJudgingProto()
	if err != nil {
		return err
	}
	cmd_ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*15))
	defer cancel()
	err = os.WriteFile("tmp.py", []byte(request.Code), 0644)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(cmd_ctx, "/bin/sh", "-c", "ulimit -v 500000 -t 20; python3 tmp.py")
	stderrBuilder := new(strings.Builder)
	cmd.Stderr = stderrBuilder
	cmd.Stdin = strings.NewReader(request.Input)
	stdout, err := cmd.Output()
	result := schemas.ExecutionResult_OK
	if cmd_ctx.Err() == context.DeadlineExceeded {
		result = schemas.ExecutionResult_TIME_LIMIT_EXCEEDED
	}
	exit_err, _ := err.(*exec.ExitError)
	stderr := stderrBuilder.String()
	var errorCode string
	if exit_err != nil {
		errorCode = err.Error()
		switch errStr := strings.ToLower(stderr); {
		case strings.Contains(errStr, "memory"):
			result = schemas.ExecutionResult_MEMORY_ERROR
		case strings.Contains(errStr, "time"):
			result = schemas.ExecutionResult_TIME_LIMIT_EXCEEDED
		}
	}
	response := &schemas.GradingResponse{
		IdNonce:   request.IdNonce,
		Stdout:    truncate(string(stdout), OUTPUT_MAXLEN),
		Stderr:    truncate(stderr, ERR_MAXLEN),
		ErrorCode: truncate(errorCode, ERR_MAXLEN),
		Result:    result,
	}
	err = storeGradingProto(response)
	if err != nil {
		return err
	}
	return nil
}
