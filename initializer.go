package wxsdk

import (
	"log"
	"regexp"
)

// TConfig 配置
var (
	AppId          string // 应用ID
	AppSecret      string // 应用密钥
	MchId			string	//商户号
)

// Initialize 配置并初始化
func Initialize(appId, appSecret string) {
	if matched, err := regexp.MatchString("^wx[0-9a-f]{16}$", appId); err != nil || !matched {
		log.Fatalf("appId format error: %s", err)
	}
	if matched, err := regexp.MatchString("^[0-9a-f]{32}$", appSecret); err != nil || !matched {
		log.Fatalf("appSecret format error: %s", err)
	}

	AppId = appId         // 应用ID
	AppSecret = appSecret // 应用密钥
}
