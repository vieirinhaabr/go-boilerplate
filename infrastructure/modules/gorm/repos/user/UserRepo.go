package gormUserRepo

import (
	"go-boilerplate/domain/user/entities"
	"gorm.io/gorm"
)

type repo struct {
	model *gorm.DB
}

var Repo repo

func (r repo) Setup(conn *gorm.DB) {
	Repo = repo{
		model: conn,
	}
}

func (r repo) Create(user *userEntity.User) error {
	result := r.model.Model(&Schema{}).Create(user)
	return result.Error
}

func (r repo) Update(user *userEntity.User) error {
	result := r.model.Model(&Schema{}).Where("id = ?", user.ID).Updates(user)
	return result.Error
}

func (r repo) GetById(id string) (userEntity.User, error) {
	var user = new(userEntity.User)
	result := r.model.Model(&Schema{}).Where("id = ?", id).First(user)

	return *user, result.Error
}
