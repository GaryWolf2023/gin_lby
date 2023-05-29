package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var ViperMap map[string]string
var viperConfig *viper.Viper

func InitConfig() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("读取配置文件失败", err)
		}
	}()
	// 获取配置文件
	viperConfig = viper.New()
	viperConfig.AddConfigPath("./config/") //指定配置文件所在的路径
	viperConfig.SetConfigName("app")
	viperConfig.SetConfigType("yaml")
	err := viperConfig.ReadInConfig() //读取配置文件
	if err != nil {
		panic(err)
	}
	// fmt.Println("-----------------------------------", viperConfig.Get("database"))
	// fmt.Println("+++++++++++++++++++++++++++++++++==", viperConfig.Get(""))
	//获取所有配置信息
	// allsetting := viperConfig.AllSettings()
	// fmt.Println("****************************************", allsetting)
}

// 传入对应的key，获取对应的value
func GetConfig(key string) interface{} {
	return viperConfig.Get(key)
}
