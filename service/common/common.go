package common

import (
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/utils/upload"
	"mime/multipart"
)

type CommonService struct {
}

//UploadFileToLocal 上传文件到本地
func (commonService *CommonService) UploadFileToLocal(file *multipart.FileHeader) (path string, err error) {
	oss := upload.NewOss()
	uploadFile, _, err := oss.UploadFile(file)
	if err != nil {
		global.GVA_LOG.Fatal("上传文件失败")
		panic(err)
	}
	return uploadFile, err
}
