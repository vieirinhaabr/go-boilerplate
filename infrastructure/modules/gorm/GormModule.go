package gormModule

import (
	"database/sql"
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/modules/gorm/repos"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Module struct {
	conn *gorm.DB
	db   *sql.DB
}

func NewGormModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {
	var config = environment.Vars.Gorm

	var dsn = "host=" + config.Host +
		" port=" + config.Port +
		" dbname=" + config.DBName +
		" user=" + config.User +
		" password=" + config.Pass
	dialect := postgres.Open(dsn)

	conn, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, err := conn.DB()
	if err != nil {
		panic(err)
	}

	*handler = Module{conn, db}
}

func (handler *Module) Start() {
	gormRepo.Transaction.Setup(handler.conn)

	gormUserRepo.Repo.Setup(handler.conn)
}

func (handler *Module) Finish() {
	_ = handler.db.Close()
}
