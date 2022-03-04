package main

import (
	"AWC-gateway/wechat"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", wechat.VerifyURL)

	wechat.RefreshAccessToken()

	// 定时刷新信任的微信服务IP
	refreshTrustedProxies(r)
	go func() {
		for {
			refreshTrustedProxies(r)
			time.Sleep(24 * time.Hour)
		}
	}()

	err := r.Run(":80")
	if err != nil {
		return
	}
}

// refreshTrustedProxies 定时刷新信赖的callback ip
func refreshTrustedProxies(r *gin.Engine) {
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
