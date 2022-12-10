package gin

import (
	"github.com/jinzhu/gorm"
	"go-boilerplate/infrastructure/modules/gorm/entities/user"
)

type gormModule struct {
	db *gorm.DB
}

func NewGormModule() *gormModule {
	return &gormModule{}
}

func (handler *gormModule) configure(res chan any) {
	db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	if err != nil {
		panic("Can't connect to DB")
	}

	db.LogMode(true)

	*handler = gormModule{db}
	res <- handler
}

func (handler *gormModule) start(res chan any) {
	user.User{}.Setup(handler.db)

	res <- handler
}

func (handler *gormModule) finish(res chan any) {
	handler.db.Close()
	res <- handler
}
