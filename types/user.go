package types

import (
	"fmt"
	"net/http"
)

type userInfo struct{}

type WeChat struct {
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
}

type SnsOauth2 struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	Openid         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
	Unionid        string ` json:"unionid"`
}

type AccessTokenErrorResponse struct {
	ErrMsg  string `json:"err_msg"`
	ErrCode string `json:"err_code"`
}

// 授权成功返回url
func (weChat *WeChat) GetAuthUrl(redirectUrl string) string {
	oauth2url := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect", weChat.AppId, redirectUrl)
	return oauth2url
}

// 通过code获取接口调用凭证
func (weChat *WeChat) GetAccessToken(code string) (*SnsOauth2, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", weChat.AppId, weChat.AppSecret, code)
	accessToken, err := http.Get(url)
	if err != nil {
		fmt.Println("获取AccessToken错误")
		return nil, err
	}
	return &accessToken, nil
}
