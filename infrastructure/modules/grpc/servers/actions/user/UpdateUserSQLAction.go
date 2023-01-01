package userGrpcActions

import (
	"go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/utils"
)

func UpdateUserSQLAction(r *proto.UpdateUserRequest) (*proto.User, error) {
	return grpcUtils.CallUseCaseAsProto("UpdateUserSQLAction", userUseCase.UpdateUserSQLUseCase, r, new(proto.User))
}
