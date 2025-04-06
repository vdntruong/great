package router

import (
	"product-ms/internal/handler"
	"product-ms/internal/service"

	"github.com/go-chi/chi/v5"
)

type ProductRouter struct {
	productHandler *handler.ProductHandler
}

func NewProductRouter(productService service.ProductService) *ProductRouter {
	return &ProductRouter{
		productHandler: handler.NewProductHandler(productService),
	}
}

func (r *ProductRouter) RegisterRoutes(router chi.Router) {
	r.productHandler.RegisterRoutes(router)
}