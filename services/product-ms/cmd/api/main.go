package main

import (
	"commons/otel"
	otelmiddleware "commons/otel/middleware"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	chandler "commons/handler"
	"product-ms/internal/infrastructure"
	"product-ms/internal/infrastructure/config"
	"product-ms/internal/repository/dao"
	"product-ms/internal/router"
	"product-ms/internal/service"

	chi "github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("Could not load config due to " + err.Error())
	}

	fmt.Printf("configuration: %+v\n", cfg)

	cleanup, err := otel.SetupOpenTelemetry(cfg.AppName, cfg.AppVersion)
	if err != nil {
		panic("Could not setup open telemetry: " + err.Error())
	}
	defer cleanup()

	infra, err := infrastructure.Load(cfg)
	if err != nil {
		panic("Could not load infrastructure due to " + err.Error())
	}

	queries := dao.New(infra.DB)
	productService := service.NewProductService(queries)
	storeService := service.NewStoreService(queries)
	productRouter := router.NewProductRouter(productService)
	storeRouter := router.NewStoreRouter(storeService)

	r := chi.NewRouter()
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(otelmiddleware.Metrics)      // otel metrics
	r.Use(otelmiddleware.TraceRequest) // otel trace
	r.Use(chimiddleware.Timeout(cfg.ReadTimeout + cfg.WriteTimeout))
	r.Get("/healthz", chandler.HealthCheck(time.Now(), cfg.AppName)) // health check traefik

	productRouter.RegisterRoutes(r)
	storeRouter.RegisterRoutes(r)

	server := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	go func() {
		println("Server is starting up", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("Failed to start server due to " + err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	println("Server is shutting down")
	if err := server.Shutdown(ctx); err != nil {
		println("Failed to gracefully shutdown " + err.Error())
	}
}
