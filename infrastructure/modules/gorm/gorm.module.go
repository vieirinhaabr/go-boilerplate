package gorm

import (
	"github.com/jinzhu/gorm"
	userRepo "go-boilerplate/infrastructure/modules/gorm/repos/user"
)

type gormModule struct {
	db *gorm.DB
}

func NewGormModule() *gormModule {
	return &gormModule{}
}

func (handler *gormModule) Configure() {
	db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	if err != nil {
		panic("Can't connect to DB")
	}

	db.LogMode(true)

	*handler = gormModule{db}
}

func (handler *gormModule) Start() {
	userRepo.Setup(handler.db)
}

func (handler *gormModule) Finish() {
	handler.db.Close()
}
