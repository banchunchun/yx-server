package main

import (
	"database/sql"
	"go-api/core"
	"go-api/global"
	"go-api/initialize"
	"go-api/server"
	"log"
)

func init() {

}

/*
*
 */
func main() {
	log.Println("begin start yx server")
	//path := "./configs/"
	//config := core.LoadConfig(path)
	//load config
	global.Config = core.LoadConfig()
	log.Println("load configs files")
	log.Println(global.Config)
	log.Println(global.Config.Mysql.Host)
	log.Println("load configs end")
	global.MYSQL = initialize.LoadMySQL()
	global.Redis = initialize.LoadRedis()
	//主进程结束前关闭数据库链接
	sqlDB, _ := global.MYSQL.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	server.RunServer()

}
