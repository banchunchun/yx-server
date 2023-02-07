package main

import (
	"database/sql"
	"fmt"
	"go-api/core"
	"go-api/global"
	"go-api/initialize"
	"go-api/server"
)

func init() {

}

/*
*
 */
func main() {
	global.CF = core.LoadConfig()
	global.LOG = initialize.ZapSugar("yxServer") //初始化日志
	global.MYSQL = initialize.LoadMySQL()
	global.Redis = initialize.LoadRedis()
	fmt.Printf("TRANSCODER: %v\n", global.CF.Transcoder)
	//主进程结束前关闭数据库链接
	sqlDB, _ := global.MYSQL.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)

	server.RunServer()

}
