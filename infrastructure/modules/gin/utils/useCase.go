package ginUtils

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/api"
	"go-boilerplate/infrastructure/global/errors"
)

type UseCase[req any, res any] func(params *req) *api.UseCaseResponse[res]

func CallUseCase[req any, res any](useCase UseCase[req, res], c *gin.Context) {
	var r = new(req)
	_ = c.ShouldBindUri(&r)
	_ = c.ShouldBindJSON(&r)
	_ = c.ShouldBindQuery(&r)

	result := useCase(r)
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

		case errors.ItemRequestedNotFound:
			c.JSON(404, gin.H{"error": *result.ErrorMsg})
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
