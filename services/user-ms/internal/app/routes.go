package app

import (
	"net/http"
	"time"

	ghandler "commons/handler"
	gmiddleware "commons/middleware"
	otelmiddleware "commons/otel/middleware"

	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
)

func (app *Application) Routes() http.Handler {
	standardMiddleware := alice.New(
		gmiddleware.RecoverPanic,
		otelmiddleware.Metrics,
		otelmiddleware.TraceRequest,
		// https://github.com/rs/zerolog?tab=readme-ov-file#integration-with-nethttp
		hlog.NewHandler(app.logger),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("")
		}),
		hlog.RemoteAddrHandler("ip"),
		hlog.UserAgentHandler("user_agent"),
		hlog.RefererHandler("referer"),
		hlog.RequestIDHandler("request_id", "X-Request-Id"),
	)

	mux := http.NewServeMux()

	{ // root
		mux.HandleFunc("GET /healthz", ghandler.HealthCheck(time.Now(), app.cfg.AppName))
		mux.HandleFunc("GET /roll-dice", ghandler.RollDice)
	}

	muxV1 := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", muxV1))

	{ // api/v1/*
		{ // users
			muxV1.HandleFunc("POST /users/register", app.HandleRegister)
			muxV1.HandleFunc("GET /users/profile", app.HandleProfile)
		}
	}

	return standardMiddleware.Then(mux)
}
