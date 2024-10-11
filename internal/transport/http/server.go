package http

import (
	_ "card-validator/docs"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	router *gin.Engine
}

// NewServer
// @title           API
// @version         1.0
// @description     This is a sample server.
// @host localhost:8000
// @BasePath /
func NewServer(cfg *Config, handlers []Handler) *Server {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	s := &Server{
		server: &http.Server{
			Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:           router,
			ReadHeaderTimeout: 10 * time.Second,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       30 * time.Second,
		},
		router: router,
	}

	api := s.router.Group("")
	for _, handler := range handlers {
		handler.Register(api)
	}

	s.router.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return s
}

func RunServer(lc fx.Lifecycle, server *Server) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				zap.S().Info("Starting HTTP server...")
				if err := server.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					zap.S().Error("Error starting server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Info("Shutting down HTTP server...")
			if err := server.server.Shutdown(ctx); err != nil {
				zap.S().Error("Server shutdown error", zap.Error(err))
				return err
			}
			zap.S().Info("Server successfully stopped.")
			return nil
		},
	})
}
