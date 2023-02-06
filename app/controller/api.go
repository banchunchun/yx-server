package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/app/request"
	"go-api/app/service"
	"go-api/tool"
)

// Video 视频合成
func Video(context *gin.Context) {

	var videoRequest request.VideoRequest
	context.ShouldBind(&videoRequest)
	//global.LOG.Log("video request data=%s", vr)

	fmt.Printf("video request data=%v\n", videoRequest)
	fmt.Println("数组内的数据：" + videoRequest.MaterialList[0].Path)
	//ts := service.NewTranscoderService()
	ts := service.NewTranscoderService()

	transCoderTaskVO := service.TransCoderTaskVO{TaskId: 1}
	//启动协程进行运行
	go ts.StartTranscoder(&transCoderTaskVO)
	tool.WriteJson(context, 200, "操作成功", videoRequest)
}
