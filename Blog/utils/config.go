package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var (
	ProjectRootPath = getOnCurrentPath()
)

func getOnCurrentPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(path.Dir(filename))

}

func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	// 获取config目录的路径
	configPath := ProjectRootPath + "/config/"
	config.AddConfigPath(configPath)
	config.SetConfigName(file)
	config.SetConfigType("yaml")
	configFilePath := configPath + file + ".yaml"
	// 如果配置文件不存在，则创建
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found: %s", configFilePath))
		} else {
			// 配置文件有错误
			panic(fmt.Errorf("config file %s content error: %s", configFilePath, err))
		}

	}
	return config

}
