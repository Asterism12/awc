package wechat

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// HandleMessage 消息处理函数
type HandleMessage func(request ReceiveMessageRequest) string

var Handlers []HandleMessage

// ReceiveMessage 接收普通消息
func ReceiveMessage(c *gin.Context) {
	var req ReceiveMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("ReceiveMessage err : ", err)
		return
	}

	msg := Handlers[0](req)
	if msg == "" {
		c.String(http.StatusOK, "")
		return
	}

	rsp := ReceiveMessageResponse{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      "text",
		Content:      msg,
	}
	c.XML(http.StatusOK, rsp)
}
