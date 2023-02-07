package tool

import "github.com/gin-gonic/gin"

type ResultData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// WriteSuccessJson 输出成功的信息至Response
func WriteSuccessJson(context *gin.Context, code int, data interface{}) {
	result := ResultData{
		Code:    code,
		Message: "操作成功",
		Data:    data,
	}
	context.JSONP(code, result)
}

// WriteErrorJson 输出失败的信息至Response
func WriteErrorJson(context *gin.Context, code int) {
	result := ResultData{
		Code:    code,
		Message: "操作失败",
	}
	context.JSONP(code, result)
}
