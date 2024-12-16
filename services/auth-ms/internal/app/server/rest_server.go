package server

import (
	"log"
	"net/http"
	"time"

	"auth-ms/internal/app/handlers"
	"auth-ms/internal/app/repository"
	"auth-ms/internal/pkg/config"
	"auth-ms/internal/pkg/constants"

	ghandler "gcommons/handler"
	"gcommons/middleware"
	"gcommons/otel"
)

func StartRESTServer(cfg *config.Config, userRepo repository.UserRepository) error {
	authHandler := handlers.NewAuthRESTHandler(userRepo)

	route := registerRoutes(authHandler)
	handler := middleware.LoggingMiddleware(route)
	handler = middleware.RecoveryMiddleware(handler)
	handler = otel.MetricsMiddleware(handler)

	log.Printf("Authentication service starting on %s\n", cfg.RESTAddress)
	return http.ListenAndServe(cfg.RESTAddress, handler)
}

func registerRoutes(authHandler *handlers.AuthRESTHandler) *http.ServeMux {
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
			_, span := otel.GetTracer().Start(r.Context(), "v1 welcome")
			defer span.End()
			_, _ = w.Write([]byte("Hi there! This is the v1 API"))
		})

		v1.HandleFunc("/register", authHandler.RegisterHandler)
		v1.HandleFunc("/login", authHandler.LoginHandler)
	}

	return root
}
