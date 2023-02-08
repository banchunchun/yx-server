package tool

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-api/app/request"
	"go-api/app/service"
	"log"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	fileName := filepath.Join("../templates/", "test.tpl")
	mapInfo := make(map[string]interface{})
	mapInfo["test"] = "yes"
	Parse(fileName, "../templates/tmp.tpl", mapInfo)
}

func TestParseObject(t *testing.T) {
	fileName := filepath.Join("../templates/", "transcoder.tpl")
	tv := service.TransCoderTaskVO{
		TaskId:          11,
		FeatureXmlInput: "/home/a.custom",
		StreamInput:     "/home/a.mp4",
		StreamOutput:    "/home/b.mp4",
		Width:           "1920",
		Height:          "1080",
		Container:       "mp4",
		Codec:           "H264",
		BitRate:         "5000",
		AudioAacType:    "AAC",
		AudioChannel:    "2",
		AudioSampleRate: "441000",
	}
	err := ParseObject(fileName, "../templates/tmp.tpl", &tv)
	if err != nil {
		//global.LOG.Error()
		log.Panicln(err)
		return
	}
}

func TestTestReg(t *testing.T) {
	//TestReg()
	//var key = "60%"
	//value := strings.Replace(key, "%", "",1)
	//fmt.Println(value)
	//global.CF = core.LoadConfig()
	//global.LOG = initialize.ZapSugar("yxServer") //初始化日志
	//global.MYSQL = initialize.LoadMySQL()
	url := fmt.Sprintf("%s:%s", "172.17.81.21", "6379")
	redisClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     url,
		Password: "Arcvideo@01",
		DB:       13,
	})

	bv := request.Base{
		TaskId:         1,
		SourceFilePath: "/home/path",
	}
	Set(redisClient, "banchun", bv)
	var bvv request.Base
	Get(redisClient, "banchun", &bvv)
	fmt.Println(bvv.TaskId)
	fmt.Println(bvv.SourceFilePath)
}
