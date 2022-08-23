package course

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type CourseInfoRouter struct {
}

// InitCourseInfoRouter 初始化 CourseInfo 路由信息
func (s *CourseInfoRouter) InitCourseInfoRouter(Router *gin.RouterGroup) {
	courseInfoRouter := Router.Group("courseInfo").Use(middleware.OperationRecord())
	courseInfoRouterWithoutRecord := Router.Group("courseInfo")
	var courseInfoApi = v1.ApiGroupApp.CourseApiGroup.CourseInfoApi
	{
		courseInfoRouter.POST("createCourseInfo", courseInfoApi.CreateCourseInfo)             // 新建CourseInfo
		courseInfoRouter.DELETE("deleteCourseInfo", courseInfoApi.DeleteCourseInfo)           // 删除CourseInfo
		courseInfoRouter.DELETE("deleteCourseInfoByIds", courseInfoApi.DeleteCourseInfoByIds) // 批量删除CourseInfo
		courseInfoRouter.PUT("updateCourseInfo", courseInfoApi.UpdateCourseInfo)              // 更新CourseInfo
	}
	{
		courseInfoRouterWithoutRecord.GET("findCourseInfo", courseInfoApi.FindCourseInfo)       // 根据ID获取CourseInfo
		courseInfoRouterWithoutRecord.GET("getCourseInfoList", courseInfoApi.GetCourseInfoList) // 获取CourseInfo列表
	}
}
