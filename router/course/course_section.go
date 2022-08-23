package course

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type CourseSectionRouter struct {
}

// InitCourseSectionRouter 初始化 CourseSection 路由信息
func (s *CourseSectionRouter) InitCourseSectionRouter(Router *gin.RouterGroup) {
	courseSectionRouter := Router.Group("courseSection").Use(middleware.OperationRecord())
	courseSectionRouterWithoutRecord := Router.Group("courseSection")
	var courseSectionApi = v1.ApiGroupApp.CourseApiGroup.CourseSectionApi
	{
		courseSectionRouter.POST("createCourseSection", courseSectionApi.CreateCourseSection)             // 新建CourseSection
		courseSectionRouter.DELETE("deleteCourseSection", courseSectionApi.DeleteCourseSection)           // 删除CourseSection
		courseSectionRouter.DELETE("deleteCourseSectionByIds", courseSectionApi.DeleteCourseSectionByIds) // 批量删除CourseSection
		courseSectionRouter.PUT("updateCourseSection", courseSectionApi.UpdateCourseSection)              // 更新CourseSection
	}
	{
		courseSectionRouterWithoutRecord.GET("findCourseSection", courseSectionApi.FindCourseSection)       // 根据ID获取CourseSection
		courseSectionRouterWithoutRecord.GET("getCourseSectionList", courseSectionApi.GetCourseSectionList) // 获取CourseSection列表
	}
}
