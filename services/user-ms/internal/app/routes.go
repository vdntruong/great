package app

import (
	ghandler "gcommons/handler"
	gmiddleware "gcommons/middleware"
	"github.com/justinas/alice"
	"net/http"
)

func (a *Application) Routes() http.Handler {
	standardMiddleware := alice.New(
		gmiddleware.RecoverPanic,
		gmiddleware.LogRequest,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", a.ping)
	mux.HandleFunc("GET /rr", ghandler.RollDice)

	return standardMiddleware.Then(mux)
}
