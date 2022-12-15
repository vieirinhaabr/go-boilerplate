package userApi

type GetUserParams struct {
	ID string `form:"ID" uri:"id" validate:"required"`
	Type string `form:"Type" uri:"type" validate:"required"`
}
