package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/tool"
)

type root struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func testRoute(r *gin.Engine) {
	test := r.Group("/api/test")
	{
		test.GET("/home", func(context *gin.Context) {
			tool.WriteJson(context, 200, "操作成功", "first")
		})
	}
}
