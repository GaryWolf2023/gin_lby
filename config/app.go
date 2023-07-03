package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var ViperConfig *viper.Viper

func InitConfig() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("读取配置文件失败", err)
		}
	}()
	// 获取配置文件
	ViperConfig = viper.New()
	ViperConfig.AddConfigPath("./config/") //指定配置文件所在的路径
	ViperConfig.SetConfigName("app")
	ViperConfig.SetConfigType("yaml")
	err := ViperConfig.ReadInConfig() //读取配置文件
	if err != nil {
		panic(err)
	}
}
