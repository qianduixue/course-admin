package common

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/response"
	"github.com/opisnoeasy/course-service/service"
	"github.com/opisnoeasy/course-service/utils"
	"go.uber.org/zap"
	"strconv"
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

func (commonApi *CommonApi) CreateMultipartUpload(c *gin.Context) {
	path := c.Request.PostFormValue("file")
	if path == "" {
		response.FailWithMessage("路径不能为空", c)
		return
	}
	result, err := utils.AloneCreateMultipartUpload(path)
	if err != nil {
		global.GVA_LOG.Error("启动分段上传失败!", zap.Error(err))
		response.FailWithMessage("启动分段上传失败", c)
	} else {
		response.OkWithData(gin.H{"result": result}, c)
	}
}

func (commonApi *CommonApi) UploadPart(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	key := c.Request.FormValue("key")
	uploadId := c.Request.FormValue("uploadId")
	partNumber, _ := strconv.Atoi(c.Request.FormValue("partNumber"))
	partSize, _ := strconv.Atoi(c.Request.FormValue("partSize"))
	if key == "" || uploadId == "" || partNumber < 0 {
		response.FailWithMessage("参数不能为空", c)
		return
	}

	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	f, err := header.Open()
	if err != nil {
		global.GVA_LOG.Error("打开文件失败!", zap.Error(err))
		response.FailWithMessage("打开文件失败", c)
		return
	}
	defer f.Close()

	data := make([]byte, partSize)
	read, _ := f.Read(data)
	if read == 0 {
		global.GVA_LOG.Error("文件读取部分失败!", zap.Error(err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	result, err := utils.AloneUploadPart(uploadId, key, int64(partNumber), data)
	if err != nil {
		global.GVA_LOG.Error("分段上传失败!", zap.Error(err))
		response.FailWithMessage("分段上传失败", c)
	} else {
		response.OkWithData(gin.H{"result": result}, c)
	}
}

func (commonApi *CommonApi) CompleteMultipartUpload(c *gin.Context) {
	parts := c.Request.PostFormValue("parts")
	uploadId := c.Request.PostFormValue("uploadId")
	key := c.Request.PostFormValue("key")
	if parts == "" || uploadId == "" || key == "" {
		response.FailWithMessage("参数不能为空", c)
		return
	}
	var completePart []*s3.CompletedPart
	err := json.Unmarshal([]byte(parts), &completePart)
	if err != nil {
		response.FailWithMessage("json解析失败", c)
		return
	}
	result, err := utils.AloneCompleteMultipartUpload(key, uploadId, completePart)
	if err != nil {
		global.GVA_LOG.Error("完成分段上传失败!", zap.Error(err))
		response.FailWithMessage("完成分段上传失败", c)
	} else {
		response.OkWithData(gin.H{"result": result}, c)
	}
}

//MultipartUpload 启动分段上传
func (commonApi *CommonApi) MultipartUpload(c *gin.Context) {
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
