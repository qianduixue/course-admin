package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type ContactUsRouter struct {
}

// InitContactUsRouter 初始化 ContactUs 路由信息
func (s *ContactUsRouter) InitContactUsRouter(Router *gin.RouterGroup) {
	contactUsRouter := Router.Group("contactUs").Use(middleware.OperationRecord())
	contactUsRouterWithoutRecord := Router.Group("contactUs")
	var contactUsApi = v1.ApiGroupApp.WebApiGroup.ContactUsApi
	{
		contactUsRouter.POST("createContactUs", contactUsApi.CreateContactUs)             // 新建ContactUs
		contactUsRouter.DELETE("deleteContactUs", contactUsApi.DeleteContactUs)           // 删除ContactUs
		contactUsRouter.DELETE("deleteContactUsByIds", contactUsApi.DeleteContactUsByIds) // 批量删除ContactUs
		contactUsRouter.PUT("updateContactUs", contactUsApi.UpdateContactUs)              // 更新ContactUs
	}
	{
		contactUsRouterWithoutRecord.GET("findContactUs", contactUsApi.FindContactUs)       // 根据ID获取ContactUs
		contactUsRouterWithoutRecord.GET("getContactUsList", contactUsApi.GetContactUsList) // 获取ContactUs列表
	}
}
