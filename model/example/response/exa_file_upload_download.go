package response

import "github.com/opisnoeasy/course-service/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
