package main

import (
	"fmt"
	"min-dsp/pkg/config"
	"min-dsp/pkg/logger"
	"min-dsp/router"

	"go.uber.org/zap"
)

// main 是程序入口，启动HTTP服务
func main() {
	// 初始化配置
	if err := config.Init("config/config.yaml"); err != nil {
		panic(fmt.Sprintf("初始化配置失败: %v", err))
	}

	// 初始化日志模块
	logConfig := &logger.Config{
		Level:      config.GetString("log.level"),
		Filename:   config.GetString("log.filename"),
		MaxSize:    config.GetInt("log.max_size"),
		MaxBackups: config.GetInt("log.max_backups"),
		MaxAge:     config.GetInt("log.max_age"),
		Compress:   config.GetBool("log.compress"),
	}
	if err := logger.InitLogger(logConfig); err != nil {
		panic(fmt.Sprintf("初始化日志模块失败: %v", err))
	}

	// 启动HTTP服务
	r := router.SetupRouter()
	port := config.GetString("app.port")
	logger.Info("服务启动",
		zap.String("app", config.GetString("app.name")),
		zap.String("version", config.GetString("app.version")),
		zap.String("port", port),
		zap.String("mode", config.GetString("app.mode")),
	)

	if err := r.Run(":" + port); err != nil {
		logger.Fatal("服务启动失败", zap.Error(err))
	}
}
