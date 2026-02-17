// 配置层 存放程序的配置信息
package config

import (
	"golang_project_scaffold/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 配置参数
const (
	// 配置日志时间格式
	timeFormat = string("2006-01-02 15:04:05")
)

var (
	Port string
)

// 配置日志输出格式
func initLogConfig(logLevel string) {
	// 控制日志的输出级别
	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}
	// 添加文件名和行号
	logrus.SetReportCaller(true)
	// 配置日志输出为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: timeFormat})

}

func init() {
	logs.Info(nil, "Start loading the program configuration...")
	// 配置日志输出级别
	viper.SetDefault("LOG_LEVEL", "debug")
	// 配置程序监听端口
	viper.SetDefault("PORT", ":8080")
	// 绑定环境变量到配置
	viper.AutomaticEnv()
	// 环境变量赋值到程序变量
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	// 加载日志输出级别
	initLogConfig(logLevel)
	logs.Info(nil, "The log configuration loading is complete.")
}
