package web

// ContactUs 结构体
type ContactUs struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Account     string `json:"account" form:"account" gorm:"column:account;comment:账号;size:100;"`
	AccountName string `json:"account_name" form:"accountName" gorm:"column:account_name;comment:账号名称;size:20;"`
	CreatedTime int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Icon        string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:200;"`
	UpdatedTime int64  `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
}

// TableName ContactUs 表名
func (ContactUs) TableName() string {
	return "contact_us"
}
