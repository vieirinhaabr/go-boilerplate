package userGrpcClientActions

import (
	"go-boilerplate/api/user"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/types"
	"go-boilerplate/infrastructure/modules/grpc/utils"
)

func UpdateUserSQLAction(conn proto.UserClient, params *userApi.UpdateUserParams) grpcTypes.ClientResponse[userEntity.User] {
	return grpcUtils.CallGrpcServer[proto.UpdateUserRequest, userEntity.User](params, conn.UpdateUserSQL)
}
