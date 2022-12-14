package main

import (
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/global/types"
	ginModule "go-boilerplate/infrastructure/modules/gin"
	gormModule "go-boilerplate/infrastructure/modules/gorm"
)

func main() {
	environment.SetupVars()

	var gorm types.Module = gormModule.NewGormModule()
	var gin types.Module = ginModule.NewGinModule()

	gorm.Configure()
	gin.Configure()

	gorm.Start()
	gin.Start()

	defer gorm.Finish()
	defer gin.Finish()
}
