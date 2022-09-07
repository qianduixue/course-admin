package common

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/response"
	"github.com/opisnoeasy/course-service/service"
	"github.com/opisnoeasy/course-service/utils"
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
	uploadPath, err := commonService.UploadFileToAws(header)
	if err != nil {
		global.GVA_LOG.Error("上传文件失败!", zap.Error(err))
		response.FailWithMessage("上传文件失败", c)
	} else {
		response.OkWithData(gin.H{"filepath": uploadPath}, c)
	}
}

//CreateMultipartUpload 启动分段上传
func (commonApi *CommonApi) CreateMultipartUpload(c *gin.Context) {
	filePath := c.Request.PostFormValue("file")
	if filePath == "" {
		response.FailWithMessage("路径不能为空", c)
		return
	}
	path, err := utils.AwsMultipartUpload(filePath)
	if err != nil {
		global.GVA_LOG.Error("创建分段上传失败!", zap.Error(err))
		response.FailWithMessage("启动分段上传失败", c)
	} else {
		response.OkWithData(gin.H{"filepath": path}, c)
	}
}

//ListParts 上传分段列表
func (commonApi *CommonApi) ListParts(c *gin.Context) {
	key := c.Request.PostFormValue("key")
	uploadId := c.Request.PostFormValue("uploadId")
	result, err := utils.ListParts(key, uploadId)
	if err != nil {
		global.GVA_LOG.Error("获取分段列表失败!", zap.Error(err))
		response.FailWithMessage("获取分段列表失败", c)
	} else {
		response.OkWithData(gin.H{"result": result}, c)
	}
}
