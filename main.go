package main

import (
	"AWC-gateway/crontab"
	"AWC-gateway/wechat"
	_ "AWC-gateway/wolfkill"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", wechat.VerifyURL)
	r.POST("/", wechat.ReceiveMessage)

	// 初始化AccessToken及微信服务IP
	wechat.SetAccessToken()
	setTrustedProxies(r)

	// 定时刷新信任的微信服务IP
	crontab.StartCronTab(func() time.Duration {
		setTrustedProxies(r)
		return 24 * time.Hour
	})

	// 定时刷新AccessToken
	crontab.StartCronTab(func() time.Duration {
		return wechat.SetAccessToken()
	})

	err := r.Run(":80")
	if err != nil {
		return
	}
}

// setTrustedProxies 设置可信任的ip集
func setTrustedProxies(r *gin.Engine) {
	ips, err := wechat.GetCallbackIP()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("trusted ips : ", ips)

	err = r.SetTrustedProxies(ips)
	if err != nil {
		log.Println(err)
		return
	}
}
