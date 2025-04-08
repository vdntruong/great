package router

import (
	chandler "commons/handler"
	otelmiddleware "commons/otel/middleware"
	"order-ms/internal/infras/config"
	"time"

	"order-ms/internal/handler"
	"order-ms/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(
	cfg *config.Config,
	orderService service.OrderService,
	orderValidator service.OrderValidator,
	cartService service.CartService,
	cartValidator service.CartValidator,
) *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(otelmiddleware.Metrics)      // otel metrics
	r.Use(otelmiddleware.TraceRequest) // otel trace

	r.Use(middleware.Timeout(cfg.ReadTimeout + cfg.WriteTimeout))
	r.Get("/healthz", chandler.HealthCheck(time.Now(), cfg.AppName)) // health check traefik

	// APIv1 routes
	r.Route("/api/v1", func(r chi.Router) {
		// Order routes
		orderHandler := handler.NewOrderHandler(orderService, orderValidator)
		r.Mount("/orders", orderHandler.Routes())

		// Cart routes
		cartHandler := handler.NewCartHandler(cartService, cartValidator)
		r.Mount("/carts", cartHandler.Routes())
	})

	return r
}
