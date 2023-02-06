package request

// Base 公用的参数
type Base struct {
	// SourceFileType 源文件类型: VIDEO/AUDIO/IMAGE
	SourceFileType string `json:"sourceFileType"`
	//SourceFilePath 源文件路径，需要保证算法能访问的路径
	SourceFilePath string `json:"sourceFilePath"`
	// TargetFilePath 输出的文件路径
	TargetFilePath string `json:"targetFilePath"`
	//Command 操作命令，默认 COMPLETE
	Command string `json:"command"`
	// TaskId 任务id，需要业务层生成后续用来查询任务是否执行完毕
	TaskId int64 `json:"taskId"`
	// CallBackUrl  回调的url，如果不填则需要业务层主动轮询查
	CallBackUrl string `json:"callbackUrl"`
	//TaskStatus   任务状态
	TaskStatus string `json:"taskStatus"`
	// Process 任务进度 0-100
	Process int64 `json:"process"`
}
