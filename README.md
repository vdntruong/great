## the great

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
