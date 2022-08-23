package initialize

import (
	"os"

	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/example"
	"github.com/opisnoeasy/course-service/model/system"

	"github.com/opisnoeasy/course-service/model/course"
	"github.com/opisnoeasy/course-service/model/web"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		// 示例模块表
		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		// 自动化模块表
		// Code generated by github.com/opisnoeasy/course-service Begin; DO NOT EDIT.

		web.NotesInfo{},
		course.CourseInfo{},
		web.UserLevelPrice{},
		web.MemberInterests{},
		course.CourseSection{},
		web.ContactUs{},
		web.PlatformRegister{},
		web.UserInfo{},
		web.Order{},
		// Code generated by github.com/opisnoeasy/course-service End; DO NOT EDIT.
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
