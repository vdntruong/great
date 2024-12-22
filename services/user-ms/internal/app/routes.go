package app

import (
	"net/http"

	ghandler "gcommons/handler"
	gmiddleware "gcommons/middleware"
	otelmiddleware "gcommons/otel/middleware"

	"github.com/justinas/alice"
)

func (a *Application) Routes() http.Handler {
	standardMiddleware := alice.New(
		gmiddleware.RecoverPanic,
		gmiddleware.LogRequest,
		otelmiddleware.TraceRequest,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", a.ping)
	mux.HandleFunc("GET /rr", ghandler.RollDice)

	return standardMiddleware.Then(mux)
}
