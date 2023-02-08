package request

// YxCallBackRequest yx callback request
type YxCallBackRequest struct {
	TaskId  int    `json:"taskId"`
	Status  string `json:"status"`
	Process string `json:"process"`
}
