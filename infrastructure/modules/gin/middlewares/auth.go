package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/util"
)

func AuthHandler(authRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		maps, msg := util.ValidateToken(authRoles, token)

		if maps == nil {
			c.JSON(403, gin.H{"message": msg})
			c.Abort()
			return
		}

		c.Set("userId", maps.UserId)
		c.Set("username", maps.Username)
		c.Next()
	}
}
