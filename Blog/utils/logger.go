package utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var LogRus *logrus.Logger

func InitLogger(configFile string) {
	viper := CreateConfig(configFile)
	LogRus = logrus.New()

	// Set log level
	switch strings.ToLower(viper.GetString("level")) {
	case "debug":
		LogRus.SetLevel(logrus.DebugLevel)
	case "info":
		LogRus.SetLevel(logrus.InfoLevel)
	case "warn":
		LogRus.SetLevel(logrus.WarnLevel)
	case "error":
		LogRus.SetLevel(logrus.ErrorLevel)
	case "panic":
		LogRus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level: %s", viper.GetString("level")))
	}

	// Set log format
	LogRus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	logFile := ProjectRootPath + viper.GetString("path")
	fOut, err := rotatelogs.New(

		logFile+".%Y%m%d%H",                      // 日志文件名按天切分
		rotatelogs.WithLinkName(logFile),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),    // 日志文件最大保存时间7天
		rotatelogs.WithRotationTime(1*time.Hour), // 日志切割时间间隔 1小时
	)
	if err != nil {
		panic(err)
	}
	LogRus.SetOutput(fOut)       // 设置日志输出到文件
	logrus.SetReportCaller(true) // 设置显示调用函数信息

}
