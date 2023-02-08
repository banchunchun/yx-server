package custom

import (
	"encoding/json"
	"go-api/app/request"
	"go-api/global"
	"os"
	"strconv"
)

// BuildImageJson 生成图片合成的json文件
func BuildImageJson(imageRequest request.ImageRequest) (string, error) {
	targetJsonPath := ""
	targetJsonName := strconv.Itoa(imageRequest.TaskId) + ".json"
	if global.CF == nil {
		targetJsonPath = "../../templates/"
	} else {
		targetJsonPath = global.CF.Transcoder.XmlPath
	}
	err := os.MkdirAll(targetJsonPath, 0644)
	if err != nil {
		global.LOG.Errorf("BuildImageJson mkdir error %v %v", targetJsonPath, err)
		return "", err
	}
	err = os.Chdir(targetJsonPath)
	if err != nil {
		global.LOG.Errorf("BuildImageJson chdir  error %v %v", targetJsonPath, err)
		return "", err
	}
	bv, err := json.Marshal(imageRequest)
	err = os.WriteFile(targetJsonName, bv, 0644)
	if err != nil {
		global.LOG.Errorf("BuildXml OS create error %v ", err)
		return targetJsonPath, err
	}
	return targetJsonPath + "/" + targetJsonName, nil
}
