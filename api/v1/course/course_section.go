package course

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/common/response"
	"github.com/opisnoeasy/course-service/model/course"
	courseReq "github.com/opisnoeasy/course-service/model/course/request"
	"github.com/opisnoeasy/course-service/service"
	"github.com/opisnoeasy/course-service/utils"
	"go.uber.org/zap"
)

type CourseSectionApi struct {
}

var courseSectionService = service.ServiceGroupApp.CourseServiceGroup.CourseSectionService

// CreateCourseSection 创建CourseSection
// @Tags CourseSection
// @Summary 创建CourseSection
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body course.CourseSection true "创建CourseSection"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseSection/createCourseSection [post]
func (courseSectionApi *CourseSectionApi) CreateCourseSection(c *gin.Context) {
	var courseSection course.CourseSection
	_ = c.ShouldBindJSON(&courseSection)
	if err := utils.Verify(courseSection, utils.CourseInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := courseSectionService.CreateCourseSection(courseSection); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCourseSection 删除CourseSection
// @Tags CourseSection
// @Summary 删除CourseSection
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body course.CourseSection true "删除CourseSection"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseSection/deleteCourseSection [delete]
func (courseSectionApi *CourseSectionApi) DeleteCourseSection(c *gin.Context) {
	var courseSection course.CourseSection
	_ = c.ShouldBindJSON(&courseSection)
	if courseSection.ID < 1 {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if err := courseSectionService.DeleteCourseSection(courseSection); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCourseSectionByIds 批量删除CourseSection
// @Tags CourseSection
// @Summary 批量删除CourseSection
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CourseSection"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /courseSection/deleteCourseSectionByIds [delete]
func (courseSectionApi *CourseSectionApi) DeleteCourseSectionByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := courseSectionService.DeleteCourseSectionByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCourseSection 更新CourseSection
// @Tags CourseSection
// @Summary 更新CourseSection
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body course.CourseSection true "更新CourseSection"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /courseSection/updateCourseSection [put]
func (courseSectionApi *CourseSectionApi) UpdateCourseSection(c *gin.Context) {
	var courseSection course.CourseSection
	_ = c.ShouldBindJSON(&courseSection)
	if courseSection.ID < 1 || courseSection.Title == "" || courseSection.Lessons == "" || courseSection.VideoUrl == "" || courseSection.CoverImage == "" {
		response.FailWithMessage("必填参数不能为空", c)
		return
	}
	if err := courseSectionService.UpdateCourseSection(courseSection); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCourseSection 用id查询CourseSection
// @Tags CourseSection
// @Summary 用id查询CourseSection
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query course.CourseSection true "用id查询CourseSection"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /courseSection/findCourseSection [get]
func (courseSectionApi *CourseSectionApi) FindCourseSection(c *gin.Context) {
	var courseSection course.CourseSection
	_ = c.ShouldBindQuery(&courseSection)
	if recourseSection, err := courseSectionService.GetCourseSection(courseSection.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recourseSection": recourseSection}, c)
	}
}

// GetCourseSectionList 分页获取CourseSection列表
// @Tags CourseSection
// @Summary 分页获取CourseSection列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query courseReq.CourseSectionSearch true "分页获取CourseSection列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseSection/getCourseSectionList [get]
func (courseSectionApi *CourseSectionApi) GetCourseSectionList(c *gin.Context) {
	var pageInfo courseReq.CourseSectionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := courseSectionService.GetCourseSectionInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
