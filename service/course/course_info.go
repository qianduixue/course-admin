package course

import (
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/course"
	courseReq "github.com/opisnoeasy/course-service/model/course/request"
	"github.com/opisnoeasy/course-service/model/resp"
	"github.com/opisnoeasy/course-service/model/web"
	"time"
)

type CourseInfoService struct {
}

// CreateCourseInfo 创建CourseInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseInfoService *CourseInfoService) CreateCourseInfo(courseInfo course.CourseInfo) (err error) {
	nowTime := time.Now().Unix()
	courseInfo.CreatedTime = nowTime
	err = global.GVA_DB.Create(&courseInfo).Error
	return err
}

// DeleteCourseInfo 删除CourseInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseInfoService *CourseInfoService) DeleteCourseInfo(courseInfo course.CourseInfo) (err error) {
	updateInfo := map[string]interface{}{
		"status":       2,
		"updated_time": time.Now().Unix(),
	}
	err = global.GVA_DB.Model(course.CourseInfo{}).Where("id = ?", courseInfo.ID).Updates(&updateInfo).Error
	return err
}

// DeleteCourseInfoByIds 批量删除CourseInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseInfoService *CourseInfoService) DeleteCourseInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]course.CourseInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCourseInfo 更新CourseInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseInfoService *CourseInfoService) UpdateCourseInfo(courseInfo course.CourseInfo) (err error) {
	nowTime := time.Now().Unix()
	updateInfo := map[string]interface{}{
		"title":        courseInfo.Title,
		"updated_time": nowTime,
		"desc":         courseInfo.Desc,
		"abstract":     courseInfo.Abstract,
	}
	err = global.GVA_DB.Model(course.CourseInfo{}).Where("id = ?", courseInfo.ID).Updates(&updateInfo).Error
	return err
}

// GetCourseInfo 根据id获取CourseInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseInfoService *CourseInfoService) GetCourseInfo(id uint) (courseInfo course.CourseInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&courseInfo).Error
	return
}

// GetCourseInfoInfoList 分页获取CourseInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseInfoService *CourseInfoService) GetCourseInfoInfoList(info courseReq.CourseInfoSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&course.CourseInfo{})
	var courseInfos []course.CourseInfo
	if len(info.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+info.Keyword+"%")
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var levelData []web.UserLevelPrice
	err = global.GVA_DB.Model(&web.UserLevelPrice{}).Find(&levelData).Error
	if err != nil {
		return
	}
	levelInfo := make(map[int8]string, 0)
	for i := 0; i < len(levelData); i++ {
		levelInfo[levelData[i].Grade] = levelData[i].Name
	}
	err = db.Limit(limit).Offset(offset).Where("status = 1").Find(&courseInfos).Error
	respData := make([]*resp.CourseInfoResp, 0)
	for i := 0; i < len(courseInfos); i++ {
		levelName := "免费课程"
		if courseInfos[i].Level != 0 {
			levelName = levelInfo[courseInfos[i].Level]
		}
		data := &resp.CourseInfoResp{
			ID:           courseInfos[i].ID,
			Title:        courseInfos[i].Title,
			LessonNumber: courseInfos[i].SectionNumber,
			CreatedTime:  courseInfos[i].CreatedTime,
			LevelName:    levelName,
		}
		respData = append(respData, data)
	}
	return respData, total, err
}
