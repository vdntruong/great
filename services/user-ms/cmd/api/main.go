package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"commons/otel"

	"user-ms/internal/app"
	"user-ms/internal/pkg/config"
)

func main() {

	// load configuration, and infrastructures
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config failed: %s", err.Error())
	}

	log.Printf("otel endpoint config: %+v", cfg.OTLPEndpoint)

	cleanup, err := otel.SetupOpenTelemetry(cfg.AppName, cfg.AppVersion)
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry: %v", err)
	}
	defer cleanup()

	// init the main application
	application, cleanups, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to init application: %v", err)
	}
	defer func() {
		for _, c := range cleanups {
			c()
		}
	}()

	// serve the REST API
	restSrv := application.InitRestServer()
	go func() {
		log.Println("Starting REST server:", restSrv.Addr)
		if err := restSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("REST killing:", err)
		}
	}()

	// serve the gRPC API
	grpcLis, grpcSrv := application.InitGRPCServer()
	go func() {
		log.Println("Starting gRPC server:", grpcLis.Addr().String())
		if err := grpcSrv.Serve(grpcLis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	grpcSrv.GracefulStop()
	if err := restSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	log.Println("Server exiting")
}
