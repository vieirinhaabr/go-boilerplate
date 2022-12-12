package api

import "go-boilerplate/infrastructure/global/errors"

type UseCaseResponse[res interface{}] struct {
	Response  *res
	ErrorCode *errors.ErrorCode
	ErrorMsg  *string
}
