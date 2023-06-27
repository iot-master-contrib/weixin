package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {
	userRouter(app.Group("/user"))
}
