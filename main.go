package main

import (
	"github.com/gin-gonic/gin"
	"go-web/core"
	"go-web/global"
	"go-web/initialize"
	"go.uber.org/zap"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release
func main() {
	//SetMode() 应该声明在 gin.New() 前，否则配置无法更新：
	gin.SetMode(AppMode)

	// TODO：1.配置初始化
	global.WEB_VIPER = core.InitializeViper()

	// TODO：2.日志
	global.WEB_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.WEB_LOG)

	global.WEB_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))

	// TODO：3.数据库连接
	global.WEB_DB = initialize.Gorm()
	// TODO：4.其他初始化
	initialize.OtherInit()
	// TODO：5.启动服务
	core.RunServer()

}
