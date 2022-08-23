package service

import (
	"github.com/opisnoeasy/course-service/service/common"
	"github.com/opisnoeasy/course-service/service/course"
	"github.com/opisnoeasy/course-service/service/example"
	"github.com/opisnoeasy/course-service/service/system"
	"github.com/opisnoeasy/course-service/service/web"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CourseServiceGroup  course.ServiceGroup
	WebServiceGroup     web.ServiceGroup
	CommonServiceGroup  common.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
