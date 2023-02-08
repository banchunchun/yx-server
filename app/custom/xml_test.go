package custom

import (
	"encoding/xml"
	"go-api/app/request"
	"log"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	mat := Material{
		Path: "/home",
	}
	region := request.Region{
		X: 10,
		Y: 10,
		W: 120,
		H: 120,
	}
	rr := R{
		region,
	}
	var regions []R
	regions = append(regions, rr)
	mat.Regions = regions
	//mr := MR{
	//	mat,
	//}
	//var matList []MR
	//matList = append(matList, mr)

	ff := Feature{
		//MaterialList: matList,
		Type: "blendImages",
	}

	features := Features{
		Feature: ff,
	}

	con := contentdetect{
		Log:      "debug",
		LiveMode: "0",
		Features: features,
	}

	file, _ := xml.MarshalIndent(&con, "", "\t")
	file = append([]byte(xml.Header), file...)
	outFileName := "../../templates/11.xml"
	_, err := os.Create(outFileName)
	if err != nil {
		log.Panicln(err)
		return
	}
	err = os.WriteFile(outFileName, file, 0644)
	if err != nil {
		log.Panicln(err)
		return
	}
}
