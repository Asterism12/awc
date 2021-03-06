package wechat

// VerifyRequest 微信服务验证URL请求体
type VerifyRequest struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
}

// EventRequest 微信事件请求体
type EventRequest struct {
	ToUserName   string `form:"ToUserName" xml:"ToUserName"`
	FromUserName string `form:"FromUserName" xml:"FromUserName"`
	CreateTime   int    `form:"CreateTime" xml:"CreateTime"`
	MsgType      string `form:"MsgType" xml:"MsgType"`
	Content      string `form:"Content" xml:"Content"`
	MsgID        int64  `form:"MsgId" xml:"MsgId"`
}
