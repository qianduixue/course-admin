package resp

type UserLevelPriceResp struct {
	ID        uint   `json:"id"`         // 主键ID
	Name      string `json:"name"`       //课程标题
	Time      string `json:"time"`       //时间
	Price     int32  `json:"price"`      //所需价格
	LevelName string `json:"level_name"` //等级名称
}
