package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	Port         int    `envconfig:"PORT"`
	LogLevel     string `envconfig:"LOG_LEVEL"`
	DatastoreURL string `envconfig:"DATASTORE_URL"`
	BrokerURL    string `envconfig:"BROKER_URL"`
}

func init() {
	viper.SetDefault("port", "3000")
	viper.SetDefault("logLevel", "info")
}

func LoadEnv() (*Config, error) {
	var instance Config
	if err := viper.Unmarshal(&instance); err != nil {
		return nil, err
	}

	err := envconfig.Process("", &instance)
	if err != nil {
		return &instance, err
	}

	return &instance, nil
}
