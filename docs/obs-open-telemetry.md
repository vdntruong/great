# OpenTelemetry

[OpenTelemetry Course](https://youtu.be/r8UvWSX3KA8) - freeCodeCamp

Keywords:
- Telemetry
- Instrumentation

We'll use OpenTelemetry to instrument, generate, collect and export telemetry data.

# Otel Collector

## [configuration structure](https://opentelemetry.io/docs/collector/configuration)

We have four classes of pipeline components that access telemetry data:
- **Receivers**
- **Processors**
- **Exporters**
- **Connectors**

Besides pipeline components you can also configure **extensions**, which provide capabilities that can be added to the Collector, 
such as diagnostic tools.

Sample configuration file:
```yaml
receivers:
  ...
processors:
  ...
exporters:
  ...
extensions:
  ...
connectors:
  ...
service:
  pipelines:
    ...
```

> [!NOTE] 
> 
> **Receivers**, **processors**, **exporters** and **pipelines** are defined through component identifiers
> following the `type[/name]` format.
> 
> For example: `otlp` or `otlp/2` or `otlp/jaeger`.

## [Receivers](https://github.com/open-telemetry/opentelemetry-collector/blob/main/receiver/README.md)

Receivers collect telemetry from one or more sources. 
They can be pull or push based, and may support one or more [data sources](https://opentelemetry.io/docs/concepts/signals/) (traces, metrics, logs, baggage).

> [!NOTE]
> Configuring a receiver does not enable it. Receivers are enabled by adding them to the appropriate pipelines within the service section.

Example:
```yaml
receivers:
  otlp: # Data sources: traces, metrics, logs
    ...
  kafka: # Data sources: traces, metrics, logs
    ...
  jaeger: # Data sources: traces
    ...
  zipkin: # Data sources: traces
    ...
  opencensus: # Data sources: traces, metrics
    ...
  prometheus: # Data sources: metrics
    ...
  hostmetrics: # Data sources: metrics
    ...
  fluentforward: # data source: logs
    ...
```

## [Processors](https://github.com/open-telemetry/opentelemetry-collector/blob/main/processor/README.md)

Processors take the data collected by receivers and modify or transform it before sending it to the exporters.
Data processing happens according to rules or settings defined for each processor, which might include filtering, dropping, renaming, or recalculating telemetry, among other operations.

The order of the processors in a pipeline determines the order of the processing operations that the Collector applies to the signal.

Example:
```yaml
processors:
  batch: # Data sources: traces, metrics, logs
  filter: # Data sources: metrics, metrics, logs
    error_mode: ignore
    traces:
      span:
        - 'attributes["container.name"] == "app_container_1"'
        - 'resource.attributes["host.name"] == "localhost"'
        - 'name == "app_3"'
      spanevent:
        - 'attributes["grpc"] == true'
        - 'IsMatch(name, ".*grpc.*")'
    metrics:
      metric:
        - 'name == "my.metric" and resource.attributes["my_label"] == "abc123"'
        - 'type == METRIC_DATA_TYPE_HISTOGRAM'
      datapoint:
        - 'metric.type == METRIC_DATA_TYPE_SUMMARY'
        - 'resource.attributes["service.name"] == "my_service_name"'
    logs:
      log_record:
        - 'IsMatch(body, ".*password.*")'
        - 'severity_number < SEVERITY_NUMBER_WARN'
  resource: # Data sources: traces
    attributes:
      - key: cloud.zone
        value: zone-1
        action: upsert
      - key: k8s.cluster.name
        from_attribute: k8s-cluster
        action: insert
      - key: redundant-attribute
        action: delete
  attributes: # Data sources: traces
    actions:
      - key: environment
        value: production
        action: insert
      - key: db.statement
        action: delete
      - key: email
        action: hash
```

## Exporters

Exporters send data to one or more backends or destinations. Exporters can be pull or push based, and may support one or more [data sources](https://opentelemetry.io/docs/concepts/signals/).

Each key within the exporters section defines an exporter instance. The key follows the `type/name` format, where `type` specifies the exporter type (e.g., `otlp`, `kafka`, `prometheus`), and `name` (optional) can be appended to provide a unique name for multiple instance of the same type.

Example:
```yaml
exporters:
  debug: # Data sources: traces, metrics, logs
    verbosity: detailed
  file: # Data sources: traces, metrics, logs
    path: ./filename.json
  otlp: # Data sources: traces, metrics, logs
    endpoint: otelcol2:4317
    tls:
      cert_file: cert.pem
      key_file: cert-key.pem
  kafka: # Data sources: traces, metrics, logs
    protocol_version: 2.0.0
  
  opencensus:  # Data sources: traces, metrics
    endpoint: otelcol2:55678
  otlphttp: # Data sources: traces, metrics
    endpoint: https://otlp.example.com:4318
  
  zipkin: # Data sources: traces
    endpoint: http://zipkin.example.com:9411/api/v2/spans
  otlp/jaeger: # Data sources: traces
    endpoint: jaeger-server:4317
    tls:
      cert_file: cert.pem
      key_file: cert-key.pem
            
  prometheus: # Data sources: metrics
    endpoint: 0.0.0.0:8889
    namespace: default
  prometheusremotewrite: # Data sources: metrics
    endpoint: http://prometheus.example.com:9411/api/prom/push
```

## [Connectors](https://github.com/open-telemetry/opentelemetry-collector/blob/main/connector/README.md)

Connectors join two pipelines, acting as both exporter and receiver.

A connector consumes data as an exporter at the end of one pipeline and emits data as a receiver at the beginning of another pipeline.

Example:

The following example shows the `count` connector and how it’s configured in the `pipelines` section. 
Notice that the connector acts as an exporter for traces and as a receiver for metrics, connecting both pipelines:
```yaml
receivers:
  foo:

exporters:
  bar:

connectors:
  count:
    spanevents:
      my.prod.event.count:
        description: The number of span events from my prod environment.
        conditions:
          - 'attributes["env"] == "prod"'
          - 'name == "prodevent"'

service:
  pipelines:
    traces:
      receivers: [foo]
      exporters: [count]
    metrics:
      receivers: [count]
      exporters: [bar]
```

## [Extensions](https://github.com/open-telemetry/opentelemetry-collector/blob/main/extension/README.md)

Extensions are optional components that expand the capabilities of the Collector to accomplish tasks not directly involved with processing telemetry data.
For example, you can add extensions for Collector health monitoring, service discovery, or data forwarding, among others.

Example:
```yaml
extensions:
  health_check:
  pprof:
  zpages:
```

## Service section

The `service` section is used to configure what components are enabled in the Collector based on the configuration found in the receivers, processors, exporters, and extensions sections. 

If a component is configured, but not defined within the service section, then it’s not enabled.

Example:
```yaml
service:
  pipelines:
    metrics:
      receivers: [opencensus, prometheus]
      processors: [batch]
      exporters: [opencensus, prometheus]
    traces:
      receivers: [opencensus, jaeger]
      processors: [batch, memory_limiter]
      exporters: [opencensus, zipkin]
```

### Extensions

The `extensions` subsection consists of a list of desired extensions to be enabled.

### Pipelines

The `pipelines` subsection is where the pipelines are configured, which can be of the following types:
- `traces` collect and processes trace data.
- `metrics` collect and processes metric data.
- `logs` collect and processes log data.

> [!NOTE]
> Note that the order of processors dictates the order in which data is processed.
