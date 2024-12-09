package handlers

import (
	"auth-ms/internal/app/models"
	"auth-ms/internal/app/repository"
	"encoding/json"
	"gcommons/authen"
	"gcommons/password"
	"net/http"
)

type AuthRESTHandler struct {
	userRepo repository.UserRepository
}

func NewAuthRESTHandler(userRepo repository.UserRepository) *AuthRESTHandler {
	return &AuthRESTHandler{
		userRepo: userRepo,
	}
}

func (h *AuthRESTHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userRepo.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *AuthRESTHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.FindByUsername(r.Context(), credentials.Username)
	if err != nil || !password.CheckPasswordHash(credentials.Password, user.Password) {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	tokenPair, err := authen.GenerateTokenPair(user.ID, user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenPair)
}
