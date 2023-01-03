package v1

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/actions/user"
)

func userRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")

	sqlGroup := userGroup.Group("/sql")

	//group.Use(middlewares.AuthHandler("any"))
	{
		sqlGroup.POST("", userGinActions.CreateUserSQLAction)
		sqlGroup.GET("/:id", userGinActions.GetUserSQLAction)
		sqlGroup.PUT("/:id", userGinActions.UpdateUserSQLAction)
	}

	noSqlGroup := userGroup.Group("/no-sql")

	//group.Use(middlewares.AuthHandler("any"))
	{
		noSqlGroup.POST("", userGinActions.CreateUserNoSQLAction)
		noSqlGroup.GET("/:id", userGinActions.GetUserNoSQLAction)
		noSqlGroup.PUT("/:id", userGinActions.UpdateUserNoSQLAction)
	}

	redisGroup := userGroup.Group("/redis")

	//group.Use(middlewares.AuthHandler("any"))
	{
		redisGroup.POST("", userGinActions.CreateUserRedisAction)
		redisGroup.GET("/:name", userGinActions.GetUserRedisAction)
	}

	grpcGroup := userGroup.Group("/grpc")

	//group.Use(middlewares.AuthHandler("any"))
	{
		grpcGroup.POST("", userGinActions.CreateUserGrpcAction)
		grpcGroup.GET("/:id", userGinActions.GetUserGrpcAction)
		grpcGroup.PUT("/:id", userGinActions.UpdateUserGrpcAction)
	}
}
