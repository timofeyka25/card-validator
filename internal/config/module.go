package config

import (
	"card-validator/internal/transport/http"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(
		New,
		func(cfg *Config) *http.Config {
			return cfg.HTTPConfig
		},
	),
)
