package tool

import "github.com/gin-gonic/gin"

type ResultData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
}

// WriteSuccessJson 输出成功的信息至Response
func WriteSuccessJson(context *gin.Context, code int, data interface{}) {
	result := ResultData{
		Code:    code,
		Message: "操作成功",
		Result:  data,
		Success: code == 0,
	}
	context.JSONP(code, result)
}

// WriteErrorJson 输出失败的信息至Response
func WriteErrorJson(context *gin.Context, code int) {
	result := ResultData{
		Code:    code,
		Message: "操作失败",
		Success: code == 0,
	}
	context.JSONP(code, result)
}
