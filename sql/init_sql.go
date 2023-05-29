package sql

import (
	"fmt"

	"Gin_one/config"
)

func InitSql() {
	// 这里写sql的初始化-连接数据库
	database := config.GetConfig("database")
	fmt.Println(database)
}
