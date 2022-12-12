package userApi

type GetUserUri struct {
	ID string `form:"ID" uri:"id" binding:"required"`
}

type GetUserParams struct {
	GetUserUri
}
