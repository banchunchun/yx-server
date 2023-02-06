package service

import "fmt"

// TransCoderTaskVO transcoder 运行时参数对象
type TransCoderTaskVO struct {
	//TaskId 任务ID
	TaskId int64
	//FeatureXmlInput 算法特征xml路径
	FeatureXmlInput string
	//	StreamOutput 输出的视频路径
	StreamOutput string
	//StreamInput 输入的路径
	StreamInput string
	// Width 视频宽度: eg. 1920
	Width int64
	// Height 视频高度：eg. 1080
	Height int64
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

var transcoderService *TranscoderService

// NewTranscoderService 生成实例
func NewTranscoderService() *TranscoderService {
	if transcoderService == nil {
		transcoderService = &TranscoderService{}
	}
	return transcoderService
}

// StartTranscoder 启动TransCoder
func (ts *TranscoderService) StartTranscoder(vo *TransCoderTaskVO) {
	//var stdout io.ReadCloser

	fmt.Println("startTranscoder")

}
