package userGinActions

import (
	"github.com/gin-gonic/gin"
	userUseCase "go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/gin/utils"
)

func CreateUserSQLAction(c *gin.Context) {
	ginUtils.CallUseCase(userUseCase.CreateUserSQLUseCase, c)
}
