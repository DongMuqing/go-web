package initialize

import (
	"go-web/global"

	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.WEB_CONFIG.App.DbType {
	case "mysql":
		return GormMysql
	default:
		return GormMysql()
	}
}
