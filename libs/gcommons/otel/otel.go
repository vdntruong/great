package otel

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.21.0"
)

func SetupOpenTelemetry(serviceName string) (func(), error) {
	ctx := context.Background()

	// Trace Exporter
	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint("otel-collector:4317"), // TODO
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// Metric Exporter
	otlMetricExporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint("otel-collector:4317"), // TODO
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	//// Prometheus Metric Exporter
	//promExporter, err := prometheus.New()
	//if err != nil {
	//	return nil, err
	//}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion("v1.0.0"),
		),
	)
	if err != nil {
		return nil, err
	}

	// Trace Provider
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)

	// Metric Provider
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(otlMetricExporter),
		),
		//metric.WithReader(promExporter),
	)
	otel.SetMeterProvider(meterProvider)

	// Propagation
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	// Custom Metrics
	meter := meterProvider.Meter(serviceName)
	requestCounter, err := meter.Int64Counter("service_requests_total") // TODO
	if err != nil {
		return nil, err
	}

	errorCounter, err := meter.Int64Counter("service_errors_total") // TODO
	if err != nil {
		return nil, err
	}

	// Example of incrementing the counter
	requestCounter.Add(context.Background(), 5)
	errorCounter.Add(context.Background(), 0)

	return func() {
		tracerProvider.Shutdown(ctx)
		meterProvider.Shutdown(ctx)
	}, nil
}

func SetupOTelSDK(ctx context.Context, serviceName string) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	tracerProvider, err := newTraceProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	meterProvider, err := newMeterProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	loggerProvider, err := newLoggerProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	// Custom Metrics
	meter := meterProvider.Meter(serviceName)
	requestCounter, err := meter.Int64Counter("service_requests_total") // TODO
	if err != nil {
		return nil, err
	}

	errorCounter, err := meter.Int64Counter("service_errors_total") // TODO
	if err != nil {
		return nil, err
	}

	// Example of incrementing the counter
	requestCounter.Add(context.Background(), 5)
	errorCounter.Add(context.Background(), 0)

	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider() (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

func newMeterProvider() (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(metricExporter, metric.WithInterval(3*time.Second)),
		),
	)
	return meterProvider, nil
}

func newLoggerProvider() (*log.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}
