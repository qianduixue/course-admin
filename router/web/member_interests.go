package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type MemberInterestsRouter struct {
}

// InitMemberInterestsRouter 初始化 MemberInterests 路由信息
func (s *MemberInterestsRouter) InitMemberInterestsRouter(Router *gin.RouterGroup) {
	memberInterestsRouter := Router.Group("memberInterests").Use(middleware.OperationRecord())
	memberInterestsRouterWithoutRecord := Router.Group("memberInterests")
	var memberInterestsApi = v1.ApiGroupApp.WebApiGroup.MemberInterestsApi
	{
		memberInterestsRouter.POST("createMemberInterests", memberInterestsApi.CreateMemberInterests)             // 新建MemberInterests
		memberInterestsRouter.DELETE("deleteMemberInterests", memberInterestsApi.DeleteMemberInterests)           // 删除MemberInterests
		memberInterestsRouter.DELETE("deleteMemberInterestsByIds", memberInterestsApi.DeleteMemberInterestsByIds) // 批量删除MemberInterests
		memberInterestsRouter.PUT("updateMemberInterests", memberInterestsApi.UpdateMemberInterests)              // 更新MemberInterests
	}
	{
		memberInterestsRouterWithoutRecord.GET("findMemberInterests", memberInterestsApi.FindMemberInterests)       // 根据ID获取MemberInterests
		memberInterestsRouterWithoutRecord.GET("getMemberInterestsList", memberInterestsApi.GetMemberInterestsList) // 获取MemberInterests列表
	}
}
