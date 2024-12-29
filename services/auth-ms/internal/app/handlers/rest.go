package handlers

import (
	"encoding/json"
	"net/http"

	"auth-ms/internal/app/entities/dtos"
)

type AuthHandler struct {
	authSvc IAuthService
}

func NewAuthHandler(authSvc IAuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.LoginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.authSvc.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
