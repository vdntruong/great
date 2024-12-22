package app

import (
	"encoding/json"
	"net/http"
	"time"

	ghandler "gcommons/handler"

	"user-ms/internal/dto"
)

func (a *Application) ping(w http.ResponseWriter, r *http.Request) {
	ghandler.HealthCheck(time.Now(), a.cfg.AppName)(w, r)
}

func (a *Application) register(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//res, err := a.userSvc.RegisterUser(r.Context(), req)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusCreated)
	//if err := json.NewEncoder(w).Encode(res); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
}
