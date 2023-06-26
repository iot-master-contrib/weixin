package api

import (
	"github.com/gin-gonic/gin"
	"weixin/types"
)

var weChat *types.WeChat

// @Summary 获取授权引导页面
// @Schemes
// @Description 获取授权引导页面
// @Tags config
// @Accept json
// @Produce json
// @Router /getAuthUrl [get]
func userRouter(app *gin.RouterGroup) {
	app.GET("/getAuthUrl", GetAuthUrl)
}
func GetAuthUrl(c *gin.Context) {
	redirectUrl := c.Param("url")
	oauth2Url := weChat.GetAuthUrl(redirectUrl)
	c.JSON(200, gin.H{
		"oauth2Url": oauth2Url,
	})
}
