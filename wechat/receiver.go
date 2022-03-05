package wechat

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// ReceiveMessage 接收普通消息
func ReceiveMessage(c *gin.Context) {
	var req ReceiveMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("ReceiveMessage err : ", err)
		return
	}

	rsp := ReceiveMessageResponse{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      "text",
		Content:      ";" + req.Content,
	}
	c.XML(http.StatusOK, rsp)
}
