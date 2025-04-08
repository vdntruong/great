# Order Microservice

A microservice for managing orders in an e-commerce system, built with Go and following clean architecture principles.

## Architecture

The service follows a clean, layered architecture with clear separation of concerns:

### Layers

1. **API Layer** (`internal/api/`)
   - Handles HTTP requests/responses
   - Input validation
   - Error handling
   - Uses DTOs for request/response
   - Routes requests to appropriate service methods

2. **Service Layer** (`internal/service/`)
   - Contains business logic
   - Orchestrates operations
   - Validates business rules
   - Converts between models and DTOs
   - Handles transactions
   - Split into:
     - `order_service.go`: Main service implementation
     - `order_converter.go`: Type conversion utilities
     - `order_validator.go`: Input validation

3. **Repository Layer** (`internal/repository/`)
   - Data access layer
   - Handles database operations
   - Uses SQLC for type-safe database queries
   - Contains DAO (Data Access Objects)

4. **Model Layer** (`internal/models/`)
   - Defines core business entities
   - Contains validation rules
   - Used across all layers
   - Includes:
     - `Order`
     - `OrderItem`
     - `CreateOrderParams`
     - `UpdateOrderParams`
     - `ListOrdersParams`

### Features

1. **Order Management**
   - Create new orders
   - Update order status
   - Cancel orders
   - List orders with filtering
   - Get order details

2. **Order Status Flow**
   - Pending (initial state)
   - Paid
   - Shipping
   - Delivered
   - Canceled

3. **Data Validation**
   - Input validation at service layer
   - Business rule validation
   - Status transition validation

## Getting Started

### Prerequisites

- Go 1.21 or later
- PostgreSQL 15 or later
- SQLC

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up the database:
   ```bash
   psql -U postgres -f db/migrations/000002_create_orders_table.up.sql
   ```
4. Generate SQLC code:
   ```bash
   sqlc generate
   ```

### Running the Service

```bash
go run cmd/server/main.go
```

## Testing

Run tests:
```bash
go test ./...
```
