package conf

import (
	"github.com/spf13/viper"
	"log"
)

type Server struct {
	Port string `yaml:"port"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Database string `yaml:"database"`
}
type Rds struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB   string `yaml:"db"`
	Pass string `yaml:"pass"`
}
type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Rds    `yaml:"redis"`
}

// LoadConfig 加载配置文件
func LoadConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("application.yml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("load env file error!")
		} else {
			log.Fatalf("load config error %s \n", err.Error())
		}
	}
	var config *Config
	if err := viper.Unmarshal(&config); err != err {
		log.Fatal("load config value error!")
	}
	return config
}
