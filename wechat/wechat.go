// Package wechat 处理wechat服务发来的请求
package wechat

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sort"
	"strings"
)

// VerifyURL 微信服务验证URL
func VerifyURL(c *gin.Context) {
	token := config.Token
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")

	pwd := genPassword([]string{token, timestamp, nonce})

	if pwd == signature {
		echoStr := c.Query("echoStr")
		c.String(200, echoStr)
	}

	log.Println("VerifyURL result : ", pwd == signature)
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
