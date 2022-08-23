package web

type UserInfoDetail struct {
	ID              uint   `json:"id" gorm:"primarykey"`
	Uid             string `json:"uid" form:"uid" gorm:"column:uid;comment:用户UUID;size:15;"`
	TelegramAccount string `json:"telegram_account" form:"telegram_account" gorm:"column:telegram_account;comment:账号;size:50;"`
	Salt            string `json:"salt" form:"salt" gorm:"column:salt;comment:salt盐;size:50;"`
	Password        string `json:"password" form:"password" gorm:"column:password;comment:md5后的密码;size:50;"`
	CreatedTime     int64  `json:"created_time" form:"createdTime" gorm:"column:created_time;comment:创建时间;size:10;"`
	UpdatedTime     int64  `json:"updated_time" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
}

func (u *UserInfoDetail) TableName() string {
	return "user_info_detail"
}
