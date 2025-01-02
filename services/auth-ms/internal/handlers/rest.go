package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"commons/otel/trace"

	"auth-ms/internal/entities/dtos"
)

type AuthHandler struct {
	tracer  trace.Tracer
	authSvc IAuthService
}

func NewAuthHandler(tracer trace.Tracer, authSvc IAuthService) *AuthHandler {
	return &AuthHandler{
		tracer:  tracer,
		authSvc: authSvc,
	}
}

func (h *AuthHandler) HandleAccessToken(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) HandleAuthenticate(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenParts := strings.Split(authToken, " ")
	if len(tokenParts) != 2 {
		http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
		return
	}

	token := tokenParts[1]
	if err := h.authSvc.VerifyAuthToken(r.Context(), token); err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
