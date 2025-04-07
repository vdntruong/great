package router

import (
	"product-ms/internal/handler"
	"product-ms/internal/service"

	"github.com/go-chi/chi/v5"
)

type StoreRouter struct {
	storeHandler *handler.Store
}

func NewStoreRouter(storeService service.StoreService) *StoreRouter {
	return &StoreRouter{
		storeHandler: handler.NewStore(storeService),
	}
}

func (r *StoreRouter) RegisterRoutes(router chi.Router) {
	r.storeHandler.RegisterRoutes(router)
}
