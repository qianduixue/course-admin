package resp

type UserInfoResp struct {
	Id        uint   `json:"id"`         // 用户id
	UID       string `json:"uid"`        // 用户uid
	Email     string `json:"email"`      // 邮箱
	Discord   string `json:"discord"`    // discord 账号
	Telegram  string `json:"telegram"`   // telegram 账号
	LevelName string `json:"level_name"` // 账号等级
}

type UserDetailResp struct {
	Email                string           `json:"email"`                  // 邮箱
	Discord              string           `json:"discord"`                // discord 账号
	Telegram             string           `json:"telegram"`               // telegram 账号
	LevelName            string           `json:"level_name"`             // 账号等级
	MembershipExpireTime int64            `json:"membership_expire_time"` // 会员过期时间
	LevelInfoData        []*LevelInfoData `json:"level_info_data"`        // 会员等级相关信息
}

type LevelInfoData struct {
	Grade     int8   `json:"grade"`      // 会员等级
	LevelName string `json:"level_name"` // 等级名称
	Type      int8   `json:"type"`       // 会员时效
}
