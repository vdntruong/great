package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"user-ms/internal/pkg/apperror"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	app.logger.Error().Err(err).Caller().Msg("server error")

	// TODO: improve errors to they have them own response process

	if errors.Is(err, apperror.ErrEmailExisted) {
		http.Error(w, apperror.ErrEmailExisted.Error(), http.StatusInternalServerError)
		return
	}

	if errors.Is(err, apperror.ErrUsernameExisted) {
		http.Error(w, apperror.ErrUsernameExisted.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) badRequest(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusBadRequest)
}

func (app *Application) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *Application) conflict(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusConflict)
}

func (app *Application) created(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		app.serverError(w, err)
	}
}
