package userRepo

import (
	"go-boilerplate/domain/user/entities"
	"gorm.io/gorm"
)

var model *gorm.DB = nil

func Setup(conn *gorm.DB) {
	conn.AutoMigrate(&Schema{})
	model = conn.Table(Schema{}.TableName(), Schema{})
}

func (Schema) Create(user user.Entity) error {
	result := model.Create(&user)
	return result.Error
}

func (Schema) Update(user user.Entity) error {
	result := model.Save(&user)
	return result.Error
}
