package types

import "go-boilerplate/infrastructure/global/errors"

type UseCase[req any, res any] func(params *req) *UseCaseResponse[res]

type UseCaseResponse[res interface{}] struct {
	Response  *res
	ErrorCode *errors.ErrorCode
	ErrorMsg  *string
}
