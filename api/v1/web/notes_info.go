package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/common/response"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
	"github.com/opisnoeasy/course-service/service"
	"go.uber.org/zap"
)

type NotesInfoApi struct {
}

var notesInfoService = service.ServiceGroupApp.WebServiceGroup.NotesInfoService

// CreateNotesInfo 创建NotesInfo
// @Tags NotesInfo
// @Summary 创建NotesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.NotesInfo true "创建NotesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notesInfo/createNotesInfo [post]
func (notesInfoApi *NotesInfoApi) CreateNotesInfo(c *gin.Context) {
	var notesInfo web.NotesInfo
	_ = c.ShouldBindJSON(&notesInfo)
	if notesInfo.Title == "" || notesInfo.Content == "" {
		response.FailWithMessage("必填参数不能为空", c)
		return
	}
	if err := notesInfoService.CreateNotesInfo(notesInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNotesInfo 删除NotesInfo
// @Tags NotesInfo
// @Summary 删除NotesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.NotesInfo true "删除NotesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notesInfo/deleteNotesInfo [delete]
func (notesInfoApi *NotesInfoApi) DeleteNotesInfo(c *gin.Context) {
	var notesInfo web.NotesInfo
	_ = c.ShouldBindJSON(&notesInfo)
	if notesInfo.ID < 1 {
		response.FailWithMessage("笔记id不能为空", c)
		return
	}
	if err := notesInfoService.DeleteNotesInfo(notesInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNotesInfoByIds 批量删除NotesInfo
// @Tags NotesInfo
// @Summary 批量删除NotesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NotesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /notesInfo/deleteNotesInfoByIds [delete]
func (notesInfoApi *NotesInfoApi) DeleteNotesInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := notesInfoService.DeleteNotesInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNotesInfo 更新NotesInfo
// @Tags NotesInfo
// @Summary 更新NotesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.NotesInfo true "更新NotesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notesInfo/updateNotesInfo [put]
func (notesInfoApi *NotesInfoApi) UpdateNotesInfo(c *gin.Context) {
	var notesInfo web.NotesInfo
	_ = c.ShouldBindJSON(&notesInfo)
	if notesInfo.ID < 1 || notesInfo.Title == "" || notesInfo.Content == "" {
		response.FailWithMessage("必填参数不能为空", c)
		return
	}
	if err := notesInfoService.UpdateNotesInfo(notesInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNotesInfo 用id查询NotesInfo
// @Tags NotesInfo
// @Summary 用id查询NotesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.NotesInfo true "用id查询NotesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /notesInfo/findNotesInfo [get]
func (notesInfoApi *NotesInfoApi) FindNotesInfo(c *gin.Context) {
	var notesInfo web.NotesInfo
	_ = c.ShouldBindQuery(&notesInfo)
	if renotesInfo, err := notesInfoService.GetNotesInfo(notesInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renotesInfo": renotesInfo}, c)
	}
}

// GetNotesInfoList 分页获取NotesInfo列表
// @Tags NotesInfo
// @Summary 分页获取NotesInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.NotesInfoSearch true "分页获取NotesInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notesInfo/getNotesInfoList [get]
func (notesInfoApi *NotesInfoApi) GetNotesInfoList(c *gin.Context) {
	var pageInfo webReq.NotesInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := notesInfoService.GetNotesInfoInfoList(pageInfo); err != nil {
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
