package main

import (
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/global/types"
	"go-boilerplate/infrastructure/modules/gin"
	"go-boilerplate/infrastructure/modules/gorm"
	"go-boilerplate/infrastructure/modules/grpc"
	"go-boilerplate/infrastructure/modules/mongo"
	"go-boilerplate/infrastructure/modules/redis"
	"golang.org/x/sync/errgroup"
)

var (
	toStart  errgroup.Group
	toFinish []func()
)

func addModule(module types.Module) types.Module {
	module.Configure()

	toStart.Go(func() error {
		return module.Start()
	})

	toFinish = append(toFinish, func() {
		module.Finish()
	})

	return module
}

func main() {
	environment.SetupVars()

	// Configure
	addModule(redisModule.NewRedisModule())
	addModule(mongoModule.NewMongoModule())
	addModule(gormModule.NewGormModule())
	addModule(grpcModule.NewGrpcModule())
	addModule(ginModule.NewGinModule())

	// Start sync
	if err := toStart.Wait(); err != nil {
		panic(err)
	}

	// Finish
	defer func() {
		for _, finish := range toFinish {
			finish()
		}
	}()
}
