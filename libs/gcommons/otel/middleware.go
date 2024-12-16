package otel

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"net/http"
	"time"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	requestCounter, _ := GetMeter().Int64Counter("http_requests_total",
		metric.WithDescription("Total number of HTTP requests"))

	requestDuration, _ := GetMeter().Float64Histogram("http_request_duration_seconds",
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
