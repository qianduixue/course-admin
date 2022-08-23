package web

import (
	"github.com/gin-gonic/gin"
	"github.com/opisnoeasy/course-service/api/v1"
	"github.com/opisnoeasy/course-service/middleware"
)

type OrderRouter struct {
}

// InitOrderRouter 初始化 Order 路由信息
func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouter := Router.Group("order").Use(middleware.OperationRecord())
	orderRouterWithoutRecord := Router.Group("order")
	var orderApi = v1.ApiGroupApp.WebApiGroup.OrderApi
	{
		//	orderRouter.POST("createOrder", orderApi.CreateOrder)   // 新建Order
		//	orderRouter.DELETE("deleteOrder", orderApi.DeleteOrder) // 删除Order
		//	orderRouter.DELETE("deleteOrderByIds", orderApi.DeleteOrderByIds) // 批量删除Order
		orderRouter.PUT("updateOrder", orderApi.UpdateOrder) // 更新Order
	}
	{
		//orderRouterWithoutRecord.GET("findOrder", orderApi.FindOrder)        // 根据ID获取Order
		orderRouterWithoutRecord.GET("getOrderList", orderApi.GetOrderList) // 获取Order列表
	}
}
