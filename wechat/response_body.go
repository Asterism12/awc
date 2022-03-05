package wechat

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"io/ioutil"
)

// Response http请求结果通用结构
type Response struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// AccessTokenResponse getAccessToken请求体
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// CallbackIPResponse GetCallbackIP请求体
type CallbackIPResponse struct {
	IPList []string `json:"ip_list"`
}

// ReceiveMessageResponse 接收普通消息返回体
type ReceiveMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `form:"ToUserName" xml:"ToUserName"`
	FromUserName string   `form:"FromUserName" xml:"FromUserName"`
	CreateTime   int      `form:"CreateTime" xml:"CreateTime"`
	MsgType      string   `form:"MsgType" xml:"MsgType"`
	Content      string   `form:"Content" xml:"Content"`
}

// unmarshalResponse 解析请求响应结果
func unmarshalResponse(response *httpclient.Response, v interface{}) error {
	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var rspBody Response
	err = json.Unmarshal(bys, &rspBody)
	if err != nil {
		return err
	}
	if rspBody.ErrCode != 0 {
		return fmt.Errorf("rsp code : %d, msg : %s", rspBody.ErrCode, rspBody.ErrMsg)
	}

	return json.Unmarshal(bys, v)
}
