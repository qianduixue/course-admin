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

type ContactUsApi struct {
}

var contactUsService = service.ServiceGroupApp.WebServiceGroup.ContactUsService

// CreateContactUs 创建ContactUs
// @Tags ContactUs
// @Summary 创建ContactUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.ContactUs true "创建ContactUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contactUs/createContactUs [post]
func (contactUsApi *ContactUsApi) CreateContactUs(c *gin.Context) {
	var contactUs web.ContactUs
	_ = c.ShouldBindJSON(&contactUs)
	if err := utils.Verify(contactUs, utils.ContactUsVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := contactUsService.CreateContactUs(contactUs); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteContactUs 删除ContactUs
// @Tags ContactUs
// @Summary 删除ContactUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.ContactUs true "删除ContactUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contactUs/deleteContactUs [delete]
func (contactUsApi *ContactUsApi) DeleteContactUs(c *gin.Context) {
	var contactUs web.ContactUs
	_ = c.ShouldBindJSON(&contactUs)
	if contactUs.ID < 1 {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if err := contactUsService.DeleteContactUs(contactUs); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteContactUsByIds 批量删除ContactUs
// @Tags ContactUs
// @Summary 批量删除ContactUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ContactUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /contactUs/deleteContactUsByIds [delete]
func (contactUsApi *ContactUsApi) DeleteContactUsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := contactUsService.DeleteContactUsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateContactUs 更新ContactUs
// @Tags ContactUs
// @Summary 更新ContactUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.ContactUs true "更新ContactUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contactUs/updateContactUs [put]
func (contactUsApi *ContactUsApi) UpdateContactUs(c *gin.Context) {
	var contactUs web.ContactUs
	_ = c.ShouldBindJSON(&contactUs)
	if contactUs.ID < 1 || contactUs.Icon == "" || contactUs.Account == "" || contactUs.AccountName == "" {
		response.FailWithMessage("参数不能为空", c)
		return
	}
	if err := contactUsService.UpdateContactUs(contactUs); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindContactUs 用id查询ContactUs
// @Tags ContactUs
// @Summary 用id查询ContactUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.ContactUs true "用id查询ContactUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contactUs/findContactUs [get]
func (contactUsApi *ContactUsApi) FindContactUs(c *gin.Context) {
	var contactUs web.ContactUs
	_ = c.ShouldBindQuery(&contactUs)
	if recontactUs, err := contactUsService.GetContactUs(contactUs.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recontactUs": recontactUs}, c)
	}
}

// GetContactUsList 分页获取ContactUs列表
// @Tags ContactUs
// @Summary 分页获取ContactUs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.ContactUsSearch true "分页获取ContactUs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contactUs/getContactUsList [get]
func (contactUsApi *ContactUsApi) GetContactUsList(c *gin.Context) {
	var pageInfo webReq.ContactUsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := contactUsService.GetContactUsInfoList(pageInfo); err != nil {
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
