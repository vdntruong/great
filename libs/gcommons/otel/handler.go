package otel

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func RollDice(w http.ResponseWriter, r *http.Request) {
	ctx, span := GetTracer().Start(r.Context(), "RollDice")
	defer span.End()

	_, span2 := GetTracer().Start(ctx, "random number")
	roll := 1 + rand.Intn(6)
	defer span2.End()

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
