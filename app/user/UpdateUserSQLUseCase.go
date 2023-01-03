package userUseCase

import (
	userApi "go-boilerplate/api/user"
	userValidator "go-boilerplate/app/user/validators"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/global/types"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
)

func UpdateUserSQLUseCase(params *userApi.UpdateUserParams) *types.UseCaseResponse[userEntity.User] {
	err := userValidator.UpdateUserValidator(params)
	if err != nil {
		code := errors.Validation
		msg := err.Error()
		return &types.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	user, err := gormUserRepo.Repo.GetById(params.ID)
	if err != nil {
		code := errors.Internal
		msg := err.Error()
		return &types.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	user.Update(params.Email, params.Name)
	err = gormUserRepo.Repo.Update(&user)
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
