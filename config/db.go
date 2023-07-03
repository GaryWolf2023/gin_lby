package config

import (
	"Gin_one/models"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.development") {
		logMode = logger.Error
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++", ViperConfig.Get("database.dns"))
	db, err := gorm.Open(mysql.Open(ViperConfig.GetString("database.dns")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(ViperConfig.GetInt("database.MaxOpenConns"))
	sqlDB.SetMaxIdleConns(ViperConfig.GetInt("database.maxIdleConns"))

	db.AutoMigrate(&models.UserBasic{})
	return db, nil
}
