package userGrpcActions

import (
	"go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/utils"
)

func GetUserSQLAction(r *proto.GetUserRequest) (*proto.User, error) {
	return grpcUtils.CallUseCaseAsProto(userUseCase.GetUserSQLUseCase, r, new(proto.User))
}
