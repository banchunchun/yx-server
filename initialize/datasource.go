package initialize

import (
	"fmt"
	"go-api/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return LoadMySQL()
}

func LoadMySQL() *gorm.DB {
	dsn := GetMysqlDSN()
	db, _ := gorm.Open(mysql.Open(dsn), GormConfig())
	return db
}

func GetMysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		global.CF.Mysql.User,
		global.CF.Mysql.Pass,
		global.CF.Mysql.Host,
		global.CF.Mysql.Database,
		global.CF.Mysql.Param,
	)
}
func GormConfig() *gorm.Config {
	c := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		/* NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.MysqlPrefix, // 表前缀
			SingularTable: true,          // 使用单数表名
		}, */
	}
	return c
}
