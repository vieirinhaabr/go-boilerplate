package userApi

type CreateUserBody struct {
	Email string `form:"Email" json:"email" binding:"required"`
	Name  string `form:"Name" json:"name" binding:"required"`
}

type CreateUserParams struct {
	CreateUserBody
}
