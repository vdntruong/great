# Go Idiomatic ([Effective Go](https://go.dev/doc/effective_go))

## Format

Use `gofmt`!

## Imports

Imports are organized in groups, with blank lines between them. The standard library packages are always in the first group.
```
package main

import (
	"context"
	"errors"
	"net/http"

	"user-ms/internal/dto"
	"user-ms/internal/model"
	"user-ms/internal/pkg/apperror"

	"golang.org/x/sync/errgroup"
	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
)
```

[goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) can help us

## Naming

### Project Names

### Package Names

Good package names are short and clear (`http`, `sync`, `bytes`,...). They are lower case, with NO `under_scores` or `mixedCaps`.

For package, use singular noun instead of plural (they are often simple nouns).

If your package adds some extension to the current standard libraries, 
try to use another name instead of using the same name and fool yourselves with import alias. 
Use `contextutil` or `contextx` instead of `context`, use `errorsx` instead of `errors`

[package names](https://go.dev/blog/package-names).

### File Names

For file name, try to keep it short, but still meaningful.

Use lower case and use `snake_case`, DO NOT use `mixedCaps`.

If there are multiple files for different purposes in a same package, 
all the files of a same purpose should be prefixed the same way, 
that way your code would be perfectly aligned in your IDE and super easy to find.

### Function Names

Use `mixedCaps` for naming your function. Use `getArticles` or `GetArticles`

Try to void the word `Get_` from the getters of a struct.

Let's do:
```
type User struct {
    age int
}

func (u *User) Age() int {
    return u.age
}

func (u *User) SetAge(a int)  {
    u.age = a
}
```

With the function as handler function, let's name it with prefix `Handle_`: `HandleLogin()`, `HandleCancelOrder()`.

If function can through a panic (the logic error shouldn't appear) names it with prefix `Must_`: `MustParseInt(s string) int`.
```
func MustParseInt(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return i
}
```

```
func ParseInt(s string) (int, error) {
    i, err := strconv.Atoi(s)
   return i, err
}
```

### Variable Names

Name of variable should be short but should be meaningful enough.

Keep the receiver name short.

Use short names for variables with limited scope.
If a function requires only one struct as its request param, `req` is much better than `getArticleRequest`. 
If it requires multiple param, concrete names should be provided but tries to keep them short.

## Variables

## Avoid global variables

## Contexts

Context should be the first parameter on methods, functions.
If a function receives a context, the context must be the first parameter and its name should be ctx.

## Panic

## Handles Errors

### Add context to errors

## Interfaces

Interface should have suffix `_er`: `Provider`, `Adapter`, `Storer`, ...

```
type UserAdapter interface {
    Get(ctx context.context, id string) error
}
```

Good Go dev don't do suffix `_Impl` for implementation tier. Instead of that, we can name it with related prefix. 

Prefer to not using empty interface. If the function handle more than one type of data, just clarify one by one and handle it explicitly.

Don't return interfaces.

## Goroutines

## Testing

### Prefer table driven tests
### Avoid assert tests
## Line Length

## Comments

## References

1. Effective Go
2. Go Code Review Comments
3. Best practices for a new go developer
4. Error handling in Go
5. Errors are values
6. Slice tricks
7. Table driven testing 
8. HTTP services best practices - Mat Ryer
9. Go best practices - Peter Bourgon
10. Twelve Go Best Practices - Francesc Campoy Flores
11. Don't use Go's default HTTP client in production
12. Go project layout
13. Building APIs - Mat Ryer
14. Context isn't for cancellation - Dave Cheney
15. The package level logger anti pattern - Dave Cheney
16. Let's go - Alex Edwards
