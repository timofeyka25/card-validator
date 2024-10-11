package logger

import (
	"card-validator/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("logger",
	fx.Provide(NewLogger),
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	zapCfg := zap.NewDevelopmentConfig()
	level, err := zap.ParseAtomicLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	zapCfg.Level = level

	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)
	return logger, nil
}
