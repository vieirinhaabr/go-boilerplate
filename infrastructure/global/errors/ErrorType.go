package errors

type ErrorCode string

const (
	Validation ErrorCode = "validation"
	Internal   ErrorCode = "internal"
)
