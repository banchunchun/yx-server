package core

import (
	"flag"
	"fmt"
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
	Param    string `yaml:"param"`
}
type Rds struct {
	Network string `yaml:"network"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	DB      int    `yaml:"db"`
	Pass    string `yaml:"pass"`
}
type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Rds    `yaml:"redis"`
}

// LoadConfig 加载配置文件
func LoadConfig(path ...string) *Config {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			config = "./configs/"
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	}
	viper.AddConfigPath(config)
	viper.SetConfigName("application.yml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("load env file error!")
		} else {
			log.Fatalf("load config error %s \n", err.Error())
		}
	}
	var c *Config
	if err := viper.Unmarshal(&c); err != err {
		log.Fatal("load config value error!")
	}
	return c
}
