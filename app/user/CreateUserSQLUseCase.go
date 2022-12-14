package userUseCase

import (
	userApi "go-boilerplate/api/user"
	userValidator "go-boilerplate/app/user/validators"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/global/types"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
)

func CreateUserSQLUseCase(params *userApi.CreateUserParams) *types.UseCaseResponse[userEntity.User] {
	err := userValidator.CreateUserValidator(params)
	if err != nil {
		code := errors.Validation
		msg := err.Error()
		return &types.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	user := userEntity.Create(params.Email, params.Name)
	err = gormUserRepo.Repo.Create(&user)
	if err != nil {
		code := errors.Internal
		msg := err.Error()
		return &types.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	return &types.UseCaseResponse[userEntity.User]{
		Response:  &user,
		ErrorCode: nil,
		ErrorMsg:  nil,
	}
}
