package services

import "github.com/gin-gonic/gin"

func UtilRoutes(r *gin.Engine) {
	r.GET("/live", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
