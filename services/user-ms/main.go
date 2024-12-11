package main

import (
	"log"
	"net/http"
	"time"

	"user-ms/constants"

	"gcommons/handler"
	"gcommons/otel"
)

func main() {
	cleanup, err := otel.SetupOpenTelemetry(constants.ServiceName)
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry: %v", err)
	}
	defer cleanup()

	http.HandleFunc("/healthz", handler.HealthCheck(time.Now(), "User"))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
