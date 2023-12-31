package core

import (
	"go-web/core/internal"
	"go-web/global"
	"go-web/utils"

	"fmt"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitializeZap 日志初始化方法
// InitializeZap Zap 获取 zap.Logger
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.WEB_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.WEB_CONFIG.Zap.Director)
		_ = os.Mkdir(global.WEB_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.WEB_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("====2-zap====: zap log init success")
	return logger
}
