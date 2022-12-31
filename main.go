package main

import (
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/global/types"
	"go-boilerplate/infrastructure/modules/gin"
	"go-boilerplate/infrastructure/modules/gorm"
	mongoModule "go-boilerplate/infrastructure/modules/mongo"
	redisModule "go-boilerplate/infrastructure/modules/redis"
)

func main() {
	environment.SetupVars()

	var redis types.Module = redisModule.NewRedisModule()
	var mongo types.Module = mongoModule.NewMongoModule()
	var gorm types.Module = gormModule.NewGormModule()
	var gin types.Module = ginModule.NewGinModule()

	redis.Configure()
	mongo.Configure()
	gorm.Configure()
	gin.Configure()

	redis.Start()
	mongo.Start()
	gorm.Start()
	gin.Start()

	defer redis.Finish()
	defer mongo.Finish()
	defer gorm.Finish()
	defer gin.Finish()
}
