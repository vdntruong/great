package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"commons/otel"

	"product-ms/db/dao"
	"product-ms/internal/handler"
	"product-ms/internal/infras"
	"product-ms/internal/infras/config"
	"product-ms/internal/router"
	"product-ms/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("Could not load config due to " + err.Error())
	}

	fmt.Printf("configuration: %+v\n", cfg)

	cleanup, err := otel.SetupOpenTelemetry(cfg.AppName, cfg.AppVersion)
	if err != nil {
		panic("Could not setup open telemetry: " + err.Error())
	}
	defer cleanup()

	infra, err := infras.Load(cfg)
	if err != nil {
		panic("Could not load infrastructure due to " + err.Error())
	}

	queries := dao.New(infra.DB)

	// Initialize services
	productService := service.NewProductService(queries)
	storeService := service.NewStoreService(queries)
	discountService := service.NewDiscountService(queries)
	voucherService := service.NewVoucherService(queries)

	// Initialize handlers
	storeHandler := handler.NewStore(storeService)
	productHandler := handler.NewProduct(productService)
	discountHandler := handler.NewDiscount(discountService)
	voucherHandler := handler.NewVoucher(voucherService)

	route := router.NewRouter(cfg, storeHandler, productHandler, discountHandler, voucherHandler)

	server := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      route,
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		println("Server forced to shutdown " + err.Error())
	}
	println("Server exiting")
}
