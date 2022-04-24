package main

import (
	_ "170-ag/ent/generated/runtime"
	"170-ag/proto/schemas"
	"170-ag/site"
	"170-ag/site/grading"
	"170-ag/site/roles"
	"170-ag/site/scaling"

	"google.golang.org/grpc"
)

func main() {
	client, err := site.GetEntClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	scalerServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			roles.AddGradingContext(client),
			grading.InterceptErrors,
		),
	)
	schemas.RegisterExternalScalerServer(scalerServer, &scaling.ExternalScalerServer{Client: client})
	go site.ServeOnPort(scalerServer, 6001)

	gradingServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grading.InterceptErrors,
		),
	)
	schemas.RegisterJudgingServerServer(gradingServer, &grading.JudgingServer{Client: client})
	site.ServeOnPort(gradingServer, 6000)
}
