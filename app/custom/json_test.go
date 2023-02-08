package custom

import (
	"go-api/app/request"
	"testing"
)

func TestBuildImageJson(t *testing.T) {
	region := request.Region{
		X: 10,
		Y: 10,
		W: 120,
		H: 120,
	}
	var regions []request.Region
	regions = append(regions, region)
	mat := request.ImageMaterialRequest{
		Path:         "/home",
		Regions:      regions,
		Transparency: 20,
	}
	var imagesList []*request.ImageMaterialRequest
	imagesList = append(imagesList, &mat)
	base := request.Base{
		TaskId:         12,
		SourceFileType: "IMAGE",
	}
	imageRequest := request.ImageRequest{
		Base:         base,
		MaterialList: imagesList,
	}

	BuildImageJson(imageRequest)
}
