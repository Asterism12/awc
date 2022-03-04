package wechat

import (
	"github.com/ddliu/go-httpclient"
	"log"
)

// GetCallbackIP 获取微信服务回调IP，用于请求准入策略
func GetCallbackIP() ([]string, error) {
	rsp, err := httpclient.Get("https://api.weixin.qq.com/cgi-bin/get_api_domain_ip", map[string]string{
		"access_token": accessToken,
	})
	if err != nil {
		return nil, err
	}

	var body CallbackIPBody
	if err = unmarshalResponse(rsp, &body); err != nil {
		return nil, err
	}
	return body.IPList, nil
}

// getAccessToken 获取接口调用凭据
func getAccessToken() (AccessTokenBody, error) {
	rsp, err := httpclient.Get("https://api.weixin.qq.com/cgi-bin/token", map[string]string{
		"grant_type": "client_credential",
		"appid":      config.APPID,
		"secret":     config.APPSecret,
	})
	if err != nil {
		return AccessTokenBody{}, err
	}

	var body AccessTokenBody
	if err = unmarshalResponse(rsp, &body); err != nil {
		return AccessTokenBody{}, err
	}
	return body, nil
}

// RefreshAccessToken 刷新AccessToken
func RefreshAccessToken() {
	body, err := getAccessToken()
	if err != nil {
		log.Println("RefreshAccessToken err : ", err)
	}
	accessToken = body.AccessToken
}
