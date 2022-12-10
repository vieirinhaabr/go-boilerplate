package userRepo

import (
	"github.com/jinzhu/gorm"
	"go-boilerplate/domain/user/entities"
)

var model *gorm.DB = nil

func Setup(db *gorm.DB) {
	db.CreateTable(Schema{})
	model = db.Model(Schema{})
}

func (Schema) Create(user user.Entity) error {
	result := model.Create(&user)
	return result.Error
}

func (Schema) Update(user user.Entity) error {
	result := model.Update(&user)
	return result.Error
}
