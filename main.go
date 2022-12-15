package main

import (
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/global/types"
	"go-boilerplate/infrastructure/modules/gin"
	"go-boilerplate/infrastructure/modules/gorm"
	mongoModule "go-boilerplate/infrastructure/modules/mongo"
)

func main() {
	environment.SetupVars()

	var mongo types.Module = mongoModule.NewMongoModule()
	var gorm types.Module = gormModule.NewGormModule()
	var gin types.Module = ginModule.NewGinModule()

	mongo.Configure()
	gorm.Configure()
	gin.Configure()

	mongo.Start()
	gorm.Start()
	gin.Start()

	defer mongo.Finish()
	defer gorm.Finish()
	defer gin.Finish()
}
