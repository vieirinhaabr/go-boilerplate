package userGinActions

import (
	"github.com/gin-gonic/gin"
	userUseCase "go-boilerplate/app/user"
	"go-boilerplate/infrastructure/modules/gin/utils"
)

func UpdateUserGrpcAction(c *gin.Context) {
	ginUtils.CallUseCase(userUseCase.UpdateUserGrpcUseCase, c)
}
