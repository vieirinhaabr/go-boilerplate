package userUseCase

import (
	"go-boilerplate/api"
	userApi "go-boilerplate/api/user"
	userValidator "go-boilerplate/app/user/validators"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	redisUserRepo "go-boilerplate/infrastructure/modules/redis/repos/user"
)

func GetUserRedisUseCase(params *userApi.GetUserByNameParams) *api.UseCaseResponse[userEntity.User] {
	err := userValidator.GetUserByNameValidator(params)
	if err != nil {
		code := errors.Validation
		msg := err.Error()
		return &api.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	user, err := redisUserRepo.Repo.Get(params.Name)
	if err != nil {
		code := errors.Internal
		msg := err.Error()
		return &api.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	return &api.UseCaseResponse[userEntity.User]{
		Response:  user,
		ErrorCode: nil,
		ErrorMsg:  nil,
	}
}
