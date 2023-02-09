package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/global"
)

func Run() *gin.Engine {
	engine := initGin()
	loadRoute(engine)
	engine.Run(global.CF.Server.Port)
	global.LOG.Infof("yxServer run port:%v\n", global.CF.Server.Port)
	return engine
}

func initGin() *gin.Engine {
	engine := gin.New()
	engine.Use()
	return engine
}
func loadRoute(r *gin.Engine) {
	testRoute(r)
	apiRoute(r)
}
