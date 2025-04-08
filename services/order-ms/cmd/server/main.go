package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"commons/otel"

	"order-ms/db/dao"
	"order-ms/internal/infras"
	"order-ms/internal/infras/config"
	"order-ms/internal/router"
	"order-ms/internal/service"
	"order-ms/internal/service/validator"

	_ "github.com/lib/pq"
)

func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Could not load config due to " + err.Error())
	}

	fmt.Printf("configuration: %+v\n", cfg)

	cleanup, err := otel.SetupOpenTelemetry(cfg.AppName, cfg.AppVersion)
	if err != nil {
		panic("Could not setup open telemetry: " + err.Error())
	}
	defer cleanup()

	// Initialize application infrastructures
	infra, err := infras.Load(cfg)
	if err != nil {
		panic("Could not load infrastructure due to " + err.Error())
	}
	defer infra.DB.Close()

	// Initialize queries
	queries := dao.New(infra.DB)

	// Initialize services
	orderService := service.NewOrderServiceAdapter(queries)
	cartService := service.NewCartServiceAdapter(queries)

	// Initialize validators
	cartValidator := validator.NewCartValidator()
	orderValidator := validator.NewOrderValidator()

	// Initialize router
	r := router.NewRouter(cfg, orderService, orderValidator, cartService, cartValidator)

	// Create server
	srv := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
