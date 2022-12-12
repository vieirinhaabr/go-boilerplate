package v1

import (
	"github.com/gin-gonic/gin"
	userApi "go-boilerplate/api/user"
	userUseCase "go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/gin/middlewares"
)

func userRoutes(r *gin.RouterGroup) {
	group := r.Group("/user")

	//group.Use(middlewares.AuthHandler("any"))
	{
		group.POST("", middlewares.CallUseCase[any, userApi.CreateUserBody](userUseCase.CreateUserUseCase))
		group.GET("/:id", middlewares.CallUseCase[userApi.GetUserUri, any](userUseCase.GetUserUseCase))
		group.PUT("/:id", middlewares.CallUseCase[userApi.UpdateUserUri, userApi.UpdateUserBody](userUseCase.UpdateUserUseCase))
	}
}
