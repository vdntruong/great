package middleware

//https://opentelemetry.io/docs/languages/go/instrumentation/

import (
	"context"
	"log"
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
	status int
	size   int64
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.status = code
	crw.ResponseWriter.WriteHeader(code)
}

func (crw *customResponseWriter) Write(b []byte) (int, error) {
	n, err := crw.ResponseWriter.Write(b)
	crw.size += int64(n)
	return n, err
}

func Metrics(next http.Handler) http.Handler {
	if err := otel.InitHTTPMetrics(); err != nil {
		log.Println("otel.InitHTTPMetrics err:", err.Error())
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	requestCounter := otel.GetRequestCountMeter()
	requestDuration := otel.GetRequestDurationMeter()
	activeRequests := otel.GetRequestActiveMeter()
	responseSize := otel.GetResponseSizeMeter()
	requestSize := otel.GetRequestSizeMeter()
	requestsByPath := otel.GetRequestsByPathMeter()
	requestErrors := otel.GetRequestErrorsMeter()
	requestQueueTime := otel.GetRequestQueueTimeMeter()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// method, and path of the request
		methodAttr, pathAttr := attribute.String("method", r.Method), attribute.String("path", r.URL.Path)
		baseAttrs := []attribute.KeyValue{methodAttr, pathAttr}
		baseAttrsOpt := metric.WithAttributes(baseAttrs...)

		requestsByPath.Add(r.Context(), 1, metric.WithAttributes(pathAttr))
		requestCounter.Add(r.Context(), 1, baseAttrsOpt)
		if size := float64(r.ContentLength); size > 0 {
			requestSize.Record(r.Context(), size, baseAttrsOpt)
		}
		activeRequests.Add(r.Context(), 1, baseAttrsOpt)
		defer activeRequests.Add(r.Context(), -1, baseAttrsOpt)

		queueTime := time.Since(startTime).Seconds()
		requestQueueTime.Record(r.Context(), queueTime, baseAttrsOpt)

		// next handler
		crw := &customResponseWriter{ResponseWriter: w}
		next.ServeHTTP(crw, r)

		// method, path, and response status of the request
		statusAttr := attribute.Int("status", crw.status)
		fullAttrsOpt := metric.WithAttributes(append(baseAttrs, statusAttr)...)

		if crw.status >= 400 {
			requestErrors.Add(r.Context(), 1, fullAttrsOpt)
		}
		responseSize.Record(r.Context(), crw.size, baseAttrsOpt)
		requestDuration.Record(r.Context(), time.Since(startTime).Seconds(), fullAttrsOpt)
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

		span.SetAttributes(attribute.String("http.status_code", strconv.Itoa(crw.status)))
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
