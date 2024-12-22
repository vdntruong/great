package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"gcommons/otel"

	"user-ms/internal/app"
	"user-ms/internal/pkg/config"
)

func main() {

	// load configuration, and infrastructures
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config failed: %s", err.Error())
	}

	cleanup, err := otel.SetupOpenTelemetry(cfg.AppName, cfg.AppVersion)
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry: %v", err)
	}
	defer cleanup()

	// init the main application
	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to init application: %v", err)
	}

	// serve the API
	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: application.Routes(),

		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	go func() {
		log.Println("Starting server:", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Server killing:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting")
}
