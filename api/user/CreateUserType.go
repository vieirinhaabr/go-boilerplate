package userApi

type CreateUserParams struct {
	Email string `form:"Email" json:"email" validate:"required,email"`
	Name  string `form:"Name" json:"name" validate:"required,gte=4,lte=40"`
}
