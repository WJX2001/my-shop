package utils

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// SendMesseageCode 发送手机验证码。
// 配置 sms_mock=true 时走开发 Mock：只打日志，不调阿里云。
func SendMessageCode(phone string, verify_code int) bool {
	if beego.AppConfig.DefaultBool("sms_mock", false) {
		logs.Info("[SMS Mock] phone=%s code=%d (开发环境，未真实发送)", phone, verify_code)
		return true
	}

	return false
}
