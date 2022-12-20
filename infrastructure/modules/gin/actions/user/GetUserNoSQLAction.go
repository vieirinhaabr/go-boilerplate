package userGinActions

import (
	"github.com/gin-gonic/gin"
	userUseCase "go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/gin/utils"
)

func GetUserNoSQLAction(c *gin.Context) {
	ginUtils.CallUseCase(userUseCase.GetUserNoSQLUseCase, c)
}
