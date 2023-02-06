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
		api.POST("video", controller.Video)
	}
}
