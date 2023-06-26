package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {
	configRouter(app.Group("/config"))
	userRouter(app.Group("/user"))
}
