package request

import (
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/web"
)

type UserLevelPriceSearch struct {
	web.UserLevelPrice
	request.PageInfo
}
