package request

import (
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/course"
)

type CourseSectionSearch struct {
	course.CourseSection
	request.PageInfo
}
