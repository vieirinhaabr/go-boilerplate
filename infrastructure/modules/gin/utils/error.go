package ginUtils

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/global/types"
)

func ResolveError[res any](c *gin.Context, result *types.UseCaseResponse[res]) {
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
