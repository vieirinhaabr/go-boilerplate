package v1

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/actions/user"
)

func userRoutes(r *gin.RouterGroup) {
	sqlGroup := r.Group("/user/sql")

	//group.Use(middlewares.AuthHandler("any"))
	{
		sqlGroup.POST("", userGinActions.CreateUserSQLAction)
		sqlGroup.GET("/:id", userGinActions.GetUserSQLAction)
		sqlGroup.PUT("/:id", userGinActions.UpdateUserSQLAction)
	}

	noSqlGroup := r.Group("/user/no-sql")

	//group.Use(middlewares.AuthHandler("any"))
	{
		noSqlGroup.POST("", userGinActions.CreateUserNoSQLAction)
		noSqlGroup.GET("/:id", userGinActions.GetUserNoSQLAction)
		noSqlGroup.PUT("/:id", userGinActions.UpdateUserNoSQLAction)
	}
}
