package web

import (
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/resp"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
)

type OrderService struct {
}

// CreateOrder 创建Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) CreateOrder(order web.Order) (err error) {
	err = global.GVA_DB.Create(&order).Error
	return err
}

// DeleteOrder 删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrder(order web.Order) (err error) {
	err = global.GVA_DB.Delete(&order).Error
	return err
}

// DeleteOrderByIds 批量删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.Order{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrder 更新Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) UpdateOrder(order web.Order) (err error) {
	err = global.GVA_DB.Save(&order).Error
	return err
}

// GetOrder 根据id获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrder(id uint) (order web.Order, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	return
}

// GetOrderInfoList 分页获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrderInfoList(info webReq.OrderSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.Order{})
	var orders []web.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.Keyword) > 0 {
		db = db.Where("email = ?", info.Keyword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("LevelData").Find(&orders).Error
	respData := make([]*resp.OrderResp, 0)
	for i := 0; i < len(orders); i++ {
		data := &resp.OrderResp{
			Id:          orders[i].ID,
			Email:       orders[i].Email,
			Discord:     orders[i].Discord,
			Telegram:    orders[i].Telegram,
			OrderSn:     orders[i].OrderSn,
			Price:       orders[i].Price,
			PayVoucher:  orders[i].PayVoucher,
			LevelName:   orders[i].LevelData.Name,
			CreatedTime: orders[i].CreatedTime,
		}
		respData = append(respData, data)
	}
	return respData, total, err
}
