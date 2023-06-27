package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"weixin/types"
)

// 后期填充
var weChat = &types.WeChat{
	AppId:       "",
	AppSecret:   "",
	AccessToken: "",
}
//var userInfo *types.UserInfo

// GetAuthUrl @Summary 获取授权引导页面
// @Schemes
// @Description 获取授权引导页面
// @Tags user
// @Accept json
// @Produce json
// @Router /getAuthUrl [get]
func GetAuthUrl(c *gin.Context) {
	redirectUrl := c.Param("url")
	oauth2Url := weChat.GetAuthUrl(redirectUrl)
	curd.OK(c, oauth2Url)
}

// GetAccessToken @Summary 获取AccessToken
// @Schemes
// @Description 获取AccessToken
// @Tags user
// @Accept json
// @Produce json
// @Router /getAccessToken [get]
func GetAccessToken(c *gin.Context) {
	code := c.Param("code")
	at, err := weChat.GetAccessToken(code)
	if err != nil {
		curd.Error(c, err)
		return
	}
	curd.OK(c, at)
}

// GetUserInfo @Summary 获取用户信息
// @Schemes
// @Description 获取用户信息
// @Tags user
// @Accept json
// @Produce json
// @Router /getUserInfo [get]
func GetUserInfo(c *gin.Context) {
	snsOauth2 := types.SnsOauth2{}
	err := c.ShouldBind(&snsOauth2)
	if err != nil {
		curd.Error(c, err)
	}
	info, err := weChat.GetUserInfo(snsOauth2.AccessToken, snsOauth2.Openid)
	if err != nil {
		curd.Error(c, err)
	}
	curd.OK(c, info)
}
func userRouter(app *gin.RouterGroup) {
	app.GET("/getAuthUrl", GetAuthUrl)
	app.GET("/getAccessToken", GetAccessToken)
	app.POST("/getUserInfo", GetUserInfo)
}
