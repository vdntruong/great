package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"user-ms/constants"

	ghandler "gcommons/handler"
	"gcommons/otel"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cleanup, err := otel.SetupOpenTelemetry(constants.ServiceName, constants.ServiceVersion)
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry: %v", err)
	}
	defer cleanup()

	log.Println("Authentication service starting on :8080")
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case err = <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}

	err = srv.Shutdown(context.Background())
	return
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	otel.HandleFunc(mux)("/healthz", ghandler.HealthCheck(time.Now(), "User"))
	otel.HandleFunc(mux)("/rolldice/", otel.RollDice)

	handler := otel.Handler(mux)
	return handler
}

//func initTracer() (*trace.TracerProvider, error) {
//	ctx := context.Background()
//
//	exporter, err := otlptracehttp.New(ctx,
//		otlptracehttp.WithEndpoint("otel-collector:4318"),
//		otlptracehttp.WithInsecure(),
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	tp := trace.NewTracerProvider(
//		trace.WithBatcher(exporter),
//		trace.WithSampler(trace.AlwaysSample()),
//		trace.WithResource(resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceName("service-user"),
//		)),
//	)
//	otel.SetTracerProvider(tp)
//	otel.SetTextMapPropagator(propagation.TraceContext{})
//
//	return tp, nil
//}
