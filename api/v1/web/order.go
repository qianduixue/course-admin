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

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.WebServiceGroup.OrderService

// CreateOrder 创建Order
// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var order web.Order
	_ = c.ShouldBindJSON(&order)
	if err := orderService.CreateOrder(order); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除Order
// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	var order web.Order
	_ = c.ShouldBindJSON(&order)
	if err := orderService.DeleteOrder(order); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除Order
// @Tags Order
// @Summary 批量删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /order/deleteOrderByIds [delete]
func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := orderService.DeleteOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新Order
// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var order web.Order
	_ = c.ShouldBindJSON(&order)
	if err := orderService.UpdateOrder(order); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询Order
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	var order web.Order
	_ = c.ShouldBindQuery(&order)
	if reorder, err := orderService.GetOrder(order.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// GetOrderList 分页获取Order列表
// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.OrderSearch true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo webReq.OrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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
