package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/miaojiaxi2004/go_server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	VIPER  *viper.Viper
	DB     *gorm.DB
	REDIS  *redis.Client
	Router *gin.Engine
	LOG    *zap.Logger
)
