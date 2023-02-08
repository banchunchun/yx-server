package service

import (
	"fmt"
	"go-api/app/callback"
	"go-api/app/request"
	"go-api/constants"
	"go-api/global"
	"go-api/tool"
	"go.uber.org/zap"
	"os/exec"
)

type ImageService struct {
}

var imageHelper *ImageService

// NewImageService 生成实例
func NewImageService() *ImageService {
	if imageHelper == nil {
		imageHelper = &ImageService{}
	}
	return imageHelper
}

// BuildImage 制作图片
func (imageHelper *ImageService) BuildImage(imageRequest request.ImageRequest, jsonPath string) {
	var cmdFmt = "./imageBlend %s"
	var targetCmd = fmt.Sprintf(cmdFmt, jsonPath)
	global.LOG.Infof("[imageBlend] exe cmd %v", targetCmd)
	cmd := exec.Command("/bin/bash", "-c", targetCmd)
	f := func() {
		cmd.Dir = global.CF.Transcoder.ImageExeWorkDir
		err := cmd.Run()
		if err != nil {
			global.LOG.Errorf("[imageBlend] run error,%d,%v", imageRequest.TaskId, zap.Error(err))
			return
		}
		pid := cmd.Process.Pid
		global.LOG.Infof("[imageBlend] Start Success taskId: %d pid:%d \n", imageRequest.TaskId, pid)
		exitCode := cmd.ProcessState.ExitCode()
		if exitCode == 0 {
			global.LOG.Infof("[imageBlend] run Success taskId: %d pid:%d  exitCode:%d \n", imageRequest.TaskId, pid, exitCode)
			doUpdateRedis(imageRequest, constants.STATUS_SUCCESS, 100)
			doCallBack(imageRequest, constants.STATUS_SUCCESS, "100")
		} else {
			global.LOG.Infof("[imageBlend] run error taskId: %d pid:%d  exitCode:%d \n", imageRequest.TaskId, pid, exitCode)
			doUpdateRedis(imageRequest, constants.STATUS_ERROR, 0)
			doCallBack(imageRequest, constants.STATUS_ERROR, "0")
		}
	}
	f()
}
func doUpdateRedis(imageRequest request.ImageRequest, status string, process int) {
	//执行回调，并更新redis值
	redisTaskKey := fmt.Sprintf(constants.YX_KEY_PREFIX, imageRequest.TaskId)
	imageRequest.TaskStatus = status
	imageRequest.Process = process
	tool.Set(global.Redis, redisTaskKey, imageRequest)
}

// doCallBack 回调
func doCallBack(imageRequest request.ImageRequest, status string, process string) {
	//notify callback url
	callbackUrl := imageRequest.CallBackUrl
	if callbackUrl != "" {
		go callback.CallBack(callbackUrl, status, process, imageRequest.TaskId)
	}
}
