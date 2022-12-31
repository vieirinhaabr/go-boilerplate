package userApi

type GetUserByNameParams struct {
	Name string `form:"Name" uri:"name" validate:"required"`
}
