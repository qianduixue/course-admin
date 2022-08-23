package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type UserInfoRouter struct {
}

// InitUserInfoRouter 初始化 UserInfo 路由信息
func (s *UserInfoRouter) InitUserInfoRouter(Router *gin.RouterGroup) {
	userInfoRouter := Router.Group("userInfo").Use(middleware.OperationRecord())
	userInfoRouterWithoutRecord := Router.Group("userInfo")
	var userInfoApi = v1.ApiGroupApp.WebApiGroup.UserInfoApi
	{
		//userInfoRouter.POST("createUserInfo", userInfoApi.CreateUserInfo)   // 新建UserInfo
		//userInfoRouter.DELETE("deleteUserInfo", userInfoApi.DeleteUserInfo) // 删除UserInfo
		//userInfoRouter.DELETE("deleteUserInfoByIds", userInfoApi.DeleteUserInfoByIds) // 批量删除UserInfo
		userInfoRouter.PUT("updateUserInfo", userInfoApi.UpdateUserInfo) // 更新UserInfo
	}
	{
		userInfoRouterWithoutRecord.GET("findUserInfo", userInfoApi.FindUserInfo)       // 根据ID获取UserInfo
		userInfoRouterWithoutRecord.GET("getUserInfoList", userInfoApi.GetUserInfoList) // 获取UserInfo列表
	}
}
