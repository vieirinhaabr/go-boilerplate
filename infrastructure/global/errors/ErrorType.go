package errors

type RepoErrorCode string
type ErrorCode string

const (
	ItemRequestedNotFound ErrorCode = "item-requested-not-found"
	Validation ErrorCode = "validation"
	Internal   ErrorCode = "internal"
)

const (
	NoItemsFound RepoErrorCode = "no items found"
)
