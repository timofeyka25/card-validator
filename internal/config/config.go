package config

import (
	"card-validator/internal/transport/http"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sync"
)

var (
	err    error
	config *Config
	once   sync.Once
)

type Config struct {
	LogLevel   string
	HTTPConfig *http.Config
}

func New() (*Config, error) {
	once.Do(func() {
		config = &Config{}

		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		httpConfig := viper.Sub("server")

		if err = parseSubConfig(httpConfig, &config.HTTPConfig); err != nil {
			return
		}

	})

	zap.S().Info(config)

	return config, err
}

func parseSubConfig[T any](subConfig *viper.Viper, parseTo *T) error {
	if subConfig == nil {
		return fmt.Errorf("can not read %T config: subconfig is nil", parseTo)
	}

	if err = subConfig.Unmarshal(parseTo); err != nil {
		return err
	}

	return nil
}
