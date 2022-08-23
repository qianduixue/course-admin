package common

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/response"
	"github.com/opisnoeasy/course-service/service"
	"go.uber.org/zap"
)

type CommonApi struct {
}

var commonService = service.ServiceGroupApp.CommonServiceGroup.CommonService

//Upload 上传文件
func (commonApi *CommonApi) Upload(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	uploadPath, err := commonService.UploadFileToLocal(header)
	if err != nil {
		global.GVA_LOG.Error("上传文件失败!", zap.Error(err))
		response.FailWithMessage("上传文件失败", c)
	} else {
		response.OkWithData(gin.H{"filepath": uploadPath}, c)
	}
}
