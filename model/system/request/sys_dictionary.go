package request

import (
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
