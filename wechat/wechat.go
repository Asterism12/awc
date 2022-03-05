// Package wechat 处理wechat服务发来的请求
package wechat

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"strings"
)

// VerifyURL 微信服务验证URL
func VerifyURL(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Println("VerifyURL err : ", err)
		return
	}

	pwd := genPassword([]string{config.Token, req.Timestamp, req.Nonce})

	if pwd == req.Signature {
		c.String(http.StatusOK, req.EchoStr)
	}

	log.Println("VerifyURL result : ", pwd == req.Signature)
}

// genPassword 根据微信提供的逻辑生成密码
//
// refer : https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html
func genPassword(params []string) string {
	sort.Strings(params)
	key := strings.Join(params, "")
	pwd := sha1.Sum([]byte(key))
	return fmt.Sprintf("%x", pwd[:20])
}
