package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"io/ioutil"
)

// ResponseBody http请求结果通用结构
type ResponseBody struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// AccessTokenBody getAccessToken http请求结果
type AccessTokenBody struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// CallbackIPBody GetCallbackIP http请求结果
type CallbackIPBody struct {
	IPList []string `json:"ip_list"`
}

// unmarshalResponse 解析请求响应结果
func unmarshalResponse(response *httpclient.Response, v interface{}) error {
	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var resBody ResponseBody
	err = json.Unmarshal(bys, &resBody)
	if err != nil {
		return err
	}
	if resBody.ErrCode != 0 {
		return fmt.Errorf("rsp code : %d, msg : %s", resBody.ErrCode, resBody.ErrMsg)
	}

	return json.Unmarshal(bys, v)
}
