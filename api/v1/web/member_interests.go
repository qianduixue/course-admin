package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/common/response"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
	"github.com/opisnoeasy/course-service/service"
	"github.com/opisnoeasy/course-service/utils"
	"go.uber.org/zap"
)

type MemberInterestsApi struct {
}

var memberInterestsService = service.ServiceGroupApp.WebServiceGroup.MemberInterestsService

// CreateMemberInterests 创建MemberInterests
// @Tags MemberInterests
// @Summary 创建MemberInterests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.MemberInterests true "创建MemberInterests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberInterests/createMemberInterests [post]
func (memberInterestsApi *MemberInterestsApi) CreateMemberInterests(c *gin.Context) {
	var memberInterests web.MemberInterests
	_ = c.ShouldBindJSON(&memberInterests)
	if err := utils.Verify(memberInterests, utils.MemberInterestsVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := memberInterestsService.CreateMemberInterests(memberInterests); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMemberInterests 删除MemberInterests
// @Tags MemberInterests
// @Summary 删除MemberInterests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.MemberInterests true "删除MemberInterests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memberInterests/deleteMemberInterests [delete]
func (memberInterestsApi *MemberInterestsApi) DeleteMemberInterests(c *gin.Context) {
	var memberInterests web.MemberInterests
	_ = c.ShouldBindJSON(&memberInterests)
	if memberInterests.ID < 1 {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if err := memberInterestsService.DeleteMemberInterests(memberInterests); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemberInterestsByIds 批量删除MemberInterests
// @Tags MemberInterests
// @Summary 批量删除MemberInterests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemberInterests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memberInterests/deleteMemberInterestsByIds [delete]
func (memberInterestsApi *MemberInterestsApi) DeleteMemberInterestsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := memberInterestsService.DeleteMemberInterestsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemberInterests 更新MemberInterests
// @Tags MemberInterests
// @Summary 更新MemberInterests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.MemberInterests true "更新MemberInterests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memberInterests/updateMemberInterests [put]
func (memberInterestsApi *MemberInterestsApi) UpdateMemberInterests(c *gin.Context) {
	var memberInterests web.MemberInterests
	_ = c.ShouldBindJSON(&memberInterests)
	if memberInterests.ID < 1 {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if err := utils.Verify(memberInterests, utils.MemberInterestsVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberInterestsService.UpdateMemberInterests(memberInterests); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemberInterests 用id查询MemberInterests
// @Tags MemberInterests
// @Summary 用id查询MemberInterests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.MemberInterests true "用id查询MemberInterests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memberInterests/findMemberInterests [get]
func (memberInterestsApi *MemberInterestsApi) FindMemberInterests(c *gin.Context) {
	var memberInterests web.MemberInterests
	_ = c.ShouldBindQuery(&memberInterests)
	if rememberInterests, err := memberInterestsService.GetMemberInterests(memberInterests.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rememberInterests": rememberInterests}, c)
	}
}

// GetMemberInterestsList 分页获取MemberInterests列表
// @Tags MemberInterests
// @Summary 分页获取MemberInterests列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.MemberInterestsSearch true "分页获取MemberInterests列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberInterests/getMemberInterestsList [get]
func (memberInterestsApi *MemberInterestsApi) GetMemberInterestsList(c *gin.Context) {
	var pageInfo webReq.MemberInterestsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := memberInterestsService.GetMemberInterestsInfoList(pageInfo); err != nil {
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
