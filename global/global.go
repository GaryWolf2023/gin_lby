package global

import (
	"Gin_one/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.SugaredLogger // 定义一个Logger来接收初始化日志的结构，还可以进一步的封装Error，Warning等方法方便使用
var DB *gorm.DB
var Rdb *config.RedisClient
