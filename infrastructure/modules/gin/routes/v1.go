package services

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/middlewares"
)

func V1Routes(r *gin.Engine) {
	v1 := r.Group("/v1")

	anyone := v1.Group("")
	anyone.Use(middlewares.AuthHandler("any"))
	{
		anyone.GET("/auth", app.Auth)
	}
}
