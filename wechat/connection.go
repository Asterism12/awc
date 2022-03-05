package wechat

import (
	"github.com/ddliu/go-httpclient"
	"log"
	"time"
)

// GetCallbackIP 获取微信服务回调IP，用于请求准入策略
func GetCallbackIP() ([]string, error) {
	rsp, err := httpclient.Get("https://api.weixin.qq.com/cgi-bin/get_api_domain_ip", map[string]string{
		"access_token": accessToken,
	})
	if err != nil {
		return nil, err
	}

	var body CallbackIPResponse
	if err = unmarshalResponse(rsp, &body); err != nil {
		return nil, err
	}
	return body.IPList, nil
}

// getAccessToken 获取接口调用凭据
func getAccessToken() (AccessTokenResponse, error) {
	rsp, err := httpclient.Get("https://api.weixin.qq.com/cgi-bin/token", map[string]string{
		"grant_type": "client_credential",
		"appid":      config.APPID,
		"secret":     config.APPSecret,
	})
	if err != nil {
		return AccessTokenResponse{}, err
	}

	var body AccessTokenResponse
	if err = unmarshalResponse(rsp, &body); err != nil {
		return AccessTokenResponse{}, err
	}
	return body, nil
}

// SetAccessToken 设置AccessToken
func SetAccessToken() time.Duration {
	body, err := getAccessToken()
	if err != nil {
		log.Println("SetAccessToken err : ", err)
	}

	accessToken = body.AccessToken
	return time.Duration(body.ExpiresIn) * time.Second
}
