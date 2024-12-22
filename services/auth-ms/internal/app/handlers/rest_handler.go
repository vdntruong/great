package handlers

import (
	"encoding/json"
	"net/http"

	"auth-ms/internal/app/entities/dtos"
)

type AuthRESTHandler struct {
	userSvc IUserService
}

func NewAuthRESTHandler(userSvc IUserService) *AuthRESTHandler {
	return &AuthRESTHandler{
		userSvc: userSvc,
	}
}

func (h *AuthRESTHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.userSvc.RegisterUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//func (h *AuthRESTHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
//	var credentials models.Credentials
//	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	user, err := h.userSvc.FindByUsername(r.Context(), credentials.Username)
//	if err != nil || !password.CheckPasswordHash(credentials.Password, user.Password) {
//		http.Error(w, "invalid username or password", http.StatusUnauthorized)
//		return
//	}
//
//	tokenPair, err := authen.GenerateTokenPair(user.ID, user.Username)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(tokenPair)
//}
