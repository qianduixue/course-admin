package response

import (
	"github.com/opisnoeasy/course-service/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
