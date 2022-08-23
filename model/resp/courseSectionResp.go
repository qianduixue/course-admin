package resp

type CourseSectionResp struct {
	Id          uint   `json:"id"`
	Lessons     string `json:"lessons"`      // 课节
	Tittle      string `json:"tittle"`       // 课节标题
	LongTime    int32  `json:"long_time"`    // 视频时间
	CreatedTime int64  `json:"created_time"` // 上传时间
}
