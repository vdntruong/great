package main

import (
	"log"

	"auth-ms/internal/app/server"
	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"

	"gcommons/otel"
)

func main() {
	cleanup, err := otel.SetupOpenTelemetry(constants.ServiceName)
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry: %v", err)
	}
	defer cleanup()

	if err := config.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := server.StartRESTServer(config.Cfg, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
