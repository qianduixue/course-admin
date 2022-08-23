package web

import (
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
	"time"
)

type NotesInfoService struct {
}

// CreateNotesInfo 创建NotesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (notesInfoService *NotesInfoService) CreateNotesInfo(notesInfo web.NotesInfo) (err error) {
	nowTime := time.Now().Unix()
	notesInfo.CreatedTime = nowTime
	err = global.GVA_DB.Create(&notesInfo).Error
	return err
}

// DeleteNotesInfo 删除NotesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (notesInfoService *NotesInfoService) DeleteNotesInfo(notesInfo web.NotesInfo) (err error) {
	err = global.GVA_DB.Delete(&notesInfo).Error
	return err
}

// DeleteNotesInfoByIds 批量删除NotesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (notesInfoService *NotesInfoService) DeleteNotesInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.NotesInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateNotesInfo 更新NotesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (notesInfoService *NotesInfoService) UpdateNotesInfo(notesInfo web.NotesInfo) (err error) {
	nowTime := time.Now().Unix()
	updateInfo := map[string]interface{}{
		"updated_time": nowTime,
		"title":        notesInfo.Title,
		"content":      notesInfo.Content,
	}
	err = global.GVA_DB.Model(web.NotesInfo{}).Where("id = ?", notesInfo.ID).Updates(&updateInfo).Error
	return err
}

// GetNotesInfo 根据id获取NotesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (notesInfoService *NotesInfoService) GetNotesInfo(id uint) (notesInfo web.NotesInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&notesInfo).Error
	return
}

// GetNotesInfoInfoList 分页获取NotesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (notesInfoService *NotesInfoService) GetNotesInfoInfoList(info webReq.NotesInfoSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.NotesInfo{})
	var notesInfos []web.NotesInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+info.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&notesInfos).Error
	return notesInfos, total, err
}
