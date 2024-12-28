package app

import (
	"encoding/json"
	"net/http"
)

func (app *Application) Decode(r *http.Request, v interface{}) error {
	_, span := app.tracer.Start(r.Context(), "decode request body")
	defer span.End()

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}
