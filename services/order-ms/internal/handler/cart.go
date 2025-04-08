package handler

import (
	"encoding/json"
	"net/http"

	"order-ms/internal/dto"
	"order-ms/internal/models"
	"order-ms/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Cart struct {
	service   service.CartService
	validator service.CartValidator
}

func NewCartHandler(cartService service.CartService, validator service.CartValidator) *Cart {
	return &Cart{
		service:   cartService,
		validator: validator,
	}
}

func (h *Cart) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.validator.ValidateGetCart(userUUID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart, err := h.service.GetCart(r.Context(), userUUID)
	if err != nil {
		// If cart doesn't exist, create a new one
		cart, err = h.service.CreateCart(r.Context(), userUUID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(dto.ToCartResponse(cart))
}

func (h *Cart) AddToCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var req dto.AddToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.validator.ValidateAddItem(userUUID, &models.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart, err := h.service.GetCart(r.Context(), userUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item := &models.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := h.service.AddItem(r.Context(), cart.ID, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get updated cart
	cart, err = h.service.GetCart(r.Context(), userUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dto.ToCartResponse(cart))
}

func (h *Cart) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	cartUUID, err := uuid.Parse(cartID)
	if err != nil {
		http.Error(w, "invalid cart ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.validator.ValidateUpdateItem(cartUUID, &models.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Remove old item
	if err := h.service.RemoveItem(r.Context(), cartUUID, req.ProductID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add new item
	item := &models.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := h.service.AddItem(r.Context(), cartUUID, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get updated cart
	cart, err := h.service.GetCart(r.Context(), cartUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dto.ToCartResponse(cart))
}

func (h *Cart) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	productID := chi.URLParam(r, "productId")

	cartUUID, err := uuid.Parse(cartID)
	if err != nil {
		http.Error(w, "invalid cart ID", http.StatusBadRequest)
		return
	}

	productUUID, err := uuid.Parse(productID)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.validator.ValidateRemoveItem(cartUUID, productUUID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.RemoveItem(r.Context(), cartUUID, productUUID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Cart) ClearCart(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "id")
	cartUUID, err := uuid.Parse(cartID)
	if err != nil {
		http.Error(w, "invalid cart ID", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.validator.ValidateClearCart(cartUUID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get cart items
	cart, err := h.service.GetCart(r.Context(), cartUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Remove all items
	for _, item := range cart.Items {
		if err := h.service.RemoveItem(r.Context(), cartUUID, item.ProductID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Cart) Routes() chi.Router {
	r := chi.NewRouter()

	// Cart routes
	r.Get("/", h.GetCart)
	r.Post("/items", h.AddToCart)
	r.Put("/items/{id}", h.UpdateCartItem)
	r.Delete("/items/{id}", h.RemoveFromCart)
	r.Delete("/items", h.ClearCart)

	return r
}
