package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var ViperMap map[string]string
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
	// fmt.Println("-----------------------------------", ViperConfig.Get("database"))
	// fmt.Println("+++++++++++++++++++++++++++++++++==", ViperConfig.Get(""))
	//获取所有配置信息
	// allsetting := ViperConfig.AllSettings()
	// fmt.Println("****************************************", allsetting)
}

// 传入对应的key，获取对应的value
func GetConfig(key string) interface{} {
	return ViperConfig.Get(key)
}
