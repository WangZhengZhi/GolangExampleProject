package config

import "github.com/spf13/viper"

// ReadYml 读取配置文件
func ReadYml() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yml")
	v.SetConfigName("core")
	return v
}
