package resp

type OrderResp struct {
	Id          uint   `json:"id"`           // id
	Email       string `json:"email"`        // 邮箱
	Discord     string `json:"discord"`      // discord 账号
	Telegram    string `json:"telegram"`     // telegram 账号
	OrderSn     string `json:"order_sn"`     // 订单编号
	Price       int32  `json:"price"`        // 所需金额
	PayVoucher  string `json:"pay_voucher"`  // 支付凭证
	LevelName   string `json:"level_name"`   // 等级名称
	CreatedTime int64  `json:"created_time"` // 下单时间
}
