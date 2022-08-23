package v1

import (
	"github.com/opisnoeasy/course-service/api/v1/common"
	"github.com/opisnoeasy/course-service/api/v1/course"
	"github.com/opisnoeasy/course-service/api/v1/example"
	"github.com/opisnoeasy/course-service/api/v1/system"
	"github.com/opisnoeasy/course-service/api/v1/web"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CourseApiGroup  course.ApiGroup
	WebApiGroup     web.ApiGroup
	CommonApiGroup  common.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
