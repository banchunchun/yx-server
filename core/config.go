package core

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	Port string `yaml:"port"`
}

type Config struct {
	Server     Server `yaml:"server"`
	Mysql      Mysql  `yaml:"mysql"`
	Redis      Rds    `yaml:"redis"`
	Log        Log    `yaml:"log"`
	Transcoder TS     `yaml:"ts"`
}
type Log struct {
	Level        string //等级
	Formatter    string //格式化模式
	ShowLine     bool   //是否展示行数
	LogInConsole bool   //是否写入文件的同时写入标准输出
	OutFile      bool   //是否写入文件
	LogDir       string //日志文件夹
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

type TS struct {
	WorkDir         string `yaml:"workDir"`
	XmlPath         string `yaml:"xmlPath"`
	ImageExeWorkDir string `yaml:"imageExeWorkDir"`
	ImageExeName    string `yaml:"imageExeName"`
}

// LoadConfig 加载配置文件
func LoadConfig(path ...string) *Config {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			config = "./configs/"
			fmt.Println("您正在使用config的默认值,config的路径为%" + config)
		} else {
			fmt.Println("您正在使用命令行的-c参数传递的值,config的路径为" + config)
		}
	}
	viper.AddConfigPath(config)
	viper.SetConfigName("application.yml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("load env file error!", zap.Error(err))
		} else {
			fmt.Println("load config error %s \n", zap.Error(err))
		}
	}
	var c *Config
	if err := viper.Unmarshal(&c); err != err {
		fmt.Println("load config value error %s \n", zap.Error(err))
	}
	return c
}
