package userApi

type UpdateUserUri struct {
	ID string `form:"ID" uri:"id" binding:"required"`
}

type UpdateUserBody struct {
	CreateUserBody
}

type UpdateUserParams struct {
	UpdateUserUri
	UpdateUserBody
}
