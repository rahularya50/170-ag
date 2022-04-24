package site

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func ServeOnPort(server *grpc.Server, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
