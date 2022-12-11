package userUseCase

import (
	"go-boilerplate/api/user"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
)

func CreateUserAction(params *userApi.CreateUserParams) *userApi.CreateUserResponse {
	user := userEntity.Create(params.Email, params.Name)
	err := userRepo.Repo.Create(&user)
	if err != nil {
		code := errors.Internal
		msg := "Error when saving user"
		return &userApi.CreateUserResponse{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	return &userApi.CreateUserResponse{
		Response:  &user,
		ErrorCode: nil,
		ErrorMsg:  nil,
	}
}
