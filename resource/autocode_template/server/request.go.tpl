package request

import (
	"github.com/opisnoeasy/course-service/model/{{.Package}}"
	"github.com/opisnoeasy/course-service/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
