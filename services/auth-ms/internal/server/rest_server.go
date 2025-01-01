package server

import (
	"auth-ms/internal/handlers"
	"auth-ms/internal/repository"
	"auth-ms/internal/services"
	"fmt"
	"log"
	"net/http"
	"time"

	"commons/authen"
	ghandler "commons/handler"
	gmiddleware "commons/middleware"
	"commons/otel"
	otelmiddleware "commons/otel/middleware"

	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"
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
	tokenAdapter, err := authen.NewTokenGenerator(cfg.PrivateKeyPath, cfg.PublicKeyPath)
	if err != nil {
		log.Fatal(fmt.Errorf("token generator initialization failed: %w", err))
	}

	userProvider := repository.NewUserProviderImpl(cfg)
	authService := services.NewAuthServiceImpl(cfg, tokenAdapter, userProvider)
	authHandler := handlers.NewAuthHandler(otel.GetTracer(), authService)

	root := http.NewServeMux()
	{
		root.HandleFunc("/healthz", ghandler.HealthCheck(time.Now(), constants.ServiceName))
	}

	v1 := http.NewServeMux()
	root.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))
	{
		v1.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Hi there! This is the v1 API"))
		})
		v1.HandleFunc("POST /token", authHandler.HandleAccessToken)
		v1.HandleFunc("GET /auth", authHandler.HandleAuthenticate)
	}

	return root
}
