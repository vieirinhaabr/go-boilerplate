package v1

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/actions/user"
)

func userRoutes(r *gin.RouterGroup) {
	group := r.Group("/user")

	//group.Use(middlewares.AuthHandler("any"))
	{
		group.POST("", userGinActions.CreateUserAction)
		group.GET("/:id/:type", userGinActions.GetUserAction)
		group.PUT("/:id", userGinActions.UpdateUserAction)
	}
}
