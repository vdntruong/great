# Go Architecture Patterns: DAO, DTO, ORM, Entity, and Business Objects

## Data Access Object (DAO)

### Definition
DAO is a pattern that provides an abstract interface to some type of database or persistence mechanism, encapsulating all access to the data source.

### Characteristics
- Separates persistence logic from business logic
- Provides CRUD operations for data source interaction
- Can support multiple data sources
- Promotes testability through interface abstraction

### Go Implementation Example
```go
type UserDAO interface {
    Create(user *entity.User) error
    GetByID(id string) (*entity.User, error)
    Update(user *entity.User) error
    Delete(id string) error
}

type PostgresUserDAO struct {
    db *sql.DB
}

func (dao *PostgresUserDAO) Create(user *entity.User) error {
    // Implementation for PostgreSQL
}
```

### Best Practices
- Use interfaces to define DAO contracts
- Implement specific database adapters
- Handle connection pooling
- Include proper error handling
- Consider using context for timeouts

## Data Transfer Object (DTO)

### Definition
DTOs are objects that carry data between processes, typically used to transfer data between application layers.

### Characteristics
- Lightweight objects focused on data transfer
- No business logic
- May include data validation
- Often used for API responses/requests

### Go Implementation Example
```go
type UserDTO struct {
    ID        string `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

// Converter methods
func (dto *UserDTO) ToEntity() *User {
    return &User{
        ID:        dto.ID,
        Email:     dto.Email,
        FirstName: dto.FirstName,
        LastName:  dto.LastName,
    }
}
```

### Best Practices
- Keep DTOs simple and focused
- Include validation tags
- Implement conversion methods
- Use proper JSON/XML tags
- Consider versioning for APIs

## Object-Relational Mapping (ORM)

### Definition
ORM is a technique that lets you query and manipulate data from a database using an object-oriented paradigm.

### Popular Go ORMs
1. GORM
2. SQLBoiler
3. Ent (by Facebook)

### Characteristics
- Automatic query generation
- Relationship handling
- Migration support
- Hooks and callbacks
- Caching capabilities

### Example Using GORM
```go
type Product struct {
    gorm.Model
    Code  string
    Price uint
}

// Usage
db.AutoMigrate(&Product{})
db.Create(&Product{Code: "D42", Price: 100})
```

### Best Practices
- Use transactions when needed
- Implement proper indexing
- Handle N+1 query problems
- Consider using eager loading
- Implement proper error handling

## Entity

### Definition
Entities are domain objects that have a distinct identity that runs through time and different states.

### Characteristics
- Represents business domain concepts
- Has a unique identifier
- Contains business logic
- Maintains consistency rules

### Go Implementation Example
```go
type User struct {
    ID        uuid.UUID
    Email     string
    Password  []byte
    Profile   Profile
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (u *User) ChangePassword(newPassword string) error {
    // Business logic for password change
}
```

### Best Practices
- Keep entities focused on business rules
- Implement value objects
- Use proper validation
- Consider immutability
- Include domain events

## Business Domain/Object

### Definition
Business objects represent the entities and business logic of an application, encapsulating both data and behavior.

### Characteristics
- Contains business rules and logic
- Independent of persistence
- Maintains invariants
- Represents real-world concepts

### Implementation Example
```go
type Order struct {
    ID          string
    Items       []OrderItem
    CustomerID  string
    Status      OrderStatus
    TotalAmount Money
}

func (o *Order) CalculateTotal() Money {
    // Business logic for total calculation
}

func (o *Order) CanBeCancelled() bool {
    // Business rules for cancellation
}
```

### Best Practices
- Separate business logic from infrastructure
- Use domain events
- Implement value objects
- Follow DDD principles
- Use proper error handling

## Related Techniques and Patterns

### Repository Pattern
- Mediates between domain and data mapping layers
- Provides collection-like interface for domain objects

```go
type UserRepository interface {
    Save(user *User) error
    FindByID(id string) (*User, error)
    FindAll() ([]*User, error)
}
```

### Service Layer
- Defines application's boundary
- Coordinates response to user actions
- Encapsulates business logic

```go
type UserService struct {
    repo UserRepository
}

func (s *UserService) RegisterUser(dto UserRegistrationDTO) error {
    // Business logic for user registration
}
```

### Unit of Work
- Maintains list of objects affected by transaction
- Coordinates writing out changes
- Resolves concurrency problems

```go
type UnitOfWork struct {
    db         *sql.DB
    tx         *sql.Tx
    completed  bool
}

func (uow *UnitOfWork) Begin() error {
    // Start transaction
}
```

### Value Objects
- Immutable objects that represent descriptive aspects of the domain
- No conceptual identity

```go
type Money struct {
    amount   decimal.Decimal
    currency string
}

func NewMoney(amount decimal.Decimal, currency string) Money {
    return Money{amount, currency}
}
```

## Best Practices Summary

1. **Layer Separation**
    - Clear separation between layers
    - Dependency injection
    - Interface-based design

2. **Error Handling**
    - Custom error types
    - Proper error wrapping
    - Contextual error information

3. **Testing**
    - Unit tests for business logic
    - Integration tests for DAOs
    - Mock interfaces for testing

4. **Performance**
    - Connection pooling
    - Proper indexing
    - Caching strategies
    - Query optimization

5. **Maintainability**
    - Clear documentation
    - Consistent naming
    - Code organization
    - Dependency management
