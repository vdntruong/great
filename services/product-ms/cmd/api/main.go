package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"commons/db/postgre"
	"commons/otel/db"

	"product-ms/internal/api"
	"product-ms/internal/config"
	"product-ms/internal/repository"

	chi "github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("Failed to load config due to " + err.Error())
	}

	dbCfg := postgre.Config{
		Host:               cfg.DBHost,
		Port:               cfg.DBPort,
		Username:           cfg.DBUsername,
		Password:           cfg.DBPassword,
		DatabaseName:       cfg.DBName,
		MaxConnections:     cfg.DBMaxConnections,
		MaxIdleConnections: cfg.DBMaxIdleConnections,
	}
	sqlDB, cleanup, err := db.NewDB(dbCfg.GetDataSourceName(), dbCfg.DatabaseName, cfg.DBMaxConnections, cfg.DBMaxIdleConnections)
	if err != nil {
		panic("Failed to connect database")
	}
	defer cleanup()

	productRepo := repository.NewProductRepository(sqlDB)
	productSvc := api.NewProductService(productRepo)
	productHandler := api.NewProductHandler(productSvc)

	r := chi.NewRouter()
	api.Middlewares(cfg, r)
	productHandler.RegisterRoutes(r)

	server := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	go func() {
		println("Server is starting up", cfg.Addr())
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
