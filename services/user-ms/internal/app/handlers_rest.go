package app

import (
	"net/http"

	gpassword "gcommons/password"

	"user-ms/internal/dto"
	"user-ms/internal/pkg/apperror"
)

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CreateUserDTO
	if err := app.Decode(r, &req); err != nil {
		app.badRequest(w, err.Error())
		return
	}

	_, founded, err := app.dao.GetByEmail(ctx, req.Email)
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

	user, err := app.dao.Insert(ctx, req.Email, req.Username, passwordHash)
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
