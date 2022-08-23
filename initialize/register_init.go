package initialize

import (
	_ "github.com/opisnoeasy/course-service/source/example"
	_ "github.com/opisnoeasy/course-service/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
