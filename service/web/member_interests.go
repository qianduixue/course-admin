package web

import (
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
	"time"
)

type MemberInterestsService struct {
}

// CreateMemberInterests 创建MemberInterests记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberInterestsService *MemberInterestsService) CreateMemberInterests(memberInterests web.MemberInterests) (err error) {
	nowTime := time.Now().Unix()
	memberInterests.CreatedTime = nowTime
	err = global.GVA_DB.Create(&memberInterests).Error
	return err
}

// DeleteMemberInterests 删除MemberInterests记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberInterestsService *MemberInterestsService) DeleteMemberInterests(memberInterests web.MemberInterests) (err error) {
	err = global.GVA_DB.Delete(&memberInterests).Error
	return err
}

// DeleteMemberInterestsByIds 批量删除MemberInterests记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberInterestsService *MemberInterestsService) DeleteMemberInterestsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.MemberInterests{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMemberInterests 更新MemberInterests记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberInterestsService *MemberInterestsService) UpdateMemberInterests(memberInterests web.MemberInterests) (err error) {
	updateInfo := map[string]interface{}{
		"updated_time": time.Now().Unix(),
		"title":        memberInterests.Title,
		"linked":       memberInterests.Linked,
	}
	err = global.GVA_DB.Model(web.MemberInterests{}).Where("id = ?", memberInterests.ID).Updates(&updateInfo).Error
	return err
}

// GetMemberInterests 根据id获取MemberInterests记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberInterestsService *MemberInterestsService) GetMemberInterests(id uint) (memberInterests web.MemberInterests, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&memberInterests).Error
	return
}

// GetMemberInterestsInfoList 分页获取MemberInterests记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberInterestsService *MemberInterestsService) GetMemberInterestsInfoList(info webReq.MemberInterestsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.MemberInterests{})
	var memberInterestss []web.MemberInterests
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.PageInfo.Keyword) > 0 {
		db.Where("title = ?", info.Title)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Select("id,title,linked").Find(&memberInterestss).Error
	return memberInterestss, total, err
}
