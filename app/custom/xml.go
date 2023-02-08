package custom

import (
	"encoding/xml"
	"go-api/app/request"
	"go-api/global"
	"os"
	"strconv"
)

// contentdetect 合成图片的xml对象
type contentdetect struct {
	Log      string   `xml:"log"`
	LiveMode string   `xml:"liveMode"`
	Features Features `xml:"features"`
}

type Features struct {
	Feature Feature `xml:"feature"`
	Type    string  `xml:"type,attr"`
}

type Feature struct {
	MaterialList MR `xml:"materialList"`
}

type MR struct {
	Material []Material `xml:"material"`
}

type Material struct {
	Path            string `xml:"path"`
	AppearDuration  string `xml:"appearDuration"`
	AppearFrequency string `xml:"appearFrequency"`
	Transparency    string `xml:"transparency"`
	StartPosition   string `xml:"startPosition"`
	EndPosition     string `xml:"endPosition"`
	Regions         []R    `xml:"regions"`
}

type R struct {
	Region request.Region `xml:"region"`
}

// BuildXml 生成图片的xml
func BuildXml(videoRequest request.VideoRequest) (string, error) {
	targetXmlPath := ""
	targetXmlName := strconv.Itoa(videoRequest.TaskId) + ".xml"
	if global.CF == nil {
		targetXmlPath = "../../templates/"
	} else {
		targetXmlPath = global.CF.Transcoder.XmlPath
	}
	err := os.MkdirAll(targetXmlPath, 0644)
	if err != nil {
		global.LOG.Errorf("BuildXml mkdir error %v %v", targetXmlPath, err)
		return "", err
	}
	err = os.Chdir(targetXmlPath)
	if err != nil {
		global.LOG.Errorf("BuildXml chdir  error %v %v", targetXmlPath, err)
		return "", err
	}
	var matList []Material
	for _, mr := range videoRequest.MaterialList {
		var regions []R

		for _, r := range mr.Regions {
			region := request.Region{
				X: r.X,
				Y: r.Y,
				W: r.W,
				H: r.H,
			}
			rr := R{
				region,
			}
			regions = append(regions, rr)
		}
		//regions = append(regions)
		mat := Material{
			Path:            mr.Path,
			AppearDuration:  strconv.Itoa(mr.AppearDuration),
			AppearFrequency: strconv.Itoa(mr.AppearFrequency),
			StartPosition:   strconv.Itoa(mr.StartPosition),
			EndPosition:     strconv.Itoa(mr.EndPosition),
			Transparency:    strconv.Itoa(mr.Transparency),
			Regions:         regions,
		}
		matList = append(matList, mat)
	}
	mr := MR{
		Material: matList,
	}
	ff := Feature{
		MaterialList: mr,
	}

	features := Features{
		Feature: ff,
		Type:    "blendImages",
	}

	con := contentdetect{
		Log:      "debug",
		LiveMode: "0",
		Features: features,
	}

	file, _ := xml.MarshalIndent(&con, "", "\t")
	file = append([]byte(xml.Header), file...)

	err = os.WriteFile(targetXmlName, file, 0644)
	if err != nil {
		global.LOG.Errorf("BuildXml OS create error %v ", err)
		return targetXmlPath, err
	}
	return targetXmlPath + "/" + targetXmlName, nil
}
