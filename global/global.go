package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go-api/core"
	"gorm.io/gorm"
)

var (
	MYSQL  *gorm.DB
	VP     *viper.Viper
	Redis  *redis.Client
	Config *core.Config
)

type Server interface {
	ListenAndServe() error
	Restart() error
	Shutdown()
}
