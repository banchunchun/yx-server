package service

import (
	"bufio"
	"fmt"
	"go-api/app/callback"
	"go-api/app/request"
	"go-api/constants"
	"go-api/global"
	"go-api/tool"
	"go.uber.org/zap"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// TransCoderTaskVO transcoder 运行时参数对象
type TransCoderTaskVO struct {
	//TaskId 任务ID
	TaskId int
	//FeatureXmlInput 算法特征xml路径
	FeatureXmlInput string
	//	StreamOutput 输出的视频路径
	StreamOutput string
	//StreamInput 输入的路径
	StreamInput string
	// Width 视频宽度: eg. 1920
	Width string
	// Height 视频高度：eg. 1080
	Height string
	// Container 视频编码格式 eg: MP4
	Container string
	// Codec 视频编码： eg. H264
	Codec string
	//BitRate 视频编码码率： eg. 5000, 无需单位，底层默认kbps
	BitRate string
	// AudioAacType 音频输出AAC: eg. AAC
	AudioAacType string
	//AudioBitRate  音频输出：eg. 64
	AudioBitRate string
	//音频声道： eg. 2
	AudioChannel string
	//AudioSampleRate 音频输出采样率： eg. 44.1
	AudioSampleRate string
}
type TranscoderService struct {
}

var transcoderHelper *TranscoderService
var reg = regexp.MustCompile("\\d+%")

// NewTranscoderService 生成实例
func NewTranscoderService() *TranscoderService {
	if transcoderHelper == nil {
		transcoderHelper = &TranscoderService{}
	}
	return transcoderHelper
}

// StartTranscoder 启动TransCoder
func (transcoderHelper *TranscoderService) StartTranscoder(vo *TransCoderTaskVO, xmlPath string) {
	var stdout io.ReadCloser
	global.LOG.Info("[transcoder] begin startTranscoder")
	var cmdFmt = "./transcoder.exe %s -g -n"
	var targetCmd = fmt.Sprintf(cmdFmt, xmlPath)
	var err error
	global.LOG.Infof("[transcoder] exe cmd %v", targetCmd)
	cmd := exec.Command("/bin/bash", "-c", targetCmd)
	f := func() {
		cmd.Dir = global.CF.Transcoder.WorkDir
		stdout, err = cmd.StdoutPipe()
		if err != nil {
			global.LOG.Errorf("stdout error,%d,%v", vo.TaskId, zap.Error(err))
			return
		}
		cmd.Stderr = cmd.Stdout
		if err := cmd.Start(); err != nil {
			global.LOG.Errorf("[transcoder] 启动失败:cmd.Start error %+v \n", err)
			return
		}
		pid := cmd.Process.Pid
		global.LOG.Infof("[transcoder] Start Success taskId: %d pid:%d \n", vo.TaskId, pid)
	}
	f()
	go addListenerStdout(stdout, vo.TaskId, vo)
}

// addListenerStdout 监听transcoder的任务输出
func addListenerStdout(stdout io.ReadCloser, key int, vo *TransCoderTaskVO) {
	reader := bufio.NewReader(stdout)
	for {
		tempLine, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			global.LOG.Errorf("[addListenerStdout] error break %v", zap.Error(err))
			break
		}
		line := string(tempLine)
		global.LOG.Infoln("[addListenerStdout] 输出 [" + line + "]")
		go doStdoutBusinessForPlay(line, key)
	}
}

// doStdoutBusinessForPlay 处理transcoder内部业务逻辑
// hang的业务暂未处理
func doStdoutBusinessForPlay(line string, key int) {
	//判断启动
	if strings.Contains(line, constants.TRANSCODER_STARTED) || strings.Contains(line, constants.TRANSCODER_RUNNING) {
		updateTranscoderStatus(key, "0", constants.STATUS_RUNNING)
	} else if strings.Contains(line, constants.TRANSCODER_COMPLETE) {
		//完成
		updateTranscoderStatus(key, "100", constants.STATUS_SUCCESS)
	} else if strings.Contains(line, constants.STATUS_ERROR) {
		//失败
		updateTranscoderStatus(key, "0", constants.STATUS_ERROR)
	} else if strings.Contains(line, constants.TRANSCODER_HANG_ERROR) {
		// HANG ，需要执行强杀并发送失败信息
		updateTranscoderStatus(key, "0", constants.STATUS_ERROR)
	} else if strings.Contains(line, constants.TRAMSCPDER_PROCESS) {
		var vals []string
		vals = reg.FindAllString(line, -1)
		if len(vals) == 0 {

		} else {
			for _, val := range vals {
				value := strings.Replace(val, "%", "", 1)
				updateTranscoderStatus(key, value, constants.STATUS_RUNNING)
			}
		}
	}

}

// updateTranscoderStatus 更新transcoder状态，并设置redis的值
func updateTranscoderStatus(taskId int, process string, status string) {
	global.LOG.Infof("[transcoder] taskId=%v status=%v process=%v\n", taskId, status, process)
	//执行回调，并更新redis值
	redisTaskKey := fmt.Sprintf(constants.YX_KEY_PREFIX, taskId)
	var vr request.VideoRequest
	//设置redis值
	tool.Get(global.Redis, redisTaskKey, &vr)
	global.LOG.Infof("updateTranscoderStatus taskId=%v rds=%v\n", taskId, vr.TaskId)
	pc, _ := strconv.Atoi(process)
	vr.Process = pc
	vr.TaskStatus = status
	tool.Set(global.Redis, redisTaskKey, vr)
	//notify callback url
	callbackUrl := vr.CallBackUrl
	if callbackUrl != "" {
		go callback.CallBack(callbackUrl, status, process, taskId)
	}

}
