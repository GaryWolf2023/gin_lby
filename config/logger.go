package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if viper.GetBool("mode.development") {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encodingConfig := zap.NewProductionEncoderConfig()
	encodingConfig.TimeKey = "time"
	encodingConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodingConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encodingConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	stSeparater := string(filepath.Separator)
	// 获取系统目录
	stRootDir, _ := os.Getwd()
	// 日志输出目录
	stLogFilePath := stRootDir + stSeparater + "log" + stSeparater + time.Now().Format(time.DateOnly)
	fmt.Println(stLogFilePath)
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.maxSize"),    // megabytes
		MaxBackups: viper.GetInt("log.maxBackups"), // 最多保留多少个备份的旧文件
		MaxAge:     viper.GetInt("log.maxAge"),     //days
		Compress:   false,                          // disabled by default
	}
	return zapcore.AddSync(lumberjackSyncer)
}
