package userRepo

import (
	"go-boilerplate/domain/user/entities"
	"gorm.io/gorm"
)

type repo struct {
	model *gorm.DB
}

var Repo repo

func (r repo) Setup(conn *gorm.DB) {
	conn.AutoMigrate(&Schema{})
	Repo = repo{
		model: conn.Model(&Schema{}),
	}
}

func (r repo) Create(user *userEntity.User) error {
	result := r.model.Create(user)
	return result.Error
}

func (r repo) Update(user *userEntity.User) error {
	result := r.model.Save(user)
	return result.Error
}

func (r repo) GetById(id string) (userEntity.User, error) {
	var user *userEntity.User
	result := r.model.First(user, id)

	return *user, result.Error
}
