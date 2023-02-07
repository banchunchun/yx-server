package request

// VideoRequest 视频合成的请求数据
type VideoRequest struct {
	Base
	// VideoOutPutRequest 视频输出参数配置
	VideoOutPutRequest VideoOutPutRequest `json:"videoOutputRequest"`
	// MaterialList 待植入的数据
	MaterialList []*VideoMaterialRequest `json:"materialList"`
}

// VideoMaterialRequest 视频的透明度
type VideoMaterialRequest struct {
	// Path 待植入素材路径
	Path string `json:"path"`
	// AppearDuration 出现时长：统一单位为秒，由前端转换
	AppearDuration int `json:"appearDuration"`
	//AppearFrequency 出现频率
	AppearFrequency int `json:"appearFrequency"`
	// Transparency 透明度，默认20
	Transparency int `json:"transparency"`
	//StartPosition 开始时间戳，到毫秒
	StartPosition int `json:"startPosition"`
	// 结束时间戳，到毫秒
	EndPosition int       `json:"endPosition"`
	Regions     []*Region `json:"regions"`
}

// VideoOutPutRequest 视频输出参数配置
type VideoOutPutRequest struct {
	// Width 视频宽度: eg. 1920
	Width int `json:"width"`
	// Height 视频高度：eg. 1080
	Height int `json:"height"`
	// Container 视频编码格式 eg: MP4
	Container string `json:"container"`
	// Codec 视频编码： eg. H264
	Codec string `json:"codec"`
	//BitRate 视频编码码率： eg. 5000, 无需单位，底层默认kbps
	BitRate string `json:"bitRate"`
	// AudioAacType 音频输出AAC: eg. AAC
	AudioAacType string `json:"audioAacType"`
	//AudioBitRate  音频输出：eg. 64
	AudioBitRate string `json:"audioBitRate"`
	//音频声道： eg. 2
	AudioChannel string `json:"audioChannel"`
	//AudioSampleRate 音频输出采样率： eg. 44.1
	AudioSampleRate string `json:"audioSampleRate"`
}
