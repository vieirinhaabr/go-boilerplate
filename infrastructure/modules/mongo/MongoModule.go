package mongoModule

import (
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/modules/mongo/repos/user"
	"go-boilerplate/infrastructure/modules/mongo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Module struct {
	client *mongo.Client
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

	*handler = Module{client, database}
}

func (handler *Module) Start() {
	mongoUserRepo.Repo.Setup(handler.database)
}

func (handler *Module) Finish() {
	ctx, _ := utils.CreateContext()

	_ = handler.client.Disconnect(ctx)
}
