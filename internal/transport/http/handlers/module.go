package handlers

import (
	"card-validator/internal/transport/http"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers",
	fx.Provide(
		NewCardHandler,
		aggregateHandlers,
	),
)

func aggregateHandlers(cardHandler *CardHandler) []http.Handler {
	return []http.Handler{
		cardHandler,
	}
}
