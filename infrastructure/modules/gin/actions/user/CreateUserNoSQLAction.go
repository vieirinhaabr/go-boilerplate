package userGinActions

import (
	"github.com/gin-gonic/gin"
	userUseCase "go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/gin/utils"
)

func CreateUserNoSQLAction(c *gin.Context) {
	ginUtils.CallUseCase(userUseCase.CreateUserNoSQLUseCase, c)
}
