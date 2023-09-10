package config

import "github.com/spf13/viper"

func (c *Config) GetListen() string {
	return viper.GetString("listen")
}
