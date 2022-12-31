package grpcServer

import (
	"context"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	proto.UnsafeUserServer
}

var UserServer userServer

func (srv userServer) ConfigureServer(grpc *grpc.Server) {
	proto.RegisterUserServer(grpc, &srv)
}

func (userServer) CreateUserSQL(context.Context, *proto.CreateUserRequest) (*proto.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSQL not implemented")
}

func (userServer) UpdateUserSQL(context.Context, *proto.UpdateUserRequest) (*proto.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSQL not implemented")
}

func (userServer) GetUserSQL(context.Context, *proto.GetUserRequest) (*proto.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSQL not implemented")
}
