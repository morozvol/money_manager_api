package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BindAddr string `mapstructure:"bind_addr"`
	LogLevel string `mapstructure:"log_level"`
}

func Init(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
