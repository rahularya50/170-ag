package judge

import "google.golang.org/grpc"

func NewConn() (*grpc.ClientConn, error) {
	return grpc.Dial("judge-scaler.default.svc.cluster.local:6000", grpc.WithInsecure())
}
