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

type UserLevelPriceApi struct {
}

var userLevelPriceService = service.ServiceGroupApp.WebServiceGroup.UserLevelPriceService

// CreateUserLevelPrice 创建UserLevelPrice
// @Tags UserLevelPrice
// @Summary 创建UserLevelPrice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.UserLevelPrice true "创建UserLevelPrice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLevelPrice/createUserLevelPrice [post]
func (userLevelPriceApi *UserLevelPriceApi) CreateUserLevelPrice(c *gin.Context) {
	var userLevelPrice web.UserLevelPrice
	_ = c.ShouldBindJSON(&userLevelPrice)
	if err := utils.Verify(userLevelPrice, utils.UserLevelPriceVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userLevelPriceService.CreateUserLevelPrice(userLevelPrice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserLevelPrice 删除UserLevelPrice
// @Tags UserLevelPrice
// @Summary 删除UserLevelPrice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.UserLevelPrice true "删除UserLevelPrice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userLevelPrice/deleteUserLevelPrice [delete]
func (userLevelPriceApi *UserLevelPriceApi) DeleteUserLevelPrice(c *gin.Context) {
	var userLevelPrice web.UserLevelPrice
	_ = c.ShouldBindJSON(&userLevelPrice)
	if userLevelPrice.ID < 1 {
		response.FailWithMessage("id不能为空", c)
	}
	if err := userLevelPriceService.DeleteUserLevelPrice(userLevelPrice); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserLevelPriceByIds 批量删除UserLevelPrice
// @Tags UserLevelPrice
// @Summary 批量删除UserLevelPrice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserLevelPrice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userLevelPrice/deleteUserLevelPriceByIds [delete]
func (userLevelPriceApi *UserLevelPriceApi) DeleteUserLevelPriceByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := userLevelPriceService.DeleteUserLevelPriceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserLevelPrice 更新UserLevelPrice
// @Tags UserLevelPrice
// @Summary 更新UserLevelPrice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.UserLevelPrice true "更新UserLevelPrice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userLevelPrice/updateUserLevelPrice [put]
func (userLevelPriceApi *UserLevelPriceApi) UpdateUserLevelPrice(c *gin.Context) {
	var userLevelPrice web.UserLevelPrice
	_ = c.ShouldBindJSON(&userLevelPrice)
	if userLevelPrice.ID < 1 {
		response.FailWithMessage("id不能为空", c)
	}
	if err := utils.Verify(userLevelPrice, utils.UserLevelPriceVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := userLevelPriceService.UpdateUserLevelPrice(userLevelPrice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserLevelPrice 用id查询UserLevelPrice
// @Tags UserLevelPrice
// @Summary 用id查询UserLevelPrice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.UserLevelPrice true "用id查询UserLevelPrice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userLevelPrice/findUserLevelPrice [get]
func (userLevelPriceApi *UserLevelPriceApi) FindUserLevelPrice(c *gin.Context) {
	var userLevelPrice web.UserLevelPrice
	_ = c.ShouldBindQuery(&userLevelPrice)
	if reuserLevelPrice, err := userLevelPriceService.GetUserLevelPrice(userLevelPrice.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserLevelPrice": reuserLevelPrice}, c)
	}
}

// GetUserLevelPriceList 分页获取UserLevelPrice列表
// @Tags UserLevelPrice
// @Summary 分页获取UserLevelPrice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.UserLevelPriceSearch true "分页获取UserLevelPrice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLevelPrice/getUserLevelPriceList [get]
func (userLevelPriceApi *UserLevelPriceApi) GetUserLevelPriceList(c *gin.Context) {
	var pageInfo webReq.UserLevelPriceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := userLevelPriceService.GetUserLevelPriceInfoList(pageInfo); err != nil {
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
