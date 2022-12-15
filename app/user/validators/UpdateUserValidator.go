package userValidator

import (
	"github.com/go-playground/validator/v10"
	userApi "go-boilerplate/api/user"
)

func UpdateUserValidator(params *userApi.UpdateUserParams) error {
	var validate = validator.New()
	err := validate.Struct(params)

	return err
}
