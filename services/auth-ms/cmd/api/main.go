package main

import (
	"log"

	"auth-ms/internal/app/server"
	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"

	"gcommons/otel"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	cleanup, err := otel.SetupOpenTelemetry(constants.ServiceName, constants.ServiceVersion)
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry: %v", err)
	}
	defer cleanup()

	if err := server.StartRESTServer(cfg); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
