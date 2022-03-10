package wechat

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	beFollowedText = `hello world !
感谢关注这个公众号。
此公众号还在施工中，只开放了狼人杀法官插件。
试着输入 help 查看狼人杀插件提供的功能。`
)

// HandleMessage 消息处理函数
type HandleMessage func(request EventRequest) string

var Handlers []HandleMessage

// HandleEvent 处理微信事件
func HandleEvent(c *gin.Context) {
	var req EventRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("HandleEvent err : ", err)
		return
	}

	var rspMsg string
	switch req.MsgType {
	case "event":
		rspMsg = beFollowedText
	case "text":
		rspMsg = Handlers[0](req)
	}
	if rspMsg == "" {
		c.String(http.StatusOK, "")
		return
	}

	rsp := EventResponse{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      "text",
		Content:      rspMsg,
	}
	c.XML(http.StatusOK, rsp)
}
