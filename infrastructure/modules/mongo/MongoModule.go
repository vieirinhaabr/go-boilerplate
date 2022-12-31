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
	var config = environment.Vars.Mongo

	ctx, _ := utils.CreateContext()

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

func (handler *Module) Start() {
	mongoUserRepo.Repo.Setup(handler.database)
}

func (handler *Module) Finish() {
	ctx, _ := utils.CreateContext()

	_ = handler.client.Disconnect(ctx)
}
