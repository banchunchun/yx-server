package response

// NoticeEventResponse 通知对象
type NoticeEventResponse struct {
	TaskId  int    `json:"taskId"`
	Status  string `json:"status"`
	Url     string `json:"url"`
	Process string `json:"process"`
}
