package callback

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	fmt.Println(postJson)
	post, err := http.Post(callbackUrl, "application/json", bytes.NewBuffer(postJson))
	global.LOG.Infof("[callback] url:%v,post:%v\n", callbackUrl, post)
	if err != nil {
		global.LOG.Errorf("[callback] post url: %v error:%v\n", callbackUrl, zap.Error(err))
		return
	}
	responseBody(&backRequest, post)
}

// responseBody 获取http 输出
func responseBody(req *request.YxCallBackRequest, r *http.Response) {
	content, _ := io.ReadAll(r.Body)
	global.LOG.Infof("[callback] url=%v response=%v\n", req.TaskId, content)
}
