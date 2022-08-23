package request

import (
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/web"
)

type ContactUsSearch struct {
	web.ContactUs
	request.PageInfo
}
