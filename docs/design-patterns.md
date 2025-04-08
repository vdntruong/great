# Design Patterns

[code example](https://refactoring.guru/design-patterns/go) 

(((Design pattern) < Architecture pattern) < Architecture style)

## 1. Creational Patterns

### Factory Method aka Virtual Constructor

**Factory method** is a creational design pattern that **provides an interface for creating objects** in a superclass, 
but allows subclasses to alter the type of objects that will be created.

```go
package main

// super class/interface

type Payment interface {
    Process(amount int) error
}
```

```go
package main
// sub classes

type CryptoPayment struct {}

func (p *CryptoPayment) Process(amount int) error {
    // ...
    return nil
}

type CreditCardPayment struct {}

func (p *CreditCardPayment) Process(amount int) error {
    // ...
    return nil
}

type PayPalPayment struct {}

func (p *PayPalPayment) Process(amount int) error {
    // ...
    return nil
}
```

```go
package main

type PaymentMethod string

// or we can use iota (enum)

const (
    CreditCard PaymentMethod = "creditcard"
    PayPal     PaymentMethod = "paypal"
    Crypto     PaymentMethod = "crypto"
)

// Factory method

func CreatePayment(method PaymentMethod) (Payment, error) {
    switch method {
    case CreditCard:
        return &CreditCardPayment{}, nil
    case PayPal:
        return &PayPalPayment{}, nil
    case Crypto:
        return &CryptoPayment{}, nil
    default:
        return nil, fmt.Errorf("unsupported payment method: %v", method)
    }
}
```
### Abstract factory

### Builder

### Options

### Prototype

### Singleton

## 2. Structural Patterns

### Facade

### Adapter

### Composite

## 3. Behavioral Patterns

### Strategy

**Strategy** is a behavioral design pattern that lets you define a family of algorithms, 
put each of them into a separate class, and make their objects interchangeable.

```go
type Notifier interface {
    SendNotification(context.Context, interface{}) error
}

type EmailNotifier struct {}

func (e *EmailNotifier) Send(ctx context.Context, payload interface{}) error {
    // ...
    return nil
}

type InAppNotifier struct {}

func (e *InAppNotifier) Send(ctx context.Context, payload interface{}) error {
    // ...
    return nil
}
```

### Chain of Responsibility

### State

### Observer
