# Product Microservice

A microservice for managing products in an e-commerce system, built with Go and following clean architecture principles.

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
     - `product_service.go`: Main service implementation
     - `product_converter.go`: Type conversion utilities
     - `product_validator.go`: Input validation

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
     - `Product`
     - `CreateProductParams`
     - `UpdateProductParams`
     - `ListProductsParams`

5. **DTO Layer** (`internal/dto/`)
   - Data Transfer Objects
   - Used for API input/output
   - Separates internal models from external API
   - Includes:
     - `CreateProductRequest`
     - `UpdateProductRequest`
     - `ProductResponse`

### Data Flow

```
HTTP Request → API Layer → Service Layer → Repository Layer → Database
Response ← API Layer ← Service Layer ← Repository Layer ← Database
```

### Key Principles

1. **Separation of Concerns**
   - Each layer has a specific responsibility
   - Clear boundaries between layers
   - Dependencies flow in one direction

2. **Type Safety**
   - Strong typing throughout
   - SQLC for type-safe database queries
   - Validation at multiple levels

3. **Clean Code**
   - Single Responsibility Principle
   - DRY (Don't Repeat Yourself)
   - Clear naming conventions
   - Proper error handling
   - Input validation

4. **Dependency Management**
   - Dependencies flow inward
   - Outer layers depend on inner layers
   - No circular dependencies

### Error Handling

- Consistent error types
- Proper error propagation
- Validation at multiple levels
- Clear error messages

### Type Safety

- Strong typing throughout
- SQLC for type-safe database queries
- Validation at multiple levels
- Clear conversion between types

## Benefits

- Clear separation of concerns
- Easy to test each layer independently
- Maintainable and scalable
- Type-safe operations
- Clear data flow
- Proper error handling
- Easy to extend with new features

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
   psql -U postgres -f migrations/000001_init_schema.up.sql
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

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
