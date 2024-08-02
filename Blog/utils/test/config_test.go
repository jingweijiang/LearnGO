package utils

import (
	"blog/utils"
	"fmt"
	"testing"
	"time"
)

func TestGetOnCurrentPath(t *testing.T) {
	fmt.Println(utils.ProjectRootPath)
}

func TestConfig(t *testing.T) {
	dbViper := utils.CreateConfig("mysql")
	dbViper.WatchConfig() // 监听配置文件变化
	// 方式一：直接获取配置项
	// 检查是否有配置项
	if !dbViper.IsSet("blog.host") {
		t.Fail()
	}
	fmt.Println(dbViper.GetString("blog.host"))
	fmt.Println(dbViper.GetInt("blog.port"))
	fmt.Println(dbViper.GetString("blog.user"))
	fmt.Println(dbViper.GetString("blog.password"))
	fmt.Println(dbViper.GetString("blog.database"))
	time.Sleep(10 * time.Second)
	fmt.Println(dbViper.GetInt("blog.port"))

	// 方式二：通过结构体获取配置项
	logViper := utils.CreateConfig("log")
	type LogConfig struct {
		Level string `mapstructure:"level"`
		Path  string `mapstructure:"path"`
	}
	var logConfig LogConfig
	if err := logViper.Unmarshal(&logConfig); err != nil {
		t.Fail()
	}
	fmt.Println(logConfig.Level)
	fmt.Println(logConfig.Path)

}
