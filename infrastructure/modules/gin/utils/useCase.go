package ginUtils

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/global/types"
)

func CallUseCase[req any, res any](useCase types.UseCase[req, res], c *gin.Context) {
	var r = new(req)
	_ = c.ShouldBindUri(&r)
	_ = c.ShouldBindJSON(&r)
	_ = c.ShouldBindQuery(&r)

	result := useCase(r)
	if result.ErrorCode != nil {
		ResolveError(c, result)
		return
	}

	c.JSON(200, *result.Response)
}
