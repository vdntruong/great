package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"commons/authen"
	"commons/handler"
	"commons/middleware"
	"commons/otel"
	omiddleware "commons/otel/middleware"

	"auth-ms/internal/handlers"
	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"
	"auth-ms/internal/services"
	"auth-ms/internal/services/adapter"
)

func StartRESTServer(cfg *config.Config) error {
	route := routes(cfg)
	handler := middleware.RecoverPanic(route)
	handler = omiddleware.Metrics(handler)
	handler = omiddleware.TraceRequest(handler)
	handler = middleware.LogRequest(handler)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.RESTPort),
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	log.Printf("Authentication service starting on %s\n", cfg.RESTPort)
	if err := server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func routes(cfg *config.Config) *http.ServeMux {
	tokenAdapter, err := authen.NewTokenGenerator(cfg.PrivateKeyPath, cfg.PublicKeyPath)
	if err != nil {
		log.Fatal(fmt.Errorf("token generator initialization failed: %w", err))
	}

	userProvider := adapter.NewUserAdapter(cfg)
	authService := services.NewAuthServiceImpl(cfg, tokenAdapter, userProvider)
	authHandler := handlers.NewAuthHandler(otel.GetTracer(), authService)

	root := http.NewServeMux()
	{
		root.HandleFunc("/healthz", handler.HealthCheck(time.Now(), constants.ServiceName))
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
