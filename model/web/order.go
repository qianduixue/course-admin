package web

// Order 结构体
type Order struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	Email          string         `json:"email" form:"email" gorm:"column:email;comment:邮箱账号;size:50;"`
	CollectionAddr string         `json:"collectionAddr" form:"collectionAddr" gorm:"column:collection_addr;comment:收款地址;size:50;"`
	CreatedTime    int64          `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:下单时间;size:10;"`
	Discord        string         `json:"discord" form:"discord" gorm:"column:discord;comment:discord账号;size:50;"`
	LevelId        int8           `json:"levelId" form:"levelId" gorm:"column:level_id;comment:会员级别id;"`
	OrderSn        string         `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:订单编号;size:25;"`
	PayVoucher     string         `json:"payVoucher" form:"payVoucher" gorm:"column:pay_voucher;comment:支付凭证;size:150;"`
	Price          int32          `json:"price" form:"price" gorm:"column:price;comment:价格;size:10;"`
	Telegram       string         `json:"telegram" form:"telegram" gorm:"column:telegram;comment:telegram账号;size:50;"`
	Type           int8           `json:"type" form:"type" gorm:"column:type;comment:会员有效期 1 月 2 季 3 年;"`
	Uid            string         `json:"uid" form:"uid" gorm:"column:uid;comment:用户UID;size:15;"`
	UpdatedTime    int64          `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:操作时间;size:10;"`
	LevelData      UserLevelPrice `gorm:"foreignKey:LevelId"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "order"
}
