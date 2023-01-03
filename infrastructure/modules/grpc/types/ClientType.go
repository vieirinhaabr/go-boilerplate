package grpcTypes

import (
	"go-boilerplate/infrastructure/global/errors"
)

type ClientResponse[res interface{}] struct {
	Response *res
	Error    *ClientError
}

type ClientError struct {
	Code errors.ErrorCode
	Msg  string
}
