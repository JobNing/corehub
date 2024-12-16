package config

import (
	"github.com/spf13/viper"
)

func InitViper(add string) error {
	viper.SetConfigFile(add)    // 配置文件类型
	return viper.ReadInConfig() // Find and read the config file
}
