package userUseCase

import (
	"go-boilerplate/api"
	userApi "go-boilerplate/api/user"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
)

func UpdateUserUseCase(params *userApi.UpdateUserParams) *api.UseCaseResponse[userEntity.User] {
	user, err := userRepo.Repo.GetById(params.ID)
	if err != nil {
		code := errors.Internal
		msg := "Error when getting user"
		return &api.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	user.Update(params.Email, params.Name)
	err = userRepo.Repo.Update(&user)
	if err != nil {
		code := errors.Internal
		msg := "Error when saving user"
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
