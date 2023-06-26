package api

import "github.com/gin-gonic/gin"

// @Summary 获取用户信息
// @Schemes
// @Description 获取用户信息
// @Tags config
// @Accept json
// @Produce json
// @Router /getUserInfo [get]
func userRouter(app *gin.RouterGroup) {
	app.GET("/getUserInfo", GetWechatUserInfo)
}
func GetWechatUserInfo(c *gin.Context) {
	//	获取code
	//	调用微信接口换取unionID、openID
	//	获取用户信息
	//	存储用户信息
}
