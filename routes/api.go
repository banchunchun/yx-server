package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/app/controller"
)

// apiRoute 封装yx server api的接口，暂时不定义后台api
func apiRoute(context *gin.Engine) {

	// api
	api := context.Group("/api/yx")
	{
		//提交视频制作任务
		api.POST("video", controller.Video)
		//提交图片制作任务
		api.POST("image", controller.Image)
		//提交音频制作任务
		api.POST("audio", controller.Audio)
		//根据任务id获取任务的状态
		api.GET("result", controller.Result)
	}
}
