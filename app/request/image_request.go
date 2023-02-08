package request

// ImageRequest 图片合成的请求数据
type ImageRequest struct {
	Base
	// MaterialList 待植入的数据
	MaterialList []*ImageMaterialRequest `json:"materialList"`
}

// ImageMaterialRequest 视频的透明度
type ImageMaterialRequest struct {
	// Path 待植入素材路径
	Path string `json:"path"`
	// Transparency 透明度，默认20
	Transparency int      `json:"transparency"`
	Regions      []Region `json:"regions"`
}
