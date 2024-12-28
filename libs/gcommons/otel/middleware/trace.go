package middleware

//https://opentelemetry.io/docs/languages/go/instrumentation/

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gcommons/otel"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type tracerKey struct{}

func WithTracer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), tracerKey{}, otel.GetTracer())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

func Metrics(next http.Handler) http.Handler {
	requestCounter, _ := otel.GetMeter().Int64Counter("http_requests_total",
		metric.WithDescription("Total number of HTTP requests"))

	requestDuration, _ := otel.GetMeter().Float64Histogram("http_request_duration_seconds",
		metric.WithDescription("HTTP request duration"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		crw := &customResponseWriter{ResponseWriter: w}
		next.ServeHTTP(crw, r)

		duration := time.Since(startTime).Seconds()

		requestCounter.Add(r.Context(), 1,
			metric.WithAttributes(
				attribute.String("method", r.Method),
				attribute.String("path", r.URL.Path),
				attribute.Int("status", crw.statusCode),
			))

		requestDuration.Record(r.Context(), duration,
			metric.WithAttributes(
				attribute.String("method", r.Method),
				attribute.String("path", r.URL.Path),
				attribute.Int("status", crw.statusCode),
			))
	})
}

func TraceRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.GetTracer().Start(r.Context(), r.URL.Path,
			trace.WithAttributes(
				attribute.String("http.method", r.Method),
				attribute.String("http.path", r.URL.Path),
			),
		)
		defer span.End()

		header := r.Header
		extractTraceFieldFromHeader(ctx, header, "x-os-platform", "os-platform")
		extractTraceFieldFromHeader(ctx, header, "x-app-version", "app-version")
		extractTraceFieldFromHeader(ctx, header, "x-request-id", "request-id")
		extractTraceFieldFromHeader(ctx, header, "x-transaction-id", "transaction-id")
		extractTraceFieldFromHeader(ctx, header, "x-correlation-id", "correlation-id")

		crw := &customResponseWriter{ResponseWriter: w}
		next.ServeHTTP(crw, r.WithContext(ctx))

		span.SetAttributes(attribute.String("http.status_code", strconv.Itoa(crw.statusCode)))
	})
}

func extractTraceFieldFromHeader(
	ctx context.Context,
	header http.Header,
	headerKey string,
	attributeKey string,
) {
	if headerValue := header.Get(headerKey); headerValue != "" {
		trace.SpanFromContext(ctx).SetAttributes(attribute.String(attributeKey, headerValue))
	}
}
