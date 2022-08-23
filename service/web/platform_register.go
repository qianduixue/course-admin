package web

import (
	"context"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
	"github.com/opisnoeasy/course-service/rk"
	"time"
)

type PlatformRegisterService struct {
}

// CreatePlatformRegister 创建PlatformRegister记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformRegisterService *PlatformRegisterService) CreatePlatformRegister(platformRegister web.PlatformRegister) (err error) {
	nowTime := time.Now().Unix()
	platformRegister.CreatedTime = nowTime
	err = global.GVA_DB.Create(&platformRegister).Error
	if err != nil {
		return err
	}
	err = platformRegisterService.DelPlatformRegisterCache()
	return err
}

// DeletePlatformRegister 删除PlatformRegister记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformRegisterService *PlatformRegisterService) DeletePlatformRegister(platformRegister web.PlatformRegister) (err error) {
	err = global.GVA_DB.Delete(&platformRegister).Error
	if err != nil {
		return err
	}
	err = platformRegisterService.DelPlatformRegisterCache()
	return err
}

// DeletePlatformRegisterByIds 批量删除PlatformRegister记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformRegisterService *PlatformRegisterService) DeletePlatformRegisterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.PlatformRegister{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePlatformRegister 更新PlatformRegister记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformRegisterService *PlatformRegisterService) UpdatePlatformRegister(platformRegister web.PlatformRegister) (err error) {
	updateInfo := map[string]interface{}{
		"updated_time":     time.Now().Unix(),
		"background_image": platformRegister.BackgroundImage,
		"linked":           platformRegister.Linked,
	}
	err = global.GVA_DB.Model(web.PlatformRegister{}).Where("id = ?", platformRegister.ID).Updates(&updateInfo).Error
	if err != nil {
		return err
	}
	err = platformRegisterService.DelPlatformRegisterCache()
	return err
}

// GetPlatformRegister 根据id获取PlatformRegister记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformRegisterService *PlatformRegisterService) GetPlatformRegister(id uint) (platformRegister web.PlatformRegister, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&platformRegister).Error
	return
}

// GetPlatformRegisterInfoList 分页获取PlatformRegister记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformRegisterService *PlatformRegisterService) GetPlatformRegisterInfoList(info webReq.PlatformRegisterSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.PlatformRegister{})
	var platformRegisters []web.PlatformRegister
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&platformRegisters).Error
	return platformRegisters, total, err
}

func (platformRegister *PlatformRegisterService) DelPlatformRegisterCache() error {
	err := global.GVA_REDIS.Del(context.Background(), rk.PlatformReg).Err()
	return err
}
