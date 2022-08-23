// 自动生成模板MemberInterests
package web

// MemberInterests 结构体
type MemberInterests struct {
	ID          uint   `json:"id" gorm:"primarykey"` // 主键ID
	CreatedTime int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Linked      string `json:"linked" form:"linked" gorm:"column:linked;comment:视频链接;size:150;"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:标题;size:30;"`
	UpdatedTime int64  `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
}

// TableName MemberInterests 表名
func (MemberInterests) TableName() string {
	return "member_interests"
}
