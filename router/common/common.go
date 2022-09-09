package common

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type CommonRouter struct {
}

// InitCommonRouter 初始化 Common 路由信息
func (c *CommonRouter) InitCommonRouter(Router *gin.RouterGroup) {
	commonRouter := Router.Group("common").Use(middleware.OperationRecord())
	var commonApi = v1.ApiGroupApp.CommonApiGroup
	{
		commonRouter.POST("uploadFile", commonApi.Upload)                               // 文件上传
		commonRouter.POST("createMultipartUpload", commonApi.CreateMultipartUpload)     // 启动分段上传
		commonRouter.POST("uploadPart", commonApi.UploadPart)                           // 分段上传
		commonRouter.POST("completeMultipartUpload", commonApi.CompleteMultipartUpload) // 完成分段上传
		commonRouter.POST("listParts", commonApi.ListParts)                             // 分段列表
		commonRouter.POST("multipartUpload", commonApi.MultipartUpload)                 // 分段上传
	}
}
