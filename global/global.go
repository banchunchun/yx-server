package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go-api/core"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MYSQL *gorm.DB
	VP    *viper.Viper
	Redis *redis.Client
	CF    *core.Config
	LOG   *zap.SugaredLogger
)

type Server interface {
	ListenAndServe() error
	Restart() error
	Shutdown()
}
