package request

import (
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/web"
)

type NotesInfoSearch struct {
	web.NotesInfo
	request.PageInfo
}
