package otel

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type HandlerFunc func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request))

func HandleFunc(mux *http.ServeMux) HandlerFunc {
	return func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}
}

func Handler(mux *http.ServeMux) http.Handler {
	return otelhttp.NewHandler(mux, "/")
}
