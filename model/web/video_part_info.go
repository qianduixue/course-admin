package web

type VideoPartInfo struct {
	ID          uint   `gorm:"primarykey"`
	UploadId    string `json:"upload_id" form:"upload_id" gorm:"column:upload_id;comment:上传id;"`
	Etag        string `json:"etag" form:"etag" gorm:"column:etag;comment:分段视频标记;"`
	PartNumber  int64  `json:"part_number" form:"part_number" gorm:"column:part_number;comment:分段id,size:10;"`
	Key         string `json:"key" form:"key" gorm:"column:key;comment:视频路径;size:10;"`
	CreatedTime int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
}

func (v *VideoPartInfo) TableName() string {
	return "video_part_info"
}
