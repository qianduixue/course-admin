package course

import (
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/common/request"
	"github.com/opisnoeasy/course-service/model/course"
	courseReq "github.com/opisnoeasy/course-service/model/course/request"
	"github.com/opisnoeasy/course-service/model/resp"
	"gorm.io/gorm"
	"time"
)

type CourseSectionService struct {
}

// CreateCourseSection 创建CourseSection记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseSectionService *CourseSectionService) CreateCourseSection(courseSection course.CourseSection) (err error) {
	nowTime := time.Now().Unix()
	courseSection.CreatedTime = nowTime
	tx := global.GVA_DB.Begin()
	err = tx.Create(&courseSection).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//给课程增加一小节
	err = tx.Model(course.CourseInfo{}).Where("id = ?", courseSection.CourseId).Update("section_number", gorm.Expr("section_number + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// DeleteCourseSection 删除CourseSection记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseSectionService *CourseSectionService) DeleteCourseSection(courseSection course.CourseSection) (err error) {
	updateInfo := map[string]interface{}{
		"updated_time": time.Now().Unix(),
		"status":       2,
	}
	//查询课程
	var section course.CourseSection
	err = global.GVA_DB.Model(course.CourseSection{}).Where("id = ?", courseSection.ID).First(&section).Error
	if err != nil {
		return err
	}
	if section.Status == 2 {
		return nil
	}
	tx := global.GVA_DB.Begin()
	err = tx.Model(course.CourseSection{}).Where("id = ?", courseSection.ID).Updates(&updateInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(course.CourseInfo{}).Where("id = ?", section.CourseId).Update("section_number", gorm.Expr("section_number - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// DeleteCourseSectionByIds 批量删除CourseSection记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseSectionService *CourseSectionService) DeleteCourseSectionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]course.CourseSection{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCourseSection 更新CourseSection记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseSectionService *CourseSectionService) UpdateCourseSection(courseSection course.CourseSection) (err error) {
	updateInfo := map[string]interface{}{
		"course_id":    courseSection.CourseId,
		"cover_image":  courseSection.CoverImage,
		"lessons":      courseSection.Lessons,
		"long_time":    courseSection.LongTime,
		"title":        courseSection.Title,
		"video_url":    courseSection.VideoUrl,
		"updated_time": time.Now().Unix(),
	}
	err = global.GVA_DB.Model(course.CourseSection{}).Where("id = ?", courseSection.ID).Updates(&updateInfo).Error
	return err
}

// GetCourseSection 根据id获取CourseSection记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseSectionService *CourseSectionService) GetCourseSection(id uint) (courseSection course.CourseSection, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&courseSection).Error
	return
}

// GetCourseSectionInfoList 分页获取CourseSection记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseSectionService *CourseSectionService) GetCourseSectionInfoList(info courseReq.CourseSectionSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&course.CourseSection{})
	var courseSections []course.CourseSection
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if len(info.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+info.Keyword+"%")
	}
	err = db.Limit(limit).Offset(offset).Where("status = 1").Find(&courseSections).Error
	respData := make([]*resp.CourseSectionResp, 0)
	for i := 0; i < len(courseSections); i++ {
		data := &resp.CourseSectionResp{
			Id:          courseSections[i].ID,
			Lessons:     courseSections[i].Lessons,
			Tittle:      courseSections[i].Title,
			LongTime:    courseSections[i].LongTime,
			CreatedTime: courseSections[i].CreatedTime,
		}
		respData = append(respData, data)
	}
	return respData, total, err
}
