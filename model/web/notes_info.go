package web

// NotesInfo 结构体
type NotesInfo struct {
	ID          uint   `gorm:"primarykey"` // 主键ID
	Content     string `json:"content" form:"content" gorm:"column:content;comment:内容简介;size:200;"`
	CreatedTime int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Image       string `json:"image" form:"image" gorm:"column:image;comment:笔记图片;size:200;"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:笔记标题;size:50;"`
	UpdatedTime int64  `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
}

// TableName NotesInfo 表名
func (NotesInfo) TableName() string {
	return "notes_info"
}
