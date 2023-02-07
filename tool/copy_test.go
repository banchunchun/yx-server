package tool

import (
	"fmt"
	"go-api/app/request"
	"go-api/app/service"
	"testing"
)

func TestSimpleCopyProperties(t *testing.T) {

	var matReq = request.VideoOutPutRequest{
		Width:  "1920",
		Height: "1080",
	}

	vo := service.TransCoderTaskVO{}

	SimpleCopyProperties(matReq, &vo)

	fmt.Println(vo)
}
