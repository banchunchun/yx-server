package callback

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api/app/request"
	"go-api/global"
	"go.uber.org/zap"
	"net/http"
)

// CallBack 回调url
func CallBack(callbackUrl string, status string, targetUrl string, taskId int) {
	backRequest := request.YxCallBackRequest{
		TaskId: taskId,
		Status: status,
		Url:    targetUrl,
	}
	postJson, _ := json.Marshal(backRequest)
	fmt.Println(postJson)
	post, err := http.Post(callbackUrl, "application/json", bytes.NewBuffer(postJson))
	global.LOG.Infoln("[callback] url:%v,post:%v", callbackUrl, post)
	if err != nil {
		global.LOG.Errorf("[callback] post url: %v error:%v\n", callbackUrl, zap.Error(err))
		return
	}
}
