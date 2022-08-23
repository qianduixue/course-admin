package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type PlatformRegisterRouter struct {
}

// InitPlatformRegisterRouter 初始化 PlatformRegister 路由信息
func (s *PlatformRegisterRouter) InitPlatformRegisterRouter(Router *gin.RouterGroup) {
	platformRegisterRouter := Router.Group("platformRegister").Use(middleware.OperationRecord())
	platformRegisterRouterWithoutRecord := Router.Group("platformRegister")
	var platformRegisterApi = v1.ApiGroupApp.WebApiGroup.PlatformRegisterApi
	{
		platformRegisterRouter.POST("createPlatformRegister", platformRegisterApi.CreatePlatformRegister)             // 新建PlatformRegister
		platformRegisterRouter.DELETE("deletePlatformRegister", platformRegisterApi.DeletePlatformRegister)           // 删除PlatformRegister
		platformRegisterRouter.DELETE("deletePlatformRegisterByIds", platformRegisterApi.DeletePlatformRegisterByIds) // 批量删除PlatformRegister
		platformRegisterRouter.PUT("updatePlatformRegister", platformRegisterApi.UpdatePlatformRegister)              // 更新PlatformRegister
	}
	{
		platformRegisterRouterWithoutRecord.GET("findPlatformRegister", platformRegisterApi.FindPlatformRegister)       // 根据ID获取PlatformRegister
		platformRegisterRouterWithoutRecord.GET("getPlatformRegisterList", platformRegisterApi.GetPlatformRegisterList) // 获取PlatformRegister列表
	}
}
