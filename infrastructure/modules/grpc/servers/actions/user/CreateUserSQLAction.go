package userGrpcActions

import (
	"go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/utils"
)

func CreateUserSQLAction(r *proto.CreateUserRequest) (*proto.User, error) {
	return grpcUtils.CallUseCaseAsProto(userUseCase.CreateUserSQLUseCase, r, new(proto.User))
}
