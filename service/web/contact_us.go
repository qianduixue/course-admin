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

type ContactUsService struct {
}

// CreateContactUs 创建ContactUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactUsService *ContactUsService) CreateContactUs(contactUs web.ContactUs) (err error) {
	nowTime := time.Now().Unix()
	contactUs.CreatedTime = nowTime
	err = global.GVA_DB.Create(&contactUs).Error
	if err != nil {
		return err
	}
	err = contactUsService.DelContactUsCache()
	return err
}

// DeleteContactUs 删除ContactUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactUsService *ContactUsService) DeleteContactUs(contactUs web.ContactUs) (err error) {
	err = global.GVA_DB.Delete(&contactUs).Error
	if err != nil {
		return err
	}
	err = contactUsService.DelContactUsCache()
	return err
}

// DeleteContactUsByIds 批量删除ContactUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactUsService *ContactUsService) DeleteContactUsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.ContactUs{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateContactUs 更新ContactUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactUsService *ContactUsService) UpdateContactUs(contactUs web.ContactUs) (err error) {
	updateInfo := map[string]interface{}{
		"updated_time": time.Now().Unix(),
		"account":      contactUs.Account,
		"account_name": contactUs.AccountName,
		"icon":         contactUs.Icon,
	}
	err = global.GVA_DB.Model(web.ContactUs{}).Where("id = ?", contactUs.ID).Updates(&updateInfo).Error
	if err != nil {
		return err
	}
	err = contactUsService.DelContactUsCache()
	return err
}

// GetContactUs 根据id获取ContactUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactUsService *ContactUsService) GetContactUs(id uint) (contactUs web.ContactUs, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&contactUs).Error
	return
}

// GetContactUsInfoList 分页获取ContactUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactUsService *ContactUsService) GetContactUsInfoList(info webReq.ContactUsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.ContactUs{})
	var contactUss []web.ContactUs
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&contactUss).Error
	return contactUss, total, err
}

//DelContactUsCache 删除缓存
func (contactUs *ContactUsService) DelContactUsCache() error {
	err := global.GVA_REDIS.Del(context.Background(), rk.ContactUs).Err()
	return err
}
