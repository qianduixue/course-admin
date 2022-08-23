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

type CourseInfoApi struct {
}

var courseInfoService = service.ServiceGroupApp.CourseServiceGroup.CourseInfoService

// CreateCourseInfo 创建CourseInfo
// @Tags CourseInfo
// @Summary 创建CourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body course.CourseInfo true "创建CourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseInfo/createCourseInfo [post]
func (courseInfoApi *CourseInfoApi) CreateCourseInfo(c *gin.Context) {
	var courseInfo course.CourseInfo
	_ = c.ShouldBindJSON(&courseInfo)
	if err := utils.Verify(courseInfo, utils.CourseInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseInfoService.CreateCourseInfo(courseInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCourseInfo 删除CourseInfo
// @Tags CourseInfo
// @Summary 删除CourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body course.CourseInfo true "删除CourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseInfo/deleteCourseInfo [delete]
func (courseInfoApi *CourseInfoApi) DeleteCourseInfo(c *gin.Context) {
	var courseInfo course.CourseInfo
	_ = c.ShouldBindJSON(&courseInfo)
	if courseInfo.ID < 1 {
		response.FailWithMessage("删除课程id不能为空", c)
		return
	}
	if err := courseInfoService.DeleteCourseInfo(courseInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCourseInfoByIds 批量删除CourseInfo
// @Tags CourseInfo
// @Summary 批量删除CourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /courseInfo/deleteCourseInfoByIds [delete]
func (courseInfoApi *CourseInfoApi) DeleteCourseInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := courseInfoService.DeleteCourseInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCourseInfo 更新CourseInfo
// @Tags CourseInfo
// @Summary 更新CourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body course.CourseInfo true "更新CourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /courseInfo/updateCourseInfo [put]
func (courseInfoApi *CourseInfoApi) UpdateCourseInfo(c *gin.Context) {
	var courseInfo course.CourseInfo
	_ = c.ShouldBindJSON(&courseInfo)
	if courseInfo.ID < 1 {
		response.FailWithMessage("更新id不能为空", c)
		return
	}
	if err := utils.Verify(courseInfo, utils.CourseInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseInfoService.UpdateCourseInfo(courseInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCourseInfo 用id查询CourseInfo
// @Tags CourseInfo
// @Summary 用id查询CourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query course.CourseInfo true "用id查询CourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /courseInfo/findCourseInfo [get]
func (courseInfoApi *CourseInfoApi) FindCourseInfo(c *gin.Context) {
	var courseInfo course.CourseInfo
	_ = c.ShouldBindQuery(&courseInfo)
	if recourseInfo, err := courseInfoService.GetCourseInfo(courseInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recourseInfo": recourseInfo}, c)
	}
}

// GetCourseInfoList 分页获取CourseInfo列表
// @Tags CourseInfo
// @Summary 分页获取CourseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query courseReq.CourseInfoSearch true "分页获取CourseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseInfo/getCourseInfoList [get]
func (courseInfoApi *CourseInfoApi) GetCourseInfoList(c *gin.Context) {
	var pageInfo courseReq.CourseInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := courseInfoService.GetCourseInfoInfoList(pageInfo); err != nil {
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
