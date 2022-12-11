package actions

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/api/user"
	"go-boilerplate/app/user"
	"go-boilerplate/infrastructure/global/errors"
)

func CreateUserAction(c *gin.Context) {
	var req userApi.CreateUserParams
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Status(412)
		c.JSON(412, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	result := userUseCase.CreateUserAction(&req)
	if result.ErrorCode != nil {
		switch *result.ErrorCode {
		case errors.Internal:
			c.JSON(500, gin.H{"error": *result.ErrorMsg})
			c.Abort()
			return

		case errors.Validation:
			c.JSON(412, gin.H{"error": *result.ErrorMsg})
			c.Abort()
			return

		default:
			c.JSON(500, gin.H{"error": "Cannot handle the error"})
			c.Abort()
			return
		}
	}

	c.JSON(200, *result.Response)
}
