package services

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/middlewares"
)

func SetupRouter(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}

func InitRouter(engine *gin.Engine) {
	SetupRouter(engine)

	V1Routes(engine)
	UtilRoutes(engine)
}
