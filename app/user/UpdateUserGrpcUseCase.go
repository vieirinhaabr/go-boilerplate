package userUseCase

import (
	userApi "go-boilerplate/api/user"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/types"
	grpcClient "go-boilerplate/infrastructure/modules/grpc/clients"
)

func UpdateUserGrpcUseCase(params *userApi.UpdateUserParams) *types.UseCaseResponse[userEntity.User] {
	response := grpcClient.UpdateUserSQL(params)

	if response.Error != nil {
		return &types.UseCaseResponse[userEntity.User]{
			ErrorCode: &response.Error.Code,
			ErrorMsg:  &response.Error.Msg,
		}
	}

	return &types.UseCaseResponse[userEntity.User]{
		Response:  response.Response,
		ErrorCode: nil,
		ErrorMsg:  nil,
	}
}
