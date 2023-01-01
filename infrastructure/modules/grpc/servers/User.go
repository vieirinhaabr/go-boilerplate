package grpcServer

import (
	"context"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/servers/actions/user"
	"google.golang.org/grpc"
)

type userServer struct {
	proto.UnsafeUserServer
}

var UserServer userServer

func (srv userServer) ConfigureServer(grpc *grpc.Server) {
	proto.RegisterUserServer(grpc, &srv)
}

func (userServer) CreateUserSQL(_ context.Context, r *proto.CreateUserRequest) (*proto.User, error) {
	return userGrpcActions.CreateUserSQLAction(r)
}

func (userServer) UpdateUserSQL(_ context.Context, r *proto.UpdateUserRequest) (*proto.User, error) {
	return userGrpcActions.UpdateUserSQLAction(r)
}

func (userServer) GetUserSQL(_ context.Context, r *proto.GetUserRequest) (*proto.User, error) {
	return userGrpcActions.GetUserSQLAction(r)
}
