package userUseCase

import (
	"go-boilerplate/api"
	userApi "go-boilerplate/api/user"
	userValidator "go-boilerplate/app/user/validators"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	mongoUserRepo "go-boilerplate/infrastructure/modules/mongo/repos/user"
)

func CreateUserNoSQLUseCase(params *userApi.CreateUserParams) *api.UseCaseResponse[userEntity.User] {
	err := userValidator.CreateUserValidator(params)
	if err != nil {
		code := errors.Validation
		msg := err.Error()
		return &api.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	userMongo := userEntity.Create(params.Email, params.Name)
	err = mongoUserRepo.Repo.Create(&userMongo)
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
		Response:  &userMongo,
		ErrorCode: nil,
		ErrorMsg:  nil,
	}
}