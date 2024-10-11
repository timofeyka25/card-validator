package di

import (
	"card-validator/internal/config"
	"card-validator/internal/logger"
	"card-validator/internal/repositories"
	"card-validator/internal/services"
	"card-validator/internal/transport/http"
	"card-validator/internal/transport/http/handlers"
	"card-validator/pkg/validator"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		config.Module,
		logger.Module,
		services.Module,
		handlers.Module,
		repositories.Module,
		http.Module,
		validator.Module,
	)
}
