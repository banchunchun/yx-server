package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/app/custom"
	"go-api/app/request"
	"go-api/app/service"
	"go-api/constants"
	"go-api/global"
	"go-api/tool"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
)

// Video 视频合成
func Video(context *gin.Context) {

	var videoRequest request.VideoRequest
	err := context.ShouldBind(&videoRequest)
	if err != nil {
		global.LOG.Errorf("[VIDEO] bind data error %v", zap.Error(err))
		tool.WriteErrorJson(context, -1)
		return
	}
	redisTaskKey := fmt.Sprintf(constants.YX_KEY_PREFIX, videoRequest.TaskId)

	tool.Set(global.Redis, redisTaskKey, videoRequest)

	global.LOG.Infof("video request data=%v\n", videoRequest)
	// 生成 transcoder helper
	transcoderHelper := service.NewTranscoderService()

	transCoderTaskVO := service.TransCoderTaskVO{
		TaskId: videoRequest.TaskId,
	}
	err = tool.SimpleCopyProperties(videoRequest.VideoOutPutRequest, &transCoderTaskVO)
	transCoderTaskVO.StreamInput = videoRequest.SourceFilePath
	transCoderTaskVO.StreamOutput = videoRequest.TargetFilePath
	if err != nil {
		global.LOG.Errorf("[VIDEO] copy properties error %v", zap.Error(err))
		tool.WriteErrorJson(context, -1)
		return
	}
	xmlPath, err := custom.BuildXml(videoRequest)
	if err != nil {
		tool.WriteErrorJson(context, -1)
		return
	}
	global.LOG.Infof("[VIDEO] build xmlPath=%v", xmlPath)
	transCoderTaskVO.FeatureXmlInput = xmlPath
	//生成transcoder.xml
	templateFileName := filepath.Join("../templates/", "transcoder.tpl")
	err = os.MkdirAll(global.CF.Transcoder.XmlPath+"/tpl/", 0644)
	if err != nil {
		global.LOG.Errorf("[VIDEO] create transcoder xml folder error %v", zap.Error(err))
		tool.WriteErrorJson(context, -1)
		return
	}
	transcoderXml := global.CF.Transcoder.XmlPath + "/tpl/" + strconv.Itoa(videoRequest.TaskId) + ".xml"
	global.LOG.Infof("[VIDEO] transcoderXml=%v", transcoderXml)
	err = tool.ParseObject(templateFileName, transcoderXml, &transCoderTaskVO)
	if err != nil {
		global.LOG.Errorf("[VIDEO] make transcoder.xml error %v", zap.Error(err))
		tool.WriteErrorJson(context, -1)
		return
	}
	//启动协程进行运行
	go transcoderHelper.StartTranscoder(&transCoderTaskVO, transcoderXml)
	//返回任务id
	tool.WriteSuccessJson(context, 200, videoRequest.TaskId)
}
