package callback

import (
	"bytes"
	"encoding/json"
	"go-api/app/request"
	"go-api/global"
	"go.uber.org/zap"
	"io"
	"net/http"
)

// CallBack 回调url
func CallBack(callbackUrl string, status string, process string, taskId int) {
	backRequest := request.YxCallBackRequest{
		TaskId:  taskId,
		Status:  status,
		Process: process,
	}
	postJson, _ := json.Marshal(backRequest)
	global.LOG.Infof("[callback] url=%v,postJson=%v\n", callbackUrl, string(postJson))
	post, err := http.Post(callbackUrl, "application/json", bytes.NewBuffer(postJson))
	if err != nil {
		global.LOG.Errorf("[callback] post url: %v error:%v\n", callbackUrl, zap.Error(err))
		return
	}
	responseBody(callbackUrl, &backRequest, post)
}

// responseBody 获取http 输出
func responseBody(url string, req *request.YxCallBackRequest, r *http.Response) {
	content, _ := io.ReadAll(r.Body)
	global.LOG.Infof("[callback] url=%v taskId=%d process=%v status=%v response=%v\n", url, req.TaskId, req.Process, req.Status, string(content))
}
