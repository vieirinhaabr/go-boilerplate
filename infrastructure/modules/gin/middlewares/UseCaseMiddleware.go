package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-boilerplate/api"
	"go-boilerplate/infrastructure/global/errors"
	"reflect"
)

type UseCase[req any, res any] func(params *req) *api.UseCaseResponse[res]

func CallUseCase[uri any, body any, req any, res any](useCase UseCase[req, res]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r = new(req)
		var u = new(uri)
		var b = new(body)

		if reflect.TypeOf(u).String() != reflect.TypeOf(new(any)).String() {
			err := c.ShouldBindUri(&u)
			if err != nil {
				c.Status(412)
				c.JSON(412, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			jq, err := json.Marshal(u)
			if err != nil {
				c.Status(422)
				c.JSON(422, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			err = json.Unmarshal(jq, &r)
			if err != nil {
				c.Status(422)
				c.JSON(422, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		if reflect.TypeOf(b).String() != reflect.TypeOf(new(any)).String() {
			err := c.ShouldBindJSON(&b)
			if err != nil {
				c.Status(412)
				c.JSON(412, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			jb, err := json.Marshal(b)
			if err != nil {
				c.Status(422)
				c.JSON(422, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			err = json.Unmarshal(jb, &r)
			if err != nil {
				c.Status(422)
				c.JSON(422, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

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

			default:
				c.JSON(500, gin.H{"error": "Cannot handle the error"})
				c.Abort()
				return
			}
		}

		c.JSON(200, *result.Response)
	}
}
