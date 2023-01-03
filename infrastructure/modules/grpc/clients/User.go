package grpcClient

import (
	"go-boilerplate/api/user"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/modules/grpc/clients/actions/user"
	"go-boilerplate/infrastructure/modules/grpc/protos/build"
	"go-boilerplate/infrastructure/modules/grpc/types"
	"google.golang.org/grpc"
)

type userClient struct {
	conn proto.UserClient
}

var UserClient userClient

func ConfigureUserClient(grpc grpc.ClientConnInterface) {
	connection := proto.NewUserClient(grpc)
	UserClient = userClient{connection}
}

func CreateUserSQL(params *userApi.CreateUserParams) grpcTypes.ClientResponse[userEntity.User] {
	return userGrpcClientActions.CreateUserSQLAction(UserClient.conn, params)
}

func UpdateUserSQL(params *userApi.UpdateUserParams) grpcTypes.ClientResponse[userEntity.User] {
	return userGrpcClientActions.UpdateUserSQLAction(UserClient.conn, params)
}

func GetUserSQL(params *userApi.GetUserParams) grpcTypes.ClientResponse[userEntity.User] {
	return userGrpcClientActions.GetUserSQLAction(UserClient.conn, params)
}
