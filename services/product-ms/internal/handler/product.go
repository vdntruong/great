package handler

import (
	"net/http"
	"strconv"

	commonjson "commons/http/json"

	"product-ms/internal/dto"
	"product-ms/internal/models"
	"product-ms/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Product struct {
	service service.ProductService
}

func NewProduct(service service.ProductService) *Product {
	return &Product{
		service: service,
	}
}

// RegisterRoutes registers all product routes
func (h *Product) RegisterRoutes(r chi.Router) {
	r.Route("/stores/{store_id}/products", func(r chi.Router) {
		r.Post("/", h.HandleCreate)
		r.Get("/", h.HandleList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.HandleGet)
			r.Put("/", h.HandleUpdate)
			r.Delete("/", h.HandleDelete)
		})
	})
}

// HandleCreate handles the creation of a new product
func (h *Product) HandleCreate(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "store_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	var req dto.CreateProductRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	params := ConvertCreateProductRequestToModel(&req)
	params.StoreID = storeID

	product, err := h.service.CreateProduct(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondCreated(w, ConvertProductModelToResponse(product))
}

// HandleGet handles retrieving a product by ID
func (h *Product) HandleGet(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	product, err := h.service.GetProductByID(r.Context(), productID.String())
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondOK(w, ConvertProductModelToResponse(product))
}

// HandleList handles retrieving a list of products
func (h *Product) HandleList(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "store_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	products, err := h.service.ListProducts(r.Context(), models.ListProductsParams{
		StoreID: storeID,
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	response := dto.ListProductsResponse{
		Products: make([]dto.ProductResponse, len(products)),
		Total:    int64(len(products)),
	}

	for i, product := range products {
		response.Products[i] = ConvertProductModelToResponse(product)
	}

	commonjson.RespondOK(w, response)
}

// HandleUpdate handles updating a product
func (h *Product) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	var req dto.UpdateProductRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	params := ConvertUpdateProductRequestToModel(&req)
	params.ID = productID

	product, err := h.service.UpdateProduct(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondOK(w, ConvertProductModelToResponse(product))
}

// HandleDelete handles deleting a product
func (h *Product) HandleDelete(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	if err := h.service.DeleteProduct(r.Context(), productID.String()); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
