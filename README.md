# the great

## Development

Local env:
- [Traefik](http://localhost:8080/dashboard) as API Gateway.
- [Jaeger](http://localhost:16686/search) as tracing tool.
- [Prometheus](http://localhost:9090/query) as monitoring and alerting tool.
- [Grafana](http://localhost:3000) as visualization tool.

M.E.L.T ~ Metrics.Event.Logs.Traces ~ [MELT 101](https://newrelic.com/sites/default/files/2022-03/melt-101-four-essential-telemetry-data-types.pdf).

Instrumentation
- Source instrumentation
    - FE: Grafana Faro
    - BE: OpenTelemetry (the industry standard)
- Binary instrumentation
    - BE: Otel Collector (the agent sit one the component and then collect the information)
    - BE: Grafana Alloy (for Grafana stacks)
- External instrumentation
    - BE: Beyla for EBPF

Distributed tracing, also called distributed request tracing is a method used to debug and monitor applications built using a microservices architecture.

- What is tracing? span?
- What is context?
  - Span context?
    - Trace ID
    - Span ID
    - Trace flags
    - Trace state
  - Correlation context?
    - Customer ID
    - Host name
    - Region
- What is propagation? 
  - Propagation is the mechanism we use to bundle up our context and transfer across services.


### Integration Flow

1. Services emit telemetry data
2. OpenTelemetry Collector receives and processes data
3. Prometheus stores metrics
4. Jaeger stores and indexes traces
5. Loki stores logs
6. Grafana visualizes collected data

### OpenTelemetry Collector

- Role: Telemetry Data Collection & Processing
- Responsibilities:
  - Gather traces, metrics, and logs
  - Standardize telemetry data formats
  - Route and export data to multiple backends
  - Support multiple ingestion protocols
  - Provide unified observability pipeline
- [Collector setup](https://opentelemetry.io/docs/collector/quick-start/)
- [Collector installation](https://opentelemetry.io/docs/collector/installation/)
- [Collector configuration](https://opentelemetry.io/docs/collector/configuration/)
- [Docker compose samples](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/examples)

The Otel Collector's pipeline has 3 steps:

Receivers -> Processors -> Exporters

It collects telemetry data via the receivers and applies transformations in the processors stage before sending it to various outputs via exporters.

We can have pipelines configured, one for traces, one for metrics, and one for logs... 

Let's take a look at this example:

```yaml
service:
  pipelines:
    metrics:
      receivers: [otlp, opencensus, prometheus]
      processors: [batch]
      exporters: [logging]
```

Basically, the metrics pipeline can receive metrics via OTLP, OpenCensus, and Prometheus, 
and the batch processor batches together several metrics before sending them to the logging exporters.

### Prometheus

- Role: Metrics collection and storage
- Responsibilities:
  - Scrape metrics from services and infrastructure
  - Store time-series data
  - Support querying and alerting based on collected metrics
  - Provide pull-based monitoring mechanism
- [Installation](https://prometheus.io/docs/prometheus/latest/installation/)
- [How Prometheus monitoring works?](https://youtu.be/h4Sl21AKiDg?list=LL)

### Jaeger

- Role: Stores and indexes traces, provide searchable trace visualization
- Responsibilities:
    - Collect and visualize distributed traces
    - Track request flow across microservices
    - Performance bottleneck identification
    - Service dependency mapping
- [Installation](https://www.jaegertracing.io/docs/1.6/getting-started/)
- Alternatives: 
  - [Zipkin](https://zipkin.io/) is a distributed tracing system.

### Loki

- Role: Log Aggregation & Management
- Responsibilities:
  - Collect and store log data from multiple sources
  - Support distributed log storage
  - Enable log searching and filtering
  - Integrate with Grafana for visualization
- [Installation](https://grafana.com/docs/loki/latest/setup/install/docker/)

### Grafana

- Role: Visualization & Dashboarding
- Responsibilities:
  - Create interactive dashboards
  - Visualize metrics from Prometheus
  - Display log data from Loki
  - Support alerting and notification mechanisms
  - Provide multi-datasource visualization
- [Installation](https://grafana.com/docs/grafana/latest/setup-grafana/installation/docker/)


### DRAFT
this project was created with intention of researching and practicing about microservices.
I don't know now far I can go with this :)

Microservices
- Services
- APIs
- Service Discovery
- Load Balancer
- Service Mesh
- Message Queue
- Databases

Programing languages
- Go
- Python

API tool
- Bruno

Design architectures / patterns:
- Clean architecture
- Domain driven design
    - Domain
    - Sub domain
    - Aggregate
    - Entity
- Event driven architecture
- Event sourcing architecture
- CQRS (Command Query Responsibility Segregation)
- Circuit Breaker

API Gateway in Red zone / Internet Facing zone / Public network
- Single entry point for FE, routing
- Security implementation
- Policies, rate limit implementations
- Protocol translation (REST -> Gateway -> gRPC / SOAP / GraphQL...)

Service registry & service discovery & Config server / Remote config
- https://traefik.io/glossary/service-discovery/

Services, containers in Green zone / Private network

Service composition

Database challenges
- Sharding
- Indexing
- Partitioning
- Backup
- Validation data in microservices (validate keys)

Caching
- Redis

Deployment
- Container
- What is Cloud native?
- What is Service Mesh?

CI/CD
- Jenkins
- GitHub Actions

Logging, Monitoring & Tracing => Centralized logging & monitoring (log/metric agent push logs to Log Ingest)
- Open-telemetry
- Grafana
- Prometheus / Loki
- Correlation ID

Check these out:
"High traffic" have to apply Event Driven Architecture
How Order service check Product inventory -> Service composition

Testing
- Unit test
- Benchmark
- Stress test

Data Visualization
- Superset
- ETL


How to handle transactions in microservices / distributed computing?
- distributed transaction
    - 2-phase commit
    - saga go with event sourcing to perform rollback
- eventually consistency
