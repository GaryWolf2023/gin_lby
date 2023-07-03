package sql

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitSql() {
	// 这里写sql的初始化-连接数据库
	database := viper.GetString("database")
	fmt.Println(database)
}
