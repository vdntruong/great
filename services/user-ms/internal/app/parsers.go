package app

import (
	"encoding/json"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func (app *Application) Decode(r *http.Request, v interface{}) error {
	span := trace.SpanFromContext(r.Context())
	span.SetName("decode request body")
	defer span.End()

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}
