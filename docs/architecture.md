# Domain Driven Design

DDD is not about code.

> [!INFO]
> DDD does not resonate in Golang 

**Strategic Design:** \
It is the one in which you and the domain experts analyze a domain, define its bounded contexts, and look for the best way to let them communicate. \
As you can easily assume, the strategic design is programming language agnostic.

```
app
│
└─customer
│ │ customer.go
└─product
│ │ product.go
└─user
│ │ user.go
│─events
  │ customer_events.go
  │ product_events.go
  │ user_events.go
```

```
app
│
└─customer
│ │ events.go
│ │ customer.go
└─product
│ │ events.go
│ │ product.go
└─user
  │ events.go
  │ user.go
```

The decoupling problem:

```
app
│
└─delivery
│ │ delivery.go  // may import product
└─product
  │ product.go   // or may import delivery
```

We can use the **Anti-Corruption Layer Pattern** (the first introduced by Eric Evans in the book Domain-Driven Design, 2003). \
This pattern is typically implemented using a combination of the Facade and Adapter design patterns.

```
app
│
└─delivery
│ │ delivery.go
└─product
│ │ product.go
└─productdelivery 
  │ product_delivery.go  // import product and delivery
```

The **"Big Ball of Mud"** is a term used in software to describe a system that has evolved into a tangled, disorganized mess over time. \
It lacks a clear architectural structure, with components and modules tightly intertwined confusingly. \
This anti-pattern arises when developers prioritize short-term fixes and feature additions over long-term architectural planning and design. \
As a result, the codebase becomes difficult to understand, maintain, and extend, leading to increased development costs and decreased system reliability.

A little copy is better than a little dependency.

**Tactical Design:** \
Describes a group of patterns to use to shape as code the invariants and models defined in a domain analytics, often driven by the strategic design. \
The end goal of applying those patterns is to model the code in a simple but expressive and safe way.

Value Type (aka Value Object) \
The always valid state, the unachievable always valid state.

```go
package demo

// Title represents a tab title
type Title string

func (t Title) String() string {
    return string(t)
}

func (t Title) Equals(t2 Title) bool {
    return t.String() == t2.String()
}
```

Repository

```go
package demo

import "errors"

type Title string

type ID string 

type Tab struct {
    ID ID
    Title Title
}

var (
    // Errors returned by the repository
    ErrRepoNextID = errors.New("tab: could not return next id")
    ErrRepoList   = errors.New("tab: could not list")
    ErrNotFound   = errors.New("tab: could not find")
    ErrRepoGet    = errors.New("tab: could not get")
    ErrRepoAdd    = errors.New("tab: could not add")
    ErrRepoRemove = errors.New("tab: could not remove")
)

type ReadRepo interface {
    // List returns a tab slice and an error in case of failure
    List() ([]*Tab, error)
    // Find returns a tab or nil if it is not found and an error in case of failure
    Find(id ID) (*Tab, error)
    // Get returns a tab and error in case is not found or failure
    Get(id ID) (*Tab, error)
}

type WriteRepo interface {
    // NextID returns the next free ID and an error in case of failure
    NextID() (ID, error)
    // Add persists a tab (already existing or not) and returns an error in case of failure
    Add(t *Tab) error
    // Remove removes a tab and returns an error in case is not found or failure
    Remove(id ID) error
}

type Repo interface {
    // List returns a tab slice and an error in case of failure
    List() ([]*Tab, error)
    // Find returns a tab or nil if it is not found and an error in case of failure
    Find(id ID) (*Tab, error)
    // Get returns a tab and error in case is not found or failure
    Get(id ID) (*Tab, error)
    // NextID returns the next free ID and an error in case of failure
    NextID() (ID, error)
    // Add persists a tab (already existing or not) and returns an error in case of failure
    Add(t *Tab) error
    // Remove removes a tab and returns an error in case is not found or failure
    Remove(id ID) error
}
```

Repository: hint

Don't
```go
package demo

_, err := repo.List(repo.Filter{"active": true})
```

Do
```go
package demo

_, err := repo.ListActive()
```

Repository: where to place it

```
app
│
└─internal
│ │─tab
│   │─reop.go  // here a MySQL implementation
└─tab
  │ repo.go    // here a the interface and the errors
```

Other patterns:
- Entity
- Aggregate
- Aggregate Root
- Domain Service

# Object Oriented Programming

Golang is not OOP
