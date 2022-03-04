package wechat

import (
	"encoding/json"
	"github.com/ddliu/go-httpclient"
	"io/ioutil"
)

type ResponseBody struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// AccessTokenBody getAccessToken http请求结果
type AccessTokenBody struct {
	ResponseBody
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// CallbackIPBody GetCallbackIP http请求结果
type CallbackIPBody struct {
	ResponseBody
	IPList []string `json:"ip_list"`
}

// unmarshalResponse 解析请求响应结果
func unmarshalResponse(response *httpclient.Response, v interface{}) error {
	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bys, v)
}
