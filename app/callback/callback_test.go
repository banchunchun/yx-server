package callback

import (
	"go-api/core"
	"go-api/global"
	"go-api/initialize"
	"testing"
)

func TestCallBack(t *testing.T) {
	global.CF = core.LoadConfig()
	global.LOG = initialize.ZapSugar("yxServer") //初始化日志
	callbackUrl := "http://172.17.81.21:17704/api/yx/callback"
	CallBack(callbackUrl, "SUCCESS", "100", 48)
}
