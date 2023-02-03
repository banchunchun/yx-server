package tool

import "github.com/gin-gonic/gin"

type ResultData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteJson(context *gin.Context, code int, message string, data interface{}) {
	result := ResultData{
		Code:    code,
		Message: message,
		Data:    data,
	}
	context.JSONP(code, result)
}
