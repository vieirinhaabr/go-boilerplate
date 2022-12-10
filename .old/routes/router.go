/**
 * Created by VoidArtanis on 10/22/2017
 */

package routes

import (
	"github.com/VoidArtanis/go-rest-boilerplate/controllers"
	"github.com/VoidArtanis/go-rest-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	InitMiddleware(engine)
	authController := new(controllers.AuthController)
	engine.POST("/login", authController.HandleLogin)
	RegisterProtectedRoutes(engine)
	RegisterPublicRoutes(engine)
	RegisterUtilityRoutes(engine)
}

func InitMiddleware(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}
