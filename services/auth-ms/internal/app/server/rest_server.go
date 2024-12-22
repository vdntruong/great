package server

import (
	"log"
	"net/http"
	"time"

	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"

	ghandler "gcommons/handler"
	gmiddleware "gcommons/middleware"
	gotel "gcommons/otel"
	otelmiddleware "gcommons/otel/middleware"
)

func StartRESTServer(cfg *config.Config) error {
	route := routes()
	handler := gmiddleware.LogRequest(route)
	handler = gmiddleware.RecoverPanic(handler)
	handler = otelmiddleware.Metrics(handler)

	log.Printf("Authentication service starting on %s\n", cfg.RESTAddress)
	return http.ListenAndServe(cfg.RESTAddress, handler)
}

func routes() *http.ServeMux {
	root := http.NewServeMux()

	v1 := http.NewServeMux()
	root.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	// root
	{
		root.HandleFunc("/healthz", ghandler.HealthCheck(time.Now(), constants.ServiceName))
	}

	// /api/v1/
	{
		v1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, span := gotel.GetTracer().Start(r.Context(), "v1 welcome")
			defer span.End()
			_, _ = w.Write([]byte("Hi there! This is the v1 API"))
		})
	}

	return root
}
