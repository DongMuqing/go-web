package global

import (
	"github.com/go-redis/redis/v8"
	"go-web/config"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

// 项目全局变量
var (
	WEB_CONFIG config.Configuration
	WEB_VIPER  *viper.Viper
	WEB_LOG    *zap.Logger
	WEB_DB     *gorm.DB
	WEB_REDIS  *redis.Client
)
