package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	APIEndpoint       string `mapstructure:"api_endpoint"`
	NotificationEmail string `mapstructure:"notification_email"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
