package userValidator

import (
	"github.com/go-playground/validator/v10"
	userApi "go-boilerplate/api/user"
)

func CreateUserValidator(params *userApi.CreateUserParams) error {
	var validate = validator.New()
	err := validate.Struct(params)

	return err
}
