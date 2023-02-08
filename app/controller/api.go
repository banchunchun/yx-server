package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/app/custom"
	"go-api/app/request"
	"go-api/app/response"
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
	fmt.Println(xmlPath)
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
	tool.WriteSuccessJson(context, 0, videoRequest.TaskId)
}

// Image  图片合成
func Image(context *gin.Context) {
	var imageRequest request.ImageRequest
	err := context.ShouldBind(&imageRequest)
	if err != nil {
		global.LOG.Errorf("[Image] bind data error %v", zap.Error(err))
		tool.WriteErrorJson(context, constants.PARAMS_BIND_ERROR)
		return
	}

	redisTaskKey := fmt.Sprintf(constants.YX_KEY_PREFIX, imageRequest.TaskId)

	tool.Set(global.Redis, redisTaskKey, imageRequest)

	global.LOG.Infof("image request data=%v\n", imageRequest)

	jsonPath, err := custom.BuildImageJson(imageRequest)
	if err != nil {
		tool.WriteErrorJson(context, -1)
		return
	}
	global.LOG.Infof("[IMAGE] build jsonPath=%v", jsonPath)
	//TODO 异步返回图片制作
	imageHelper := service.NewImageService()
	go imageHelper.BuildImage(imageRequest, jsonPath)
	//返回任务id
	tool.WriteSuccessJson(context, 0, imageRequest.TaskId)
}

// Audio 音频合成
func Audio(context *gin.Context) {
	tool.WriteErrorJson(context, -1)
}

// Result 根据任务id和taskType获取结果
func Result(context *gin.Context) {
	queryTaskId, exists := context.GetQuery(constants.NOTIFY_TASK_ID)
	if !exists {
		tool.WriteErrorJson(context, constants.PARAMS_NOT_EXISTS_ERROR)
		return
	}
	taskType, exists := context.GetQuery(constants.NOTIFY_TASK_TYPE)
	if !exists {
		tool.WriteErrorJson(context, constants.PARAMS_NOT_EXISTS_ERROR)
		return
	}
	global.LOG.Infof("[result] taskId=%v taskType=%v\n", queryTaskId, taskType)
	taskId, _ := strconv.Atoi(queryTaskId)
	redisTaskKey := fmt.Sprintf(constants.YX_KEY_PREFIX, taskId)
	//通知对象
	notice := response.NoticeEventResponse{}
	switch taskType {
	case "VIDEO":
		var video request.VideoRequest
		tool.Get(global.Redis, redisTaskKey, &video)
		notice.TaskId = video.TaskId
		notice.Process = strconv.Itoa(video.Process)
		notice.Url = video.TargetFilePath
		notice.Status = video.TaskStatus
		break
	case "IMAGE":
		var video request.ImageRequest
		tool.Get(global.Redis, redisTaskKey, &video)
		notice.TaskId = video.TaskId
		notice.Process = strconv.Itoa(video.Process)
		notice.Url = video.TargetFilePath
		notice.Status = video.TaskStatus
		break
	case "AUDIO":
		global.LOG.Errorf("[result] taskId=%v taskType=%v\n unsupported", queryTaskId, taskType)
		tool.WriteErrorJson(context, constants.PARAMS_NOT_SUPPORT)
		break
	default:
		global.LOG.Errorf("[result] taskId=%v taskType=%v\n unsupported", queryTaskId, taskType)
		tool.WriteErrorJson(context, constants.PARAMS_NOT_SUPPORT)
		break
	}
	global.LOG.Infof("[result] taskId=%v notice event value=%v\n", queryTaskId, notice)
	tool.WriteSuccessJson(context, 0, notice)

}
