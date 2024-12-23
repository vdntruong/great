package app

import (
	"encoding/json"
	"net/http"
	"time"
	"user-ms/internal/pkg/apperror"

	ghandler "gcommons/handler"
	gpassword "gcommons/password"

	"user-ms/internal/dto"
)

func (app *Application) ping(w http.ResponseWriter, r *http.Request) {
	ghandler.HealthCheck(time.Now(), app.cfg.AppName)(w, r)
}

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		app.badRequest(w, err.Error())
		return
	}

	ctx := r.Context()

	_, founded, err := app.users.GetByEmail(ctx, req.Email)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if founded {
		app.conflict(w, apperror.ErrEmailExisted.Error())
		return
	}

	passwordHash, err := gpassword.Hash(req.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	user, err := app.users.Insert(ctx, req.Email, req.Username, passwordHash)
	if err != nil {
		app.serverError(w, err)
		return
	}

	resp := dto.UserDTO{
		Email:    user.Email,
		Username: user.Username,
	}
	app.created(w, resp)
}
