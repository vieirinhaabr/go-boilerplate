package userGrpcClientActions

import (
	"go-boilerplate/api/user"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/types"
	"go-boilerplate/infrastructure/modules/grpc/utils"
)

func CreateUserSQLAction(conn proto.UserClient, params *userApi.CreateUserParams) grpcTypes.ClientResponse[userEntity.User] {
	return grpcUtils.CallGrpcServer[proto.CreateUserRequest, userEntity.User](params, conn.CreateUserSQL)
}
