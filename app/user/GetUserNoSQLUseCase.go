package userUseCase

import (
	userApi "go-boilerplate/api/user"
	userValidator "go-boilerplate/app/user/validators"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/global/types"
	mongoUserRepo "go-boilerplate/infrastructure/modules/mongo/repos/user"
)

func GetUserNoSQLUseCase(params *userApi.GetUserParams) *types.UseCaseResponse[userEntity.User] {
	err := userValidator.GetUserValidator(params)
	if err != nil {
		code := errors.Validation
		msg := err.Error()
		return &types.UseCaseResponse[userEntity.User]{
			Response:  nil,
			ErrorCode: &code,
			ErrorMsg:  &msg,
		}
	}

	user, err := mongoUserRepo.Repo.GetById(params.ID)
	if err != nil {
		code := errors.Internal
		if err.Error() == string(errors.NoItemsFound) {
			code = errors.ItemRequestedNotFound
		}

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
