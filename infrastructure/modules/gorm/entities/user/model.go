package user

import "github.com/jinzhu/gorm"

var model *gorm.DB = nil

func (User) Setup(db *gorm.DB) {
	db.CreateTable(User{})
	model = db.Model(User{})
}

func (User) Create() (bool, error) {
	return false, nil
}

func (User) Update() (bool, error) {
	return false, nil
}

func (User) Delete() (bool, error) {
	return false, nil
}
