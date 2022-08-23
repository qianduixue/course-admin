package request

import (
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/course"
)

type CourseInfoSearch struct {
	course.CourseInfo
	request.PageInfo
}
