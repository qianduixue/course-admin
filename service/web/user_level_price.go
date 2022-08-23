package web

import (
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/resp"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
	"time"
)

type UserLevelPriceService struct {
}

// CreateUserLevelPrice 创建UserLevelPrice记录
// Author [piexlmax](https://github.com/piexlmax)
func (userLevelPriceService *UserLevelPriceService) CreateUserLevelPrice(userLevelPrice web.UserLevelPrice) (err error) {
	nowTime := time.Now().Unix()
	userLevelPrice.CreatedTime = nowTime
	err = global.GVA_DB.Create(&userLevelPrice).Error
	return err
}

// DeleteUserLevelPrice 删除UserLevelPrice记录
// Author [piexlmax](https://github.com/piexlmax)
func (userLevelPriceService *UserLevelPriceService) DeleteUserLevelPrice(userLevelPrice web.UserLevelPrice) (err error) {
	err = global.GVA_DB.Delete(&userLevelPrice).Error
	return err
}

// DeleteUserLevelPriceByIds 批量删除UserLevelPrice记录
// Author [piexlmax](https://github.com/piexlmax)
func (userLevelPriceService *UserLevelPriceService) DeleteUserLevelPriceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.UserLevelPrice{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserLevelPrice 更新UserLevelPrice记录
// Author [piexlmax](https://github.com/piexlmax)
func (userLevelPriceService *UserLevelPriceService) UpdateUserLevelPrice(userLevelPrice web.UserLevelPrice) (err error) {
	updateInfo := map[string]interface{}{
		"updated_time": time.Now().Unix(),
		"type":         userLevelPrice.Type,
		"grade":        userLevelPrice.Grade,
		"price":        userLevelPrice.Price,
		"name":         userLevelPrice.Name,
	}
	err = global.GVA_DB.Model(web.UserLevelPrice{}).Where("id = ?", userLevelPrice.ID).Save(&updateInfo).Error
	return err
}

// GetUserLevelPrice 根据id获取UserLevelPrice记录
// Author [piexlmax](https://github.com/piexlmax)
func (userLevelPriceService *UserLevelPriceService) GetUserLevelPrice(id uint) (userLevelPrice web.UserLevelPrice, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userLevelPrice).Error
	return
}

// GetUserLevelPriceInfoList 分页获取UserLevelPrice记录
// Author [piexlmax](https://github.com/piexlmax)
func (userLevelPriceService *UserLevelPriceService) GetUserLevelPriceInfoList(info webReq.UserLevelPriceSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.UserLevelPrice{})
	var userLevelPrices []web.UserLevelPrice
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	levelTime := map[int8]string{
		1: "每月",
		2: "每季",
		3: "每年",
	}
	levelData := map[int8]string{
		1: "一级",
		2: "二级",
		3: "三级",
	}
	err = db.Limit(limit).Offset(offset).Find(&userLevelPrices).Error
	respData := make([]*resp.UserLevelPriceResp, 0)
	for i := 0; i < len(userLevelPrices); i++ {
		data := &resp.UserLevelPriceResp{
			ID:        userLevelPrices[i].ID,
			Name:      userLevelPrices[i].Name,
			Time:      levelTime[userLevelPrices[i].Type],
			Price:     userLevelPrices[i].Price,
			LevelName: levelData[userLevelPrices[i].Grade],
		}
		respData = append(respData, data)
	}
	return respData, total, err
}
