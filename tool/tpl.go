package tool

import (
	"fmt"
	"go-api/global"
	"os"
	"regexp"
	"text/template"
)

// Parse 解析template
func Parse(fileName string, outFileName string, mapInfo map[string]interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		global.LOG.Errorf("failed to parse tpl %s %s", fileName, err.Error())
		return err
	}
	f, err := os.Create(outFileName)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			global.LOG.Errorf("close file %s %s", fileName, err.Error())
		}
	}(f)
	if err != nil {
		global.LOG.Errorf("failed to open writer %s %s", outFileName, err.Error())
		return err
	}
	err = t.Execute(f, mapInfo)
	return err
}

// ParseObject 通过结构体对象透传解析模板对象
func ParseObject(fileName string, outFileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		global.LOG.Error("failed to parse tpl %s %s", fileName, err.Error())
		return err
	}
	f, err := os.Create(outFileName)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			global.LOG.Error("close file %s %s", fileName, err.Error())
		}
	}(f)
	if err != nil {
		global.LOG.Error("failed to open writer %s %s", outFileName, err.Error())
		return err
	}
	err = t.Execute(f, data)
	return err
}

func TestReg() {
	ss := " [\n \\ [ 60% ] finished, speed [ 6.74 ], duration [ 202772(ms) ]\n | [ 61% ] finished, speed [ 6.76 ], duration [ 202772(ms) ]]"
	reg := regexp.MustCompile("\\d+\\(?=%\\)")
	var ssss []string
	ssss = reg.FindAllString(ss, -1)
	fmt.Println(ssss)
}
