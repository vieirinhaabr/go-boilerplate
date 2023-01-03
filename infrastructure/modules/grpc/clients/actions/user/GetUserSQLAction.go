package userGrpcClientActions

import (
	"go-boilerplate/api/user"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/types"
	"go-boilerplate/infrastructure/modules/grpc/utils"
)

func GetUserSQLAction(conn proto.UserClient, params *userApi.GetUserParams) grpcTypes.ClientResponse[userEntity.User] {
	return grpcUtils.CallGrpcServer[proto.GetUserRequest, userEntity.User](params, conn.GetUserSQL)
}
