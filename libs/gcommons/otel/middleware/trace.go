package middleware

import (
	"context"
	"net/http"
	"time"
	
	"gcommons/otel"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

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
		crw := &customResponseWriter{ResponseWriter: w}

		startTime := time.Now()

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
		ctx, span := otel.GetTracer().Start(
			r.Context(),
			r.URL.Path,
			trace.WithAttributes(
				attribute.String("http.method", r.Method),
				attribute.String("http.path", r.URL.Path),
			),
		)
		defer span.End()

		header := r.Header
		extractTraceFieldFromHeader(ctx, header, "x-request-id", "request-id")
		extractTraceFieldFromHeader(ctx, header, "x-transaction-id", "transaction-id")
		extractTraceFieldFromHeader(ctx, header, "x-correlation-id", "correlation-id")

		extractTraceFieldFromHeader(ctx, header, "x-app-version", "app-version")
		extractTraceFieldFromHeader(ctx, header, "x-release-timestamp", "release-timestamp")

		next.ServeHTTP(w, r.WithContext(ctx))
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
