package userApi

type GetUserParams struct {
	ID string `form:"ID" uri:"id" validate:"required"`
}
