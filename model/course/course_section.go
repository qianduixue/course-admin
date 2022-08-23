package course

// CourseSection 结构体
type CourseSection struct {
	ID          uint   `gorm:"primarykey"` // 主键ID
	CourseId    int64  `json:"course_id" form:"courseId" gorm:"column:course_id;comment:课程id;size:19;"`
	CoverImage  string `json:"cover_image" form:"coverImage" gorm:"column:cover_image;comment:课节封面;size:200;"`
	CreatedTime int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Lessons     string `json:"lessons" form:"lessons" gorm:"column:lessons;comment:课节;size:20;"`
	Status      int8   `json:"status" form:"status" gorm:"column:status;comment:状态  1 正常 2 已删除;"`
	LongTime    int32  `json:"long_time" form:"longTime" gorm:"column:long_time;comment:课节时长;size:10;"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:课节标题;size:50;"`
	UpdatedTime int64  `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
	VideoUrl    string `json:"video_url" form:"videoUrl" gorm:"column:video_url;comment:课节视频地址;size:100;"`
}

// TableName CourseSection 表名
func (CourseSection) TableName() string {
	return "course_section"
}
