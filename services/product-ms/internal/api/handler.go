package api

import (
	"errors"
	"net/http"
	"strconv"

	wjson "commons/http/json"

	"product-ms/internal/dao"
	"product-ms/internal/dto"

	chi "github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductSvc *ProductSvc
}

func NewProductHandler(svc *ProductSvc) *ProductHandler {
	return &ProductHandler{
		ProductSvc: svc,
	}
}

func (h *ProductHandler) RegisterRoutes(r chi.Router) {
	r.Post("/products", h.HandleCreateProduct)
	r.Get("/products/{id}", h.HandleGetProductByID)
	r.Get("/products", h.HandleListProducts)
	r.Put("/products/{id}", h.HandleUpdateProduct)
	r.Delete("/products/{id}", h.HandleDeleteProduct)
}

func (h *ProductHandler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatingProductRequest
	if err := wjson.DecodeRequest(r, &req); err != nil {
		wjson.RespondBadRequestError(w, errors.New("invalid request body"))
		return
	}

	product, err := h.ProductSvc.ProductRepo.DAO.CreateProduct(r.Context(), dao.CreateProductParams{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		wjson.RespondInternalServerError(w, errors.New("failed to create product"))
		return
	}

	wjson.RespondCreated(w, product)
}

func (h *ProductHandler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wjson.RespondBadRequestError(w, errors.New("invalid id"))
		return
	}

	var req dto.UpdatingProductReq
	if err := wjson.DecodeRequest(r, &req); err != nil {
		wjson.RespondBadRequestError(w, errors.New("invalid request body"))
		return
	}

	product, err := h.ProductSvc.ProductRepo.DAO.UpdateProduct(r.Context(), dao.UpdateProductParams{
		ID:          int32(id),
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		wjson.RespondInternalServerError(w, errors.New("failed to update product"))
		return
	}

	wjson.RespondOK(w, product)
}

func (h *ProductHandler) HandleGetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wjson.RespondBadRequestError(w, errors.New("invalid id"))
		return
	}

	product, err := h.ProductSvc.ProductRepo.DAO.GetProductByID(r.Context(), int32(id))
	if err != nil {
		wjson.RespondNotFoundError(w, errors.New("product not found"))
		return
	}

	wjson.RespondOK(w, product)
}

func (h *ProductHandler) HandleListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductSvc.ProductRepo.DAO.ListProducts(r.Context())
	if err != nil {
		wjson.RespondInternalServerError(w, err)
		return
	}

	wjson.RespondOK(w, products)
}

func (h *ProductHandler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wjson.RespondBadRequestError(w, errors.New("invalid id"))
		return
	}

	if err := h.ProductSvc.ProductRepo.DAO.DeleteProduct(r.Context(), int32(id)); err != nil {
		wjson.RespondInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
