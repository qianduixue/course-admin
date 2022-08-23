package web

// UserLevelPrice 结构体
type UserLevelPrice struct {
	ID          uint   `gorm:"primarykey"` // 主键ID
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Grade       int8   `json:"grade" form:"grade" gorm:"column:grade;comment:等级;size:10;"`
	Name        string `json:"name" form:"name" gorm:"column:name;comment:等级名称;size:15;"`
	Price       int32  `json:"price" form:"price" gorm:"column:price;comment:所需价格;size:10;"`
	Type        int8   `json:"type" form:"type" gorm:"column:type;comment:会员有效期 1 月 2 季 3 年;"`
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
}

// TableName UserLevelPrice 表名
func (UserLevelPrice) TableName() string {
	return "user_level_price"
}
