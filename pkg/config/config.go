package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	// Config 全局配置实例
	Config *viper.Viper
)

// Init 初始化配置
func Init(configPath string) error {
	Config = viper.New()

	// 设置配置文件路径
	Config.SetConfigFile(configPath)

	// 设置环境变量前缀
	Config.SetEnvPrefix("MIN_DSP")
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	Config.AutomaticEnv()

	// 读取配置文件
	if err := Config.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 监听配置文件变化
	Config.WatchConfig()

	return nil
}

// GetString 获取字符串配置
func GetString(key string) string {
	return Config.GetString(key)
}

// GetInt 获取整数配置
func GetInt(key string) int {
	return Config.GetInt(key)
}

// GetBool 获取布尔配置
func GetBool(key string) bool {
	return Config.GetBool(key)
}

// GetStringSlice 获取字符串切片配置
func GetStringSlice(key string) []string {
	return Config.GetStringSlice(key)
}

// GetStringMap 获取字符串映射配置
func GetStringMap(key string) map[string]interface{} {
	return Config.GetStringMap(key)
}

// GetStringMapString 获取字符串到字符串的映射配置
func GetStringMapString(key string) map[string]string {
	return Config.GetStringMapString(key)
}

// Sub 获取子配置
func Sub(key string) *viper.Viper {
	return Config.Sub(key)
}

// AllSettings 获取所有配置
func AllSettings() map[string]interface{} {
	return Config.AllSettings()
}
