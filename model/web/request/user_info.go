package request

import (
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/web"
)

type UserInfoSearch struct {
	LevelType int8 `json:"level_type" form:"level_type"` // 会员等级
	web.UserInfo
	request.PageInfo
}
