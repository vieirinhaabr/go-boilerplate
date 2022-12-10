package main

import (
	"go-boilerplate/infrastructure/global/types"
	"go-boilerplate/infrastructure/modules/gin"
	"go-boilerplate/infrastructure/modules/gorm"
)

func main() {
	var gorm types.Module = gorm.NewGormModule()
	var gin types.Module = gin.NewGinModule()

	gorm.Configure()
	gin.Configure()

	gorm.Start()
	gin.Start()

	defer gorm.Finish()
	defer gin.Finish()
}
