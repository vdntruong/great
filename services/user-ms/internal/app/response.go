package app

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/hlog"
	"net/http"

	"commons/errs"
)

func (app *Application) respondError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		ae errs.Error
		ok = errors.As(err, &ae)
	)

	if ok {
		http.Error(w, ae.GetStatusText(), ae.GetStatusCode())
		return
	}

	hlog.FromRequest(r).Error().Err(err).Msg("respond error")
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) respondBadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func (app *Application) respondNotFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *Application) respondConflict(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusConflict)
}

func (app *Application) respondCreated(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *Application) respondOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
