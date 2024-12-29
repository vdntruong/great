package otel

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel/metric"
)

var (
	// HTTP Basic Metrics
	requestCounter  metric.Int64Counter
	requestDuration metric.Float64Histogram
	activeRequests  metric.Int64UpDownCounter
	responseSize    metric.Int64Histogram
	requestSize     metric.Float64Histogram
)

func initMetricsHTTPRequestCounter() (err error) {
	requestCounter, err = meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("1"),
	)
	return
}

func GetRequestCountMeter() metric.Int64Counter {
	return requestCounter
}

func initMetricsHTTPRequestDuration() (err error) {
	requestDuration, err = meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("HTTP request duration in seconds"),
		metric.WithUnit("s"),
	)
	return
}

func GetRequestDurationMeter() metric.Float64Histogram {
	return requestDuration
}

func initMetricsHTTPActiveRequest() (err error) {
	activeRequests, err = meter.Int64UpDownCounter(
		"http_active_requests",
		metric.WithDescription("Number of active HTTP requests"),
		metric.WithUnit("1"),
	)
	return
}

func GetRequestActiveMeter() metric.Int64UpDownCounter {
	return activeRequests
}

func initMetricsHTTPResponseSize() (err error) {
	responseSize, err = meter.Int64Histogram(
		"http_response_size_bytes",
		metric.WithDescription("Size of HTTP responses in bytes"),
		metric.WithUnit("bytes"),
	)
	return
}

func GetResponseSizeMeter() metric.Int64Histogram {
	return responseSize
}

func initMetricsHTTPRequestSize() (err error) {
	requestSize, err = meter.Float64Histogram(
		"http_request_size_bytes",
		metric.WithDescription("Size of HTTP requests in bytes"),
		metric.WithUnit("bytes"),
	)
	return
}

func GetRequestSizeMeter() metric.Float64Histogram {
	return requestSize
}

var (
	// Business Metrics
	requestErrors  metric.Int64Counter
	requestsByPath metric.Int64Counter

	// Performance Metrics
	requestQueueTime metric.Float64Histogram
)

func initMetricsRequestErrors() (err error) {
	requestErrors, err = meter.Int64Counter(
		"http_request_errors_total",
		metric.WithDescription("Total number of HTTP request errors"),
		metric.WithUnit("1"),
	)
	return
}

func GetRequestErrorsMeter() metric.Int64Counter {
	return requestErrors
}

func initMetricsRequestsByPath() (err error) {
	requestsByPath, err = meter.Int64Counter(
		"http_requests_by_path_total",
		metric.WithDescription("Total requests broken down by path"),
		metric.WithUnit("1"),
	)
	return
}

func GetRequestsByPathMeter() metric.Int64Counter {
	return requestsByPath
}

func initMetricsRequestQueueTime() (err error) {
	requestQueueTime, err = meter.Float64Histogram(
		"http_request_queue_duration_seconds",
		metric.WithDescription("Time spent in request queue"),
		metric.WithUnit("s"),
	)
	return
}

func GetRequestQueueTimeMeter() metric.Float64Histogram {
	return requestQueueTime
}

func InitHTTPMetrics() error {
	if err := initMetricsHTTPRequestCounter(); err != nil {
		return err
	}
	if err := initMetricsHTTPRequestDuration(); err != nil {
		return err
	}
	if err := initMetricsHTTPActiveRequest(); err != nil {
		return err
	}
	if err := initMetricsHTTPResponseSize(); err != nil {
		return err
	}
	if err := initMetricsHTTPRequestSize(); err != nil {
		return err
	}

	if err := initMetricsRequestErrors(); err != nil {
		return err
	}
	if err := initMetricsRequestsByPath(); err != nil {
		return err
	}
	if err := initMetricsRequestQueueTime(); err != nil {
		return err
	}
	return nil
}

var (
	// System Metrics
	memoryUsage    metric.Float64ObservableGauge
	goroutineCount metric.Int64ObservableGauge

	// Performance Metrics
	cpuUsage metric.Float64ObservableGauge
)

func initMetricsMemoryUsage() (err error) {
	memoryUsage, err = meter.Float64ObservableGauge(
		"memory_usage_bytes",
		metric.WithDescription("Current memory usage in bytes"),
		metric.WithUnit("bytes"),
	)
	if err != nil {
		return
	}
	return
}

func initMetricsGoroutineCount() (err error) {
	goroutineCount, err = meter.Int64ObservableGauge(
		"goroutines_count",
		metric.WithDescription("Number of running goroutines"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return
	}
	return
}

func initMetricsCPUUsage() (err error) {
	cpuUsage, err = meter.Float64ObservableGauge(
		"cpu_usage_percentage",
		metric.WithDescription("CPU usage percentage"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return
	}
	return
}

func InitSystemMetricsAndCallbacks() error {
	if err := initMetricsCPUUsage(); err != nil {
		return err
	}
	if err := initMetricsGoroutineCount(); err != nil {
		return err
	}
	if err := initMetricsMemoryUsage(); err != nil {
		return err
	}

	if _, err := meter.RegisterCallback(
		func(_ context.Context, observer metric.Observer) error {
			observer.ObserveFloat64(cpuUsage, getCPUUsage())
			observer.ObserveFloat64(memoryUsage, getMemoryUsage())
			observer.ObserveInt64(goroutineCount, int64(runtime.NumGoroutine()))
			return nil
		},
		cpuUsage,
		memoryUsage,
		goroutineCount,
	); err != nil {
		return err
	}
	return nil
}
