package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserInfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
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

const (
	redirectAuthURL       = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	accessTokenURL        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
)

// 授权成功返回url
func (weChat *Config) GetAuthUrl(redirectUrl string) string {
	oauth2url := fmt.Sprintf(redirectAuthURL, weChat.AppId, redirectUrl, "snsapi_userinfo ")
	return oauth2url
}

// 通过code获取接口调用凭证
func (weChat *Config) GetAccessToken(code string) (*SnsOauth2, error) {
	url := fmt.Sprintf(accessTokenURL, weChat.AppId, weChat.AppSecret, code)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("获取 AccessToken 错误", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取 AccessToken 读取返回body错误", err)
		return nil, err
	}
	if bytes.Contains(body, []byte("errcode")) {
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			fmt.Printf("获取 AccessToken 的错误信息 %+v\n", ater)
			return nil, err
		}
		return nil, fmt.Errorf("%s", ater.ErrMsg)
	} else {
		atr := SnsOauth2{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			fmt.Println("获取 AccessToken 返回数据json解析错误", err)
			return nil, err
		}
		return &atr, nil
	}

}
func (weChat *Config) GetUserInfo(accessToken, openId string) (*UserInfo, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken, openId)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("用户信息get请求失败")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取 用户信息 读取返回body错误", err)
		return nil, err
	}
	if bytes.Contains(body, []byte("errcode")) {
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			fmt.Printf("获取 用户信息 的错误信息 %+v\n", ater)
			return nil, err
		}
		return nil, fmt.Errorf("%s", ater.ErrMsg)
	} else {
		userInfo := UserInfo{}
		err = json.Unmarshal(body, &userInfo)
		if err != nil {
			fmt.Println("获取 用户信息 返回数据json解析错误", err)
			return nil, err
		}
		return &userInfo, nil
	}
}
func (weChat *Config) RefreshToken(refreshToken string) (*SnsOauth2, error) {
	url := fmt.Sprintf(refreshAccessTokenURL, weChat.AppId, refreshToken)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("重新获取 AccessToken get请求失败")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("重新获取 AccessToken 读取返回body错误", err)
		return nil, err
	}
	if bytes.Contains(body, []byte("errcode")) {
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			fmt.Printf("重新获取 AccessToken 的错误信息 %+v\n", ater)
			return nil, err
		}
		return nil, fmt.Errorf("%s", ater.ErrMsg)
	} else {
		so2 := SnsOauth2{}
		err = json.Unmarshal(body, &so2)
		if err != nil {
			fmt.Println("重新获取 AccessToken 返回数据json解析错误", err)
			return nil, err
		}
		return &so2, nil
	}
}
