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

	go refreshTrustedProxies(r)
	wechat.RefreshAccessToken()

	err := r.Run(":80")
	if err != nil {
		return
	}
}

// refreshTrustedProxies 定时刷新信赖的callback ip
func refreshTrustedProxies(r *gin.Engine) {
	ips, err := wechat.GetCallbackIP()
	if err != nil {
		log.Panicln(err)
	}

	err = r.SetTrustedProxies(ips)
	if err != nil {
		log.Panicln(err)
	}

	time.Sleep(24 * time.Hour)
}
