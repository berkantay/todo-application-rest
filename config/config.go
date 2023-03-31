package config

import (
	"context"

	"github.com/spf13/viper"
)

type Config struct {
	Postgresql Postgresql
}

func NewConfig(ctx context.Context, configName string, configPath string) (*Config, error) {
	v, err := readConfig(configName, configPath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func readConfig(configName, configPath string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	err := v.ReadInConfig()
	return v, err
}
