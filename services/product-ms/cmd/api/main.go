package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"product-ms/infrastructure"
	"product-ms/internal/api"
	"product-ms/internal/config"
	"product-ms/internal/repository/dao"
	"product-ms/internal/service"

	chi "github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("Could not load config due to " + err.Error())
	}

	infra, err := infrastructure.Load(cfg)
	if err != nil {
		panic("Could not load infrastructure due to " + err.Error())
	}

	queries := dao.New(infra.DB)
	svc := service.NewProductService(queries)
	handler := api.NewProductHandler(svc)

	router := chi.NewRouter()
	api.Middlewares(cfg, router)
	handler.RegisterRoutes(router)

	server := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      router,
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	go func() {
		println("Server is starting up", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("Failed to start server due to " + err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	println("Server is shutting down")
	if err := server.Shutdown(ctx); err != nil {
		println("Failed to gracefully shutdown " + err.Error())
	}
}
