package response

import "github.com/opisnoeasy/course-service/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
