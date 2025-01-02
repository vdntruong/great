# Circuit Breaker Pattern in Microservices

## Overview
Circuit Breaker is a design pattern used in microservices architecture to prevent cascading failures and improve system resilience. \
It works similarly to an electrical circuit breaker, cutting off the flow of requests when a service is experiencing problems.

Refer package: [gobreaker](https://github.com/sony/gobreaker)

## How It Works

```
CLOSED -> (failures reach threshold) -> OPEN -> (wait reset timeout) -> HALF-OPEN -> (success threshold met) -> CLOSED
           or
CLOSED -> (failures reach threshold) -> OPEN -> (wait reset timeout) -> HALF-OPEN -> (failure occurs) -> OPEN
```

### States
1. **CLOSED** (Normal Operation)
    - Requests flow through normally
    - Failures are counted
    - Switches to OPEN if failure threshold is reached

2. **OPEN** (Service Protection)
    - Requests **fail fast**
    - No calls to problematic service
    - After timeout period, switches to HALF-OPEN

3. **HALF-OPEN** (Testing Recovery)
    - Limited requests allowed through
    - Success -> CLOSED
    - Failure -> OPEN

### Configuration Parameters
- Failure Threshold: Number of failures before circuit opening
- Reset Timeout: Time to wait before retrying
- Success Threshold for Recovery: Number of successful requests required in HALF-OPEN state before circuit closes
- Failure Criteria:
  - Request timeout exceeded
  - Response time: >= 5s (normal 200ms)
  - Connection failures
  - HTTP 5xx responses
  - ... (e.g., business logic errors)

## Key Advantages

1. **Resource Protection**
    - Prevents thread/connection exhaustion
    - Allows failing services to recover
    - Maintains system resources

2. **Fast Failure Response**
    - No waiting for timeouts
    - Immediate error response
    - Better user experience

3. **System Resilience**
    - Prevents cascading failures
    - Allows partial system operation
    - Self-healing capability

4. **Monitoring and Detection**
    - Clear service health indicators
    - Early problem detection
    - Better system observability

## When to Use Circuit Breakers

1. **Recommended Scenarios**
    - External service calls
    - Database operations
    - Network-dependent services
    - Resource-intensive operations

2. **Implementation Considerations**
    - Failure threshold settings
    - Timeout configurations
    - Monitoring requirements
    - Fallback strategies

## Best Practices

1. **Configuration**
    - Set appropriate thresholds
    - Configure meaningful timeouts
    - Regular monitoring and adjustment

2. **Implementation**
    - Use with retry mechanisms
    - Implement fallback logic
    - Monitor circuit state
    - Log state transitions

3. **Maintenance**
    - Regular testing of fallback scenarios
    - Monitor failure patterns
    - Adjust parameters based on metrics

## Benefits vs Traditional Error Handling

- Fast Failure: Immediate failure instead of waiting for timeout
- Self Healing: Automatic recovery when service health improves
- Graceful Degradation: System continues partial operation
- Monitoring: Helps detect and track service issues

|                    | Without Circuit Breaker                                               | With Circuit Breaker                                                  |
|--------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|
| Performance        | Each request must wait for timeout/retry -> wastes time and resources | Fails fast when service is known to be problematic -> reduces latency |
| Service Protection | Continues sending requests despite service overload                   | Pauses requests, allowing service recovery time                       |
| System Health      | Continuous failed calls can lead to cascading failures                | Prevents error propagation through system                             |
| Monitoring         | Difficult to track service status                                     | Circuit breaker state indicates service health                        |
| Resource Usage     | Threads/connections blocked waiting for retry                         | Resources freed immediately when service is unavailable               |

### Traditional Approach (Without Circuit Breaker: Retry)
```go
func callService() error {
    for i := 0; i < maxRetries; i++ {
        err := makeRequest()
        if err == nil {
            return nil
        }
        time.Sleep(backoff)
    }
    return handleFallback()
}
```

### With Circuit Breaker
```go
func callService() error {
    if circuitBreaker.IsOpen() {
        return handleFallback()  // Fail fast
    }
    
    if err := makeRequest(); err != nil {
        circuitBreaker.RecordFailure()
        if circuitBreaker.ShouldOpen() {  // Check if threshold reached
            circuitBreaker.Open()
        }
        return handleFallback()
    }
    
    circuitBreaker.RecordSuccess()
    return nil
}
```

## Real-World Example: E-commerce System

### System Architecture
```
Order Service -> Payment Service -> Inventory Service -> Shipping Service
```

### Cascading Failure Scenario

#### Without Circuit Breaker

```
Inventory (fails) -> Payment (blocks) -> Order (exhausts resources) -> System-wide impact (cascading failure)
```

1. **Initial Failure (Inventory Service)**
    - High CPU (90%)
    - Slow response (5s vs normal 200ms)
    - Increasing error rate

2. **Cascade to Payment Service**
    - Blocked threads waiting for Inventory
    - Resource exhaustion
    - Service degradation (starting to respond error to client)

3. **Impact on Order Service**
    - Pending transactions pile up
    - Database connections exhausted
    - Increasing memory usage
    - System-wide slowdown

4. **Final State**
    - Complete system failure
    - No new orders possible
    - All related or maybe non-related functions affected

#### With Circuit Breaker
```
Order Service -> [CB] -> Payment Service -> [CB] -> Inventory Service -> [CB] -> Shipping Service
```

When Inventory fails:
1. Circuit Breaker between Payment-Inventory opens
2. Payment receives immediate failure response, without any blocking (fast fail)
3. Order service can handle fallback gracefully
4. System remains partially operational

## Conclusion
Circuit Breakers are essential in modern microservices architectures, providing:
- System stability
- Failure isolation
- Resource protection
- Better user experience

When combined with proper monitoring and fallback strategies, they significantly improve system reliability and maintainability.
