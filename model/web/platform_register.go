// 自动生成模板PlatformRegister
package web

// PlatformRegister 结构体
type PlatformRegister struct {
	ID              uint   `json:"id" gorm:"primarykey"`
	BackgroundImage string `json:"background_image" form:"backgroundImage" gorm:"column:background_image;comment:背景图片;size:200;"`
	CreatedTime     int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Linked          string `json:"linked" form:"linked" gorm:"column:linked;comment:跳转链接;size:100;"`
	UpdatedTime     int64  `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
}

// TableName PlatformRegister 表名
func (PlatformRegister) TableName() string {
	return "platform_register"
}
