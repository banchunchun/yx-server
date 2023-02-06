package tool

import (
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
		FeatureXmlInput: "/home/a.xml",
		StreamInput:     "/home/a.mp4",
		StreamOutput:    "/home/b.mp4",
		Width:           1920,
		Height:          1080,
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
