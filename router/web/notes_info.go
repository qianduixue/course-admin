package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type NotesInfoRouter struct {
}

// InitNotesInfoRouter 初始化 NotesInfo 路由信息
func (s *NotesInfoRouter) InitNotesInfoRouter(Router *gin.RouterGroup) {
	notesInfoRouter := Router.Group("notesInfo").Use(middleware.OperationRecord())
	notesInfoRouterWithoutRecord := Router.Group("notesInfo")
	var notesInfoApi = v1.ApiGroupApp.WebApiGroup.NotesInfoApi
	{
		notesInfoRouter.POST("createNotesInfo", notesInfoApi.CreateNotesInfo)             // 新建NotesInfo
		notesInfoRouter.DELETE("deleteNotesInfo", notesInfoApi.DeleteNotesInfo)           // 删除NotesInfo
		notesInfoRouter.DELETE("deleteNotesInfoByIds", notesInfoApi.DeleteNotesInfoByIds) // 批量删除NotesInfo
		notesInfoRouter.PUT("updateNotesInfo", notesInfoApi.UpdateNotesInfo)              // 更新NotesInfo
	}
	{
		notesInfoRouterWithoutRecord.GET("findNotesInfo", notesInfoApi.FindNotesInfo)       // 根据ID获取NotesInfo
		notesInfoRouterWithoutRecord.GET("getNotesInfoList", notesInfoApi.GetNotesInfoList) // 获取NotesInfo列表
	}
}
