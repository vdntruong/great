package handler

import (
	"net/http"
	"strconv"

	"product-ms/internal/dto"
	"product-ms/internal/models"
	"product-ms/internal/service"

	commonjson "commons/http/json"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Store struct {
	storeService service.StoreService
}

func NewStore(storeService service.StoreService) *Store {
	return &Store{
		storeService: storeService,
	}
}

// RegisterRoutes registers all store routes
func (h *Store) RegisterRoutes(r chi.Router) {
	r.Route("/stores", func(r chi.Router) {
		r.Post("/", h.HandleCreate)
		r.Get("/", h.HandleList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.HandleGet)
			r.Put("/", h.HandleUpdate)
			r.Delete("/", h.HandleDelete)
		})
	})
}

// HandleCreate handles the creation of a new store
func (h *Store) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateStoreRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	params := ConvertCreateStoreRequestToModel(&req)
	store, err := h.storeService.CreateStore(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondCreated(w, ConvertStoreModelToResponse(store))
}

// HandleGet handles retrieving a store by ID
func (h *Store) HandleGet(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	store, err := h.storeService.GetStoreByID(r.Context(), storeID.String())
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondOK(w, ConvertStoreModelToResponse(store))
}

// HandleList handles retrieving a list of stores
func (h *Store) HandleList(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	params := models.ListStoresParams{
		Limit:  int32(limit),
		Offset: int32((page - 1) * limit),
	}

	stores, err := h.storeService.ListStores(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	response := dto.StoreListResponse{
		Stores:     make([]dto.StoreResponse, len(stores.Stores)),
		TotalCount: stores.TotalCount,
		Page:       stores.Page,
		Limit:      stores.Limit,
	}

	for i, store := range stores.Stores {
		response.Stores[i] = ConvertStoreModelToResponse(&store)
	}

	commonjson.RespondOK(w, response)
}

// HandleUpdate handles updating a store
func (h *Store) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	var req dto.UpdateStoreRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	params := ConvertUpdateStoreRequestToModel(&req)
	store, err := h.storeService.UpdateStore(r.Context(), storeID.String(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondOK(w, ConvertStoreModelToResponse(store))
}

// HandleDelete handles deleting a store
func (h *Store) HandleDelete(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	if err := h.storeService.DeleteStore(r.Context(), storeID.String()); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}
