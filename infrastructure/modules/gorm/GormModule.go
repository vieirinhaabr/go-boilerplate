package gormModule

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	migratePg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/modules/gorm/repos"
	"go-boilerplate/infrastructure/modules/gorm/repos/user"
	"go-boilerplate/utils"
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

func runMigrations(config environment.Gorm, dsn string) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Cannot open sql instance: " + err.Error())
	}
	defer conn.Close()

	driver, err := migratePg.WithInstance(conn, &migratePg.Config{})
	if err != nil {
		panic("Cannot open migration instance: " + err.Error())
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+config.MigrationsDir,
		"postgres",
		driver,
	)
	if err != nil {
		panic("Cannot load migrations and driver: " + err.Error())
	}

	// CASE NEED TO ROLLBACK A VERSION
	// err = m.Force(1)
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		panic("Cannot run the SQL migrations: " + err.Error())
	}
}

func (handler *Module) Configure() {
	utils.Log.Warn("[GORM] ᛃ Configuring\n")

	var config = environment.Vars.Gorm

	var dsn = "host=" + config.Host +
		" port=" + config.Port +
		" dbname=" + config.DBName +
		" user=" + config.User +
		" password=" + config.Pass +
		" sslmode=disable"

	runMigrations(config, dsn)
	dialect := postgres.Open(dsn)

	conn, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic("Cannot open Gorm connection: " + err.Error())
	}

	db, err := conn.DB()
	if err != nil {
		panic("Cannot open Gorm database: " + err.Error())
	}

	*handler = Module{conn, db}
}

func (handler *Module) Start() error {
	utils.Log.Warn("[GORM] ▶ Starting\n")

	gormRepo.Transaction.Setup(handler.conn)
	gormUserRepo.Repo.Setup(handler.conn)

	return nil
}

func (handler *Module) Finish() {
	utils.Log.Warn("[GORM] ■ Finishing\n")

	_ = handler.db.Close()
}
