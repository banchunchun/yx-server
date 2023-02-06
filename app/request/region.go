package request

// Region 坐标系
type Region struct {
	// X 左上角X坐标，基于当前比例的绝对坐标
	X int `json:"x"`
	// Y 左上角Y坐标，基于当前比例的绝对坐标
	Y int `json:"y"`
	//W 当前画框的宽度
	W int `json:"w"`
	//H 当前画框的高度
	H int `json:"h"`
}
