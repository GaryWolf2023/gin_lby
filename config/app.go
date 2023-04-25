package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var ViperMap map[string]string

func InitConfig() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("读取配置文件失败", err)
		}
	}()
	// 获取配置文件
	viper.SetConfigName("app")    //指定配置文件的文件名称(不需要指定配置文件的扩展名)
	viper.AddConfigPath("config") //指定配置文件所在的路径
	err := viper.ReadInConfig()   //读取配置文件
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("sql"))
}
