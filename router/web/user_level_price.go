package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type UserLevelPriceRouter struct {
}

// InitUserLevelPriceRouter 初始化 UserLevelPrice 路由信息
func (s *UserLevelPriceRouter) InitUserLevelPriceRouter(Router *gin.RouterGroup) {
	userLevelPriceRouter := Router.Group("userLevelPrice").Use(middleware.OperationRecord())
	userLevelPriceRouterWithoutRecord := Router.Group("userLevelPrice")
	var userLevelPriceApi = v1.ApiGroupApp.WebApiGroup.UserLevelPriceApi
	{
		userLevelPriceRouter.POST("createUserLevelPrice", userLevelPriceApi.CreateUserLevelPrice)             // 新建UserLevelPrice
		userLevelPriceRouter.DELETE("deleteUserLevelPrice", userLevelPriceApi.DeleteUserLevelPrice)           // 删除UserLevelPrice
		userLevelPriceRouter.DELETE("deleteUserLevelPriceByIds", userLevelPriceApi.DeleteUserLevelPriceByIds) // 批量删除UserLevelPrice
		userLevelPriceRouter.PUT("updateUserLevelPrice", userLevelPriceApi.UpdateUserLevelPrice)              // 更新UserLevelPrice
	}
	{
		userLevelPriceRouterWithoutRecord.GET("findUserLevelPrice", userLevelPriceApi.FindUserLevelPrice)       // 根据ID获取UserLevelPrice
		userLevelPriceRouterWithoutRecord.GET("getUserLevelPriceList", userLevelPriceApi.GetUserLevelPriceList) // 获取UserLevelPrice列表
	}
}
