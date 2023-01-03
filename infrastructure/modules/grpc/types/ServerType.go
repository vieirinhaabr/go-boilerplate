package grpcTypes

import (
	"google.golang.org/grpc"
)

type Server interface {
	ConfigureServer(grpc *grpc.Server)
}
