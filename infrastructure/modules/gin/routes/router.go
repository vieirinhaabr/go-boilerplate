package router

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/middlewares"
	"go-boilerplate/infrastructure/modules/gin/routes/healthcheck"
	"go-boilerplate/infrastructure/modules/gin/routes/v1"
)

func configure(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}

func SetRoutes(engine *gin.Engine) {
	configure(engine)

	v1.Routes(engine)
	healthcheck.Routes(engine)
}
