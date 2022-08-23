package web

import (
	"context"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/resp"
	"github.com/opisnoeasy/course-service/model/web"
	webReq "github.com/opisnoeasy/course-service/model/web/request"
	"github.com/opisnoeasy/course-service/rk"
	"strings"
	"time"
)

type UserInfoService struct {
}

// CreateUserInfo 创建UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) CreateUserInfo(userInfo web.UserInfo) (err error) {
	err = global.GVA_DB.Create(&userInfo).Error
	return err
}

// DeleteUserInfo 删除UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) DeleteUserInfo(userInfo web.UserInfo) (err error) {
	err = global.GVA_DB.Delete(&userInfo).Error
	return err
}

// DeleteUserInfoByIds 批量删除UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) DeleteUserInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.UserInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserInfo 更新UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) UpdateUserInfo(userInfo web.UserInfo) (err error) {
	updateInfo := map[string]interface{}{
		"updated_time":           time.Now().Unix(),
		"membership_expire_time": userInfo.MembershipExpireTime,
		"level":                  userInfo.Level,
	}
	err = global.GVA_DB.Model(web.UserInfo{}).Where("id = ?", userInfo.ID).Updates(&updateInfo).Error
	var build strings.Builder
	build.WriteString(rk.UserInfo)
	build.WriteString(userInfo.Uid)
	key := build.String()
	err = global.GVA_REDIS.Del(context.Background(), key).Err()
	return err
}

// GetUserInfo 根据id获取UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) GetUserInfo(id uint) (userDetail resp.UserDetailResp, err error) {
	var userInfo web.UserInfo
	err = global.GVA_DB.Where("id = ?", id).Preload("DetailData").First(&userInfo).Error
	if err != nil {
		return
	}
	levelName := "普通会员"
	if userInfo.Level > 0 {
		var levelInfo web.UserLevelPrice
		err = global.GVA_DB.Model(web.UserLevelPrice{}).Where("grade = ?", userInfo.Level).First(&levelInfo).Error
		if err != nil {
			return
		}
		levelName = levelInfo.Name
	}
	respData := resp.UserDetailResp{
		Email:                userInfo.Email,
		Discord:              userInfo.Discord,
		Telegram:             userInfo.DetailData.TelegramAccount,
		LevelName:            levelName,
		MembershipExpireTime: userInfo.MembershipExpireTime,
		LevelInfoData:        userInfoService.GetUserLevelInfo(),
	}
	return respData, nil
}

// GetUserInfoInfoList 分页获取UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) GetUserInfoInfoList(info webReq.UserInfoSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.UserInfo{})
	var userInfos []web.UserInfo
	if len(info.Keyword) > 0 {
		db = db.Where("email = ?", info.Keyword)
	}
	if info.LevelType > 0 {
		db = db.Where("level = ?", info.LevelType)
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//查询会员等级
	var levelInfos []web.UserLevelPrice
	err = global.GVA_DB.Model(web.UserLevelPrice{}).Find(&levelInfos).Error
	if err != nil {
		return
	}
	levelData := make(map[int8]string, 0)
	for i := 0; i < len(levelInfos); i++ {
		levelData[levelInfos[i].Grade] = levelInfos[i].Name
	}
	err = db.Limit(limit).Offset(offset).Preload("DetailData").Find(&userInfos).Error
	respData := make([]*resp.UserInfoResp, 0)
	for i := 0; i < len(userInfos); i++ {
		levelName := "普通用户"
		if userInfos[i].Level > 0 {
			levelName = levelData[userInfos[i].Level]
		}
		data := &resp.UserInfoResp{
			Id:        userInfos[i].ID,
			UID:       userInfos[i].Uid,
			Email:     userInfos[i].Email,
			Discord:   userInfos[i].Discord,
			Telegram:  userInfos[i].DetailData.TelegramAccount,
			LevelName: levelName,
		}
		respData = append(respData, data)
	}
	return respData, total, err
}

//GetUserLevelInfo 获取会员级别信息
func (userInfoService *UserInfoService) GetUserLevelInfo() []*resp.LevelInfoData {
	//查询会员等级
	var levelInfos []web.UserLevelPrice
	err := global.GVA_DB.Model(web.UserLevelPrice{}).Find(&levelInfos).Error
	if err != nil {
		return nil
	}
	levelData := make([]*resp.LevelInfoData, 0)
	for i := 0; i < len(levelInfos); i++ {
		data := &resp.LevelInfoData{
			Grade:     levelInfos[i].Grade,
			LevelName: levelInfos[i].Name,
			Type:      levelInfos[i].Type,
		}
		levelData = append(levelData, data)
	}
	levelData = append(levelData, &resp.LevelInfoData{
		Grade:     0,
		LevelName: "普通会员",
		Type:      0,
	})
	return levelData
}
