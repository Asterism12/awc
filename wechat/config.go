package wechat

import (
	"github.com/jinzhu/configor"
	"log"
)

var config = struct {
	APPID     string `json:"app_id"`
	APPSecret string `json:"app_secret"`
	Token     string `json:"token"`
}{}

var (
	accessToken string // 接口调用凭据
)

func init() {
	err := configor.Load(&config, "wechat.json")
	if err != nil {
		panic(err)
	}

	log.Println(config)
}
