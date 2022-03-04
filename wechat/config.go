package wechat

import (
	"github.com/jinzhu/configor"
)

var config = struct {
	APPID     string
	APPSecret string
	Token     string
}{}

var (
	accessToken string // 接口调用凭据
)

func init() {
	err := configor.Load(&config, "wechat.json")
	if err != nil {
		panic(err)
	}
}
