package main

import (
	"card-validator/internal/di"
	"go.uber.org/zap"
	"time"
)

func main() {
	now := time.Now()

	app := di.Build()

	zap.S().Info("Starting application...")

	zap.S().Infof("Up and running (%s)", time.Since(now))

	app.Run()

	zap.S().Info("Service stopped.")
}
