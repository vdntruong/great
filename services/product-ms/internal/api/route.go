package api

import (
	"time"

	"commons/handler"
	otelmiddleware "commons/otel/middleware"

	"product-ms/internal/config"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Middlewares(cfg *config.Config, r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(otelmiddleware.Metrics)      // otel metrics
	r.Use(otelmiddleware.TraceRequest) // otel trace
	r.Use(middleware.Timeout(cfg.ReadTimeout + cfg.WriteTimeout))
	r.Get("/healthz", handler.HealthCheck(time.Now(), cfg.AppName)) // health check traefik
}
