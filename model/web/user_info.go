package web

// UserInfo 结构体
type UserInfo struct {
	ID                   uint           `json:"id" form:"id" gorm:"primarykey"`
	CreatedTime          int64          `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	Discord              string         `json:"discord" form:"discord" gorm:"column:discord;comment:discord账号;size:100;"`
	Email                string         `json:"email" form:"email" gorm:"column:email;comment:邮箱地址;size:50;"`
	HeadPortrait         string         `json:"head_portrait" form:"headPortrait" gorm:"column:head_portrait;comment:头像;size:200;"`
	Level                int8           `json:"level" form:"level" gorm:"column:level;comment:用户等级;size:10;"`
	MembershipExpireTime int64          `json:"membership_expire_time" form:"membershipExpireTime" gorm:"column:membership_expire_time;comment:会员到期时间;size:10;"`
	RegType              int8           `json:"reg_type" form:"regType" gorm:"column:reg_type;comment:注册类别 1 邮箱 2 discord;"`
	Uid                  string         `json:"uid" form:"uid" gorm:"column:uid;comment:用户UUID;size:15;"`
	UpdatedTime          int64          `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
	IpAddr               string         `json:"ip_addr" form:"ip_addr" gorm:"column:ip_addr;comment:ip地址;size:50;"`
	DetailData           UserInfoDetail `gorm:"foreignKey:Uid"`
}

// TableName UserInfo 表名
func (UserInfo) TableName() string {
	return "user_info"
}
