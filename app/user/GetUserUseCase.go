package userUseCase

import (
	"go-boilerplate/api"
	userApi "go-boilerplate/api/user"
	userValidator "go-boilerplate/app/user/validators"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
	mongoUserRepo "go-boilerplate/infrastructure/modules/mongo/repos/user"
)

func GetUserUseCase(params *userApi.GetUserParams) *api.UseCaseResponse[userEntity.User] {
	err := userValidator.GetUserValidator(params)
	if err != nil {
		code := errors.Validation
		msg := err.Error()
		return &api.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	if params.Type == "mongo" {
		user, err := mongoUserRepo.Repo.GetById(params.ID)
		if err != nil {
			code := errors.Internal
			if err.Error() == string(errors.NoItemsFound){
				code = errors.ItemRequestedNotFound
			}

			msg := err.Error()
			return &api.UseCaseResponse[userEntity.User]{
				Response:  nil,
				ErrorCode: &code,
				ErrorMsg:  &msg,
			}
		}

		return &api.UseCaseResponse[userEntity.User]{
			Response:  &user,
			ErrorCode: nil,
			ErrorMsg:  nil,
		}
	}

	user, err := gormUserRepo.Repo.GetById(params.ID)
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
		Response:  &user,
		ErrorCode: nil,
		ErrorMsg:  nil,
	}
}
