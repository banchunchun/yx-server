package request

// YxCallBackRequest yx callback request
type YxCallBackRequest struct {
	TaskId int    `json:"taskId"`
	Url    string `json:"url"`
	Status string `json:"status"`
}
