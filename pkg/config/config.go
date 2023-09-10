package config

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
}

func (c *Config) init() error {
	flag.String("config", "./config.yaml", "Path to the yaml config file")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return fmt.Errorf("viper.BindPFlags: %w", err)
	}
	viper.SetConfigFile(viper.GetString("config"))
	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("viper.ReadInConfig: %w", err)
	}
	return nil
}

func New() (*Config, error) {
	c := new(Config)
	err := c.init()
	if err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}
	return c, nil
}
