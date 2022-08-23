package router

import (
	"github.com/opisnoeasy/course-service/router/common"
	"github.com/opisnoeasy/course-service/router/course"
	"github.com/opisnoeasy/course-service/router/example"
	"github.com/opisnoeasy/course-service/router/system"
	"github.com/opisnoeasy/course-service/router/web"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Course  course.RouterGroup
	Web     web.RouterGroup
	Common  common.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
