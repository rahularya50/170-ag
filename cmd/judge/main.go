package main

import (
	"170-ag/judge"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	var err error

	switch mode := os.Getenv("MODE"); mode {
	case "SUBSCRIBE":
		err = judge.PopFromJudgingQueue(ctx)
	case "JUDGE":
		err = judge.JudgeLoadedRequest(ctx)
	case "PUBLISH":
		err = judge.PushToGradingQueue(ctx)
	default:
		panic(fmt.Sprintf("Unknown MODE %s", mode))
	}
	if err != nil {
		panic(err)
	}
}
