package router

import (
	"net/http"

	"product-ms/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	storeHandler    *handler.Store
	productHandler  *handler.Product
	discountHandler *handler.Discount
	voucherHandler  *handler.Voucher
}

func NewRouter(
	storeHandler *handler.Store,
	productHandler *handler.Product,
	discountHandler *handler.Discount,
	voucherHandler *handler.Voucher,
) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Store routes
	r.Route("/stores", func(r chi.Router) {
		r.Post("/", storeHandler.HandleCreate)
		r.Get("/", storeHandler.HandleList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", storeHandler.HandleGet)
			r.Put("/", storeHandler.HandleUpdate)
			r.Delete("/", storeHandler.HandleDelete)
		})
	})

	// Product routes
	r.Route("/stores/{store_id}/products", func(r chi.Router) {
		r.Post("/", productHandler.HandleCreate)
		r.Get("/", productHandler.HandleList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", productHandler.HandleGet)
			r.Put("/", productHandler.HandleUpdate)
			r.Delete("/", productHandler.HandleDelete)
		})
	})

	// Discount routes
	r.Route("/stores/{store_id}/discounts", func(r chi.Router) {
		r.Post("/", discountHandler.CreateDiscount)
		r.Get("/", discountHandler.ListDiscounts)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", discountHandler.GetDiscount)
			r.Put("/", discountHandler.UpdateDiscount)
			r.Delete("/", discountHandler.DeleteDiscount)
			r.Post("/products/{product_id}", discountHandler.AddDiscountProduct)
			r.Delete("/products/{product_id}", discountHandler.RemoveDiscountProduct)
			r.Post("/categories/{category_id}", discountHandler.AddDiscountCategory)
			r.Delete("/categories/{category_id}", discountHandler.RemoveDiscountCategory)
		})
	})

	// Voucher routes
	r.Route("/stores/{store_id}/vouchers", func(r chi.Router) {
		r.Post("/", voucherHandler.CreateVoucher)
		r.Get("/", voucherHandler.ListVouchers)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", voucherHandler.GetVoucher)
			r.Put("/", voucherHandler.UpdateVoucher)
			r.Delete("/", voucherHandler.DeleteVoucher)
			r.Put("/status", voucherHandler.UpdateVoucherStatus)
		})
	})

	return r
}
