/**
 * Created by VoidArtanis on 10/22/2017
 */

package routes

import (
	"github.com/VoidArtanis/go-rest-boilerplate/controllers"
	"github.com/VoidArtanis/go-rest-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")

	authGroup.Use(middlewares.AuthHandler("admin"))
	{
		authGroup.GET("/getmessage", controllers.GetSecretText)
	}
}
