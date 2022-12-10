package v1

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/middlewares"
)

func test(c *gin.Context) {}

func userRoutes(r *gin.RouterGroup) {
	group := r.Group("/user")

	group.Use(middlewares.AuthHandler("any"))
	{
		group.POST("", test)
	}
}
