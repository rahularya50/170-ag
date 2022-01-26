package main

import (
	_ "170-ag/ent/generated/runtime"
	"170-ag/proto/schemas"
	"170-ag/site"
	"170-ag/site/grading"
	"170-ag/site/roles"
	"170-ag/site/scaling"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func serveOnPort(server *grpc.Server, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	client, err := site.GetEntClient()
	if err != nil {
		panic(err)
	}

	scalerServer := grpc.NewServer(
		grpc.UnaryInterceptor(roles.AddGradingContext(client)),
		grpc.UnaryInterceptor(grading.InterceptErrors),
	)
	schemas.RegisterExternalScalerServer(scalerServer, &scaling.ExternalScalerServer{Client: client})
	go serveOnPort(scalerServer, 6001)

	gradingServer := grpc.NewServer()
	schemas.RegisterJudgingServerServer(gradingServer, &grading.JudgingServer{Client: client})
	serveOnPort(gradingServer, 6000)
}
