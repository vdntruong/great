## the great

## Monitoring tools

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

### Prometheus

- Role: Metrics collection and storage
- Responsibilities:
  - Scrape metrics from services and infrastructure
  - Store time-series data
  - Support querying and alerting based on collected metrics
  - Provide pull-based monitoring mechanism
- [Installation](https://prometheus.io/docs/prometheus/latest/installation/)

### Jaeger

- Role: Stores and indexes traces, provide searchable trace visualization
- Responsibilities:
    - Collect and visualize distributed traces
    - Track request flow across microservices
    - Performance bottleneck identification
    - Service dependency mapping
- [Installation](https://www.jaegertracing.io/docs/1.6/getting-started/)

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
