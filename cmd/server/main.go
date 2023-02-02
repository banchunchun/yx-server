package main

import (
	"com.arcvideo.yx.server/pkg/conf"
	"log"
)

func init() {

}

/*
*
 */
func main() {
	log.Println("begin start yx server")
	path := "./configs/"
	config := conf.LoadConfig(path)
	log.Println("load configs files")
	log.Println(config)
	log.Println(config.Mysql.Host)
	log.Println("load configs end")

}
