package userApi

type UpdateUserParams struct {
	ID string `form:"ID" uri:"id" validate:"required"`
	Email string `form:"Email" json:"email" validate:"required,email"`
	Name  string `form:"Name" json:"name" validate:"required,gte=4,lte=40"`
	Type  string `form:"Type" json:"type" validate:"required"`
}
