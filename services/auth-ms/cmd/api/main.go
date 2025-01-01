package main

import (
	"log"

	"commons/discovery"
	"commons/otel"

	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"
	"auth-ms/internal/server"
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

	if err := discovery.Register(cfg.ServiceID, cfg.ServiceName, "auth-ms", cfg.RESTPort, []string{"rest"}, map[string]string{"protocol": "http"}); err != nil {
		log.Printf("[ERROR] failed to register auth-ms: %v\n", err)
	}

	if err := server.StartRESTServer(cfg); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
