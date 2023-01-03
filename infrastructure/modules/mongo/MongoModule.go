package mongoModule

import (
	migrate "github.com/xakep666/mongo-migrate"
	"go-boilerplate/config/environment"
	_ "go-boilerplate/infrastructure/modules/mongo/migrations"
	"go-boilerplate/infrastructure/modules/mongo/repos/user"
	"go-boilerplate/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MigrationCollectionName = "migrations"

type Module struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {
	utils.Log.Warn("[MONGO] ᛃ Configuring\n")

	var config = environment.Vars.Mongo

	ctx, cancel := utils.CreateContext()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))
	if err != nil {
		panic(err)
	}

	database := client.Database(config.DBName)

	migrate.SetDatabase(database)
	migrate.SetMigrationsCollection(MigrationCollectionName)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		panic("Cannot run Mongo migrations: " + err.Error())
	}

	*handler = Module{client, database}
}

func (handler *Module) Start() error {
	utils.Log.Warn("[MONGO] ▶ Starting\n")

	mongoUserRepo.Repo.Setup(handler.database)

	return nil
}

func (handler *Module) Finish() {
	utils.Log.Warn("[MONGO] ■ Finishing\n")

	ctx, cancel := utils.CreateContext()
	defer cancel()

	_ = handler.client.Disconnect(ctx)
}
