package json

import (
	"encoding/json"
	"net/http"
)

const (
	HeaderContentType = "Content-Type"

	ContentTypeApplicationJSONUTF8 = "application/json; charset=utf-8"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondOK(w http.ResponseWriter, data interface{}) {
	responseJSON(w, http.StatusOK, data)
}

func RespondCreated(w http.ResponseWriter, data interface{}) {
	responseJSON(w, http.StatusCreated, data)
}

func RespondNotFoundError(w http.ResponseWriter, err error) {
	responseJSON(w, http.StatusNotFound, ErrorResponse{Error: err.Error()})
}

func RespondBadRequestError(w http.ResponseWriter, err error) {
	responseJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
}

func RespondInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func RespondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func responseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSONUTF8)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
