package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {
	//miniprogramRouter(app.Group("/miniprogram"))
	officialaccountRouter(app.Group("/officialaccount"))
	//payRouter(app.Group("/pay"))
}
