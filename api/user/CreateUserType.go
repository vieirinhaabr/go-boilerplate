package userApi

import (
	userEntity "go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/global/errors"
)

type CreateUserParams struct {
	Email string `form:"Email" json:"email" binding:"required"`
	Name  string `form:"Name" json:"name" binding:"required"`
}

type CreateUserResponse struct {
	Response  *userEntity.User
	ErrorCode *errors.ErrorCode
	ErrorMsg  *string
}
