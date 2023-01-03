package userUseCase

import (
	userApi "go-boilerplate/api/user"
	userValidator "go-boilerplate/app/user/validators"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/global/types"
	redisUserRepo "go-boilerplate/infrastructure/modules/redis/repos/user"
)

func CreateUserRedisUseCase(params *userApi.CreateUserParams) *types.UseCaseResponse[userEntity.User] {
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
	err = redisUserRepo.Repo.Set(&user)
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
