package course

// CourseInfo 结构体
type CourseInfo struct {
	ID            uint   `gorm:"primarykey"` // 主键ID
	Abstract      string `json:"abstract" form:"abstract" gorm:"column:abstract;comment:课程简介;size:100;"`
	CreatedTime   int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Desc          string `json:"desc" form:"desc" gorm:"column:desc;comment:标题描述;size:15;"`
	Level         int8   `json:"level" form:"level" gorm:"column:level;comment:课程级别 0 免费 1 一级会员 2 二级会员;"`
	Status        int8   `json:"status" form:"status" gorm:"column:status;comment:状态  1 正常 2 已删除;"`
	SectionNumber int8   `json:"section_number" form:"sectionNumber" gorm:"column:section_number;comment:课程小节数量;size:10;"`
	Title         string `json:"title" form:"title" gorm:"column:title;comment:课程标题;size:15;"`
	UpdatedTime   int64  `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
}

// TableName CourseInfo 表名
func (CourseInfo) TableName() string {
	return "course_info"
}
