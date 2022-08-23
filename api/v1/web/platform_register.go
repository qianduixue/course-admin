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

type PlatformRegisterApi struct {
}

var platformRegisterService = service.ServiceGroupApp.WebServiceGroup.PlatformRegisterService

// CreatePlatformRegister 创建PlatformRegister
// @Tags PlatformRegister
// @Summary 创建PlatformRegister
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.PlatformRegister true "创建PlatformRegister"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platformRegister/createPlatformRegister [post]
func (platformRegisterApi *PlatformRegisterApi) CreatePlatformRegister(c *gin.Context) {
	var platformRegister web.PlatformRegister
	_ = c.ShouldBindJSON(&platformRegister)
	if err := utils.Verify(platformRegister, utils.PlatformRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := platformRegisterService.CreatePlatformRegister(platformRegister); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlatformRegister 删除PlatformRegister
// @Tags PlatformRegister
// @Summary 删除PlatformRegister
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.PlatformRegister true "删除PlatformRegister"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platformRegister/deletePlatformRegister [delete]
func (platformRegisterApi *PlatformRegisterApi) DeletePlatformRegister(c *gin.Context) {
	var platformRegister web.PlatformRegister
	_ = c.ShouldBindJSON(&platformRegister)
	if platformRegister.ID < 1 {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if err := platformRegisterService.DeletePlatformRegister(platformRegister); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlatformRegisterByIds 批量删除PlatformRegister
// @Tags PlatformRegister
// @Summary 批量删除PlatformRegister
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlatformRegister"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /platformRegister/deletePlatformRegisterByIds [delete]
func (platformRegisterApi *PlatformRegisterApi) DeletePlatformRegisterByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := platformRegisterService.DeletePlatformRegisterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlatformRegister 更新PlatformRegister
// @Tags PlatformRegister
// @Summary 更新PlatformRegister
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.PlatformRegister true "更新PlatformRegister"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platformRegister/updatePlatformRegister [put]
func (platformRegisterApi *PlatformRegisterApi) UpdatePlatformRegister(c *gin.Context) {
	var platformRegister web.PlatformRegister
	_ = c.ShouldBindJSON(&platformRegister)
	if platformRegister.ID < 1 || platformRegister.BackgroundImage == "" || platformRegister.Linked == "" {
		response.FailWithMessage("必填参数不能为空", c)
		return
	}
	if err := platformRegisterService.UpdatePlatformRegister(platformRegister); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlatformRegister 用id查询PlatformRegister
// @Tags PlatformRegister
// @Summary 用id查询PlatformRegister
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.PlatformRegister true "用id查询PlatformRegister"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platformRegister/findPlatformRegister [get]
func (platformRegisterApi *PlatformRegisterApi) FindPlatformRegister(c *gin.Context) {
	var platformRegister web.PlatformRegister
	_ = c.ShouldBindQuery(&platformRegister)
	if replatformRegister, err := platformRegisterService.GetPlatformRegister(platformRegister.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replatformRegister": replatformRegister}, c)
	}
}

// GetPlatformRegisterList 分页获取PlatformRegister列表
// @Tags PlatformRegister
// @Summary 分页获取PlatformRegister列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.PlatformRegisterSearch true "分页获取PlatformRegister列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platformRegister/getPlatformRegisterList [get]
func (platformRegisterApi *PlatformRegisterApi) GetPlatformRegisterList(c *gin.Context) {
	var pageInfo webReq.PlatformRegisterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := platformRegisterService.GetPlatformRegisterInfoList(pageInfo); err != nil {
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
