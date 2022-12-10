package gorm

import (
	"database/sql"
	"go-boilerplate/infrastructure/modules/gorm/repos"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormModule struct {
	conn *gorm.DB
	db   *sql.DB
}

func NewGormModule() *gormModule {
	return &gormModule{}
}

func (handler *gormModule) Configure() {
	dns := postgres.Open("host=localhost port=5432 dbname=boilerplate user=dbuser password=dbpass")
	conn, err := gorm.Open(dns, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, err := conn.DB()
	if err != nil {
		panic(err)
	}

	*handler = gormModule{conn, db}
}

func (handler *gormModule) Start() {
	transaction.Setup(handler.conn)

	userRepo.Setup(handler.conn)
}

func (handler *gormModule) Finish() {
	handler.db.Close()
}
