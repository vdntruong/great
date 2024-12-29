package server

import (
	"log"
	"net/http"
	"time"

	"auth-ms/internal/app/handlers"
	"auth-ms/internal/app/repository"
	"auth-ms/internal/app/services"
	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"

	ghandler "gcommons/handler"
	gmiddleware "gcommons/middleware"
	otelmiddleware "gcommons/otel/middleware"
)

func StartRESTServer(cfg *config.Config) error {
	route := routes(cfg)
	handler := gmiddleware.RecoverPanic(route)
	handler = otelmiddleware.Metrics(handler)
	handler = otelmiddleware.TraceRequest(handler)
	handler = gmiddleware.LogRequest(handler)

	log.Printf("Authentication service starting on %s\n", cfg.RESTAddress)
	return http.ListenAndServe(cfg.RESTAddress, handler)
}

func routes(cfg *config.Config) *http.ServeMux {
	userProvider := repository.NewUserProviderImpl(cfg)
	authService := services.NewAuthServiceImpl(userProvider)
	authHandler := handlers.NewAuthHandler(authService)

	root := http.NewServeMux()
	{
		root.HandleFunc("/healthz", ghandler.HealthCheck(time.Now(), constants.ServiceName))
	}

	v1 := http.NewServeMux()
	root.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))
	{
		v1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Hi there! This is the v1 API"))
		})
		v1.HandleFunc("POST /login", authHandler.Login)
	}

	return root
}
