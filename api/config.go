package api

import "github.com/gin-gonic/gin"

func configRouter(app *gin.RouterGroup) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "config",
		})
	})
}
