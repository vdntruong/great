package otel

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel/metric"
)

var (
	// HTTP Metrics
	requestCounter  metric.Int64Counter
	requestDuration metric.Float64Histogram
	activeRequests  metric.Int64UpDownCounter
	responseSize    metric.Int64Histogram
)

func InitMetricsHTTPRequestCounter() (err error) {
	requestCounter, err = meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("1"),
	)
	return
}

func InitMetricsHTTPRequestDuration() (err error) {
	requestDuration, err = meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("HTTP request duration in seconds"),
		metric.WithUnit("s"),
	)
	return
}

func InitMetricsHTTPActiveRequest() (err error) {
	activeRequests, err = meter.Int64UpDownCounter(
		"http_active_requests",
		metric.WithDescription("Number of active HTTP requests"),
		metric.WithUnit("1"),
	)
	return
}

func InitMetricsHTTPResponseSize() (err error) {
	responseSize, err = meter.Int64Histogram(
		"http_response_size_bytes",
		metric.WithDescription("Size of HTTP responses in bytes"),
		metric.WithUnit("bytes"),
	)
	return
}

var (
	// System Metrics
	memoryUsage    metric.Float64ObservableGauge
	goroutineCount metric.Int64ObservableGauge
)

func InitMetricsMemoryUsage() (err error) {
	memoryUsage, err = meter.Float64ObservableGauge(
		"memory_usage_bytes",
		metric.WithDescription("Current memory usage in bytes"),
		metric.WithUnit("bytes"),
	)
	if err != nil {
		return
	}

	_, err = meter.RegisterCallback(
		func(_ context.Context, observer metric.Observer) error {
			observer.ObserveFloat64(memoryUsage, getMemoryUsage())
			return nil
		},
		memoryUsage,
	)

	return
}

func InitMetricsGoroutineCount() (err error) {
	goroutineCount, err = meter.Int64ObservableGauge(
		"goroutines_count",
		metric.WithDescription("Number of running goroutines"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return
	}

	_, err = meter.RegisterCallback(
		func(_ context.Context, observer metric.Observer) error {
			observer.ObserveInt64(goroutineCount, int64(runtime.NumGoroutine()))
			return nil
		},
		goroutineCount,
	)
	return
}

var (
	// Business Metrics
	requestErrors  metric.Int64Counter
	requestsByPath metric.Int64Counter
)

func InitMetricsRequestErrors() (err error) {
	requestErrors, err = meter.Int64Counter(
		"http_request_errors_total",
		metric.WithDescription("Total number of HTTP request errors"),
		metric.WithUnit("1"),
	)
	return
}

func InitMetricsRequestsByPath() (err error) {
	requestsByPath, err = meter.Int64Counter(
		"http_requests_by_path_total",
		metric.WithDescription("Total requests broken down by path"),
		metric.WithUnit("1"),
	)
	return
}

var (
	// Performance Metrics
	requestQueueTime metric.Float64Histogram
	cpuUsage         metric.Float64ObservableGauge
)

func InitMetricsRequestQueueTime() (err error) {
	requestQueueTime, err = meter.Float64Histogram(
		"http_request_queue_duration_seconds",
		metric.WithDescription("Time spent in request queue"),
		metric.WithUnit("s"),
	)
	return
}

func InitMetricsCPUUsage() (err error) {
	cpuUsage, err = meter.Float64ObservableGauge(
		"cpu_usage_percentage",
		metric.WithDescription("CPU usage percentage"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return
	}

	_, err = meter.RegisterCallback(
		func(_ context.Context, observer metric.Observer) error {
			observer.ObserveFloat64(cpuUsage, getCPUUsage())
			return nil
		},
		cpuUsage,
	)
	return
}

func initSystemMetrics() error {
	if err := InitMetricsCPUUsage(); err != nil {
		return err
	}
	if err := InitMetricsGoroutineCount(); err != nil {
		return err
	}
	if err := InitMetricsMemoryUsage(); err != nil {
		return err
	}
	return nil
}
