package healthcheck

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET("/live", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
