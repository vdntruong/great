# Microservices Communication Patterns

## Communication Types

### 1. External Communication (Public APIs)
External clients access microservices through an API Gateway (e.g., Traefik):

```
External Client → API Gateway (Traefik) → Service Instances
                          ↓
                  [Gateway Features]
                  - Load Balancing
                  - SSL/TLS
                  - Auth
                  - Rate Limiting
                  - Circuit Breaking
```

Characteristics:
- Exposed to external clients
- Requires authentication and authorization
- Implements rate limiting and security policies
- Monitored and logged at gateway level

The gateway, which provides:
- Load balancing
- SSL termination
- Authentication
- Rate limiting
- Monitoring
- Circuit breaking

### 2. Internal Communication (Service-to-Service)
Services communicate directly using Service Discovery:

```
Service A → Service Discovery → Load Balancer → Service B Instances
                    ↓
            [Discovery Process]
            1. Registration
            2. Health Check
            3. Service Lookup
            4. Load Balancing
```

Characteristics:
- Direct communication between services
- Uses service discovery (e.g., Consul, Eureka)
- More efficient (no gateway overhead)
- Secured within internal network

Service Discovery workflow:
- Service registration
- Health checking
- Service lookup
- Load balancing

## Service Discovery Flow

```
    [Service Registration]
            ↓
      [Health Check]
            ↓
     [Service Lookup] ←----→ [Cache]
            ↓
     [Load Balancing]
            ↓
[Service Instance Selection]
```

## Circuit Breaker Patterns

Circuit Breaker Failure Flow
```
Request → [Circuit Breaker Check]
           /                \
    [Circuit Open]    [Circuit Closed]
          ↓                 ↓
    Return Fallback    Forward Request
                           /        \
                    [Success]    [Failure]
                                    ↓
                            [Increment Failure Count]
```

### 1. Gateway-Level Circuit Breaker

```
             [Traefik Gateway]
                     ↓
             [Circuit Breaker]
             /      |       \
      Service A  Service B   Service C
      /   \      /   \      /   \
Instance  Instance   Instance   Instance
```

Implemented at the API Gateway (Traefik):
- Monitors all traffic to a service
- Global circuit breaker state
- Affects all instances of target service
- Centralized configuration

### 2. Service-Level Circuit Breaker

Implemented within each service:
- Monitors service-to-service calls
- Works with service discovery
- Handles multiple instances of target service
- Configured at code level

#### Circuit Breaker with Service Discovery

1. Service-Level Circuit Breaker (All Instances)

```
Service A → [Single Circuit Breaker] → Service B
                                       /    \
                                  Instance  Instance
```

2. Per-Instance Circuit Breaker

```
Service A → [Circuit Breaker 1] → Service B Instance 1
          ↘ [Circuit Breaker 2] → Service B Instance 2
```

## Multiple Instances Handling

### Circuit Breaker with Multiple Instances
When target service has multiple instances:

1. Instance-Level Approach (Common):
- Single circuit breaker for all instances
- Tracks aggregate error rate
- Opens circuit for all instances when threshold reached
- Works with service discovery load balancing

2. Per-Instance Approach (Rare):
- Separate circuit breaker per instance
- More complex management
- Rarely used as service discovery handles instance health

## Best Practices

1. Communication Patterns:
   - Use API Gateway for external communication
   - Use direct service discovery for internal calls
   - Implement proper network segmentation

2. Circuit Breaker Implementation:
   - Combine circuit breakers with retry mechanisms
   - Implement fallback behavior
   - Monitor circuit breaker metrics
   - Use appropriate thresholds based on service requirements

3. Circuit Breaker Configuration:
   - Set appropriate thresholds based on service SLAs
   - Configure meaningful timeout values
   - Implement proper fallback mechanisms 
   - Use sliding windows for better accuracy

4. Service Discovery:
   - Implement health checks 
   - Use appropriate registration/deregistration policies 
   - Configure proper refresh intervals 
   - Handle service discovery failures gracefully

5. Error Handling:
   - Implement proper fallback mechanisms
   - Log circuit breaker state changes
   - Monitor failure rates and patterns
   - Set up alerts for circuit breaker events

6. Testing:
   - Test circuit breaker behavior with chaos engineering
   - Verify fallback mechanisms
   - Test service discovery failover
   - Simulate network failures

7. Instance Handling:
   - Prefer instance-level circuit breaking
   - Let service discovery handle instance health
   - Implement proper fallback mechanisms

8. Security:
   - Secure internal communication channels
   - Implement proper authentication for service-to-service calls
   - Use network policies to restrict communication

## Monitoring and Maintenance

1. Circuit Breaker Metrics:
   - Track failure rates
   - Monitor circuit state changes
   - Alert on circuit open events

2. Service Health:
   - Implement health checks
   - Monitor service discovery registry
   - Track instance availability

3. Performance Monitoring:
   - Track response times
   - Monitor error rates
   - Measure service dependencies



> [!QUESTION]
> How circuit breaker work on Event-driven architecture?
