---
title: "Context-Aware Validation"
description: "Request-scoped validation with cancellation and timeout support"
weight: 30
---

# Context-Aware Validation

Govalid generates context-aware validation methods for every struct.
Enabling request-scoped validation with cancellation and timeout handling.

## Generated Methods

For each struct with validation markers, four methods are generated:

```go
func ValidatePersonRequest(t *PersonRequest) error

func ValidatePersonRequestContext(ctx context.Context, t *PersonRequest) error

func (t *PersonRequest) Validate() error

func (t *PersonRequest) ValidateContext(ctx context.Context) error
```

The `Validate()` and `ValidatePersonRequest()` methods delegate to their context aware equivalents using `context.Background()`.

## How Context Cancellation Works

Context cancellation is checked **between field validations**.

Generated code follows this pattern:

```go
func ValidatePersonRequestContext(ctx context.Context, t *PersonRequest) error {
    if t == nil {
        return ErrNilPersonRequest
    }

    var errs govaliderrors.ValidationErrors

    if ctx.Err() != nil {
        return ctx.Err()
    }

    if t.Name == "" {
        err := ErrPersonRequestNameRequiredValidation
        err.Value = t.Name
        errs = append(errs, err)
    }

    if ctx.Err() != nil {
        return ctx.Err()
    }

    if t.Email == "" {
        err := ErrPersonRequestEmailRequiredValidation
        err.Value = t.Email
        errs = append(errs, err)
    }

    if !validationhelper.IsValidEmail(t.Email) {
        err := ErrPersonRequestEmailEmailValidation
        err.Value = t.Email
        errs = append(errs, err)
    }

    if len(errs) > 0 {
        return errs
    }
    return nil
}
```

Each `ctx.Err()` check provides a cancellation point before the next field's validation.

## Usage Examples

### With Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
defer cancel()

err := ValidatePersonRequestContext(ctx, &person)
if err == context.DeadlineExceeded {
    log.Println("Validation timed out")
}
```

### With Cancellation

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    err := ValidatePersonRequestContext(ctx, &person)
    if err == context.Canceled {
        log.Println("Validation cancelled")
    }
}()

// Cancel from another goroutine
cancel()
```

### HTTP Handler

```go
func createPersonHandler(w http.ResponseWriter, r *http.Request) {
    var req PersonRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if err := req.ValidateContext(r.Context()); err != nil {
        if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "Request cancelled", http.StatusRequestTimeout)
            return
        }
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Process valid request...
}
```

## HTTP Middleware

govalid provides middleware for automatic request validation:

```go
import "github.com/sivchari/govalid/validation/middleware"

// Standard validation (no context awareness)
http.HandleFunc("/users", middleware.ValidateRequest[PersonRequest](handler))

// Context-aware validation (honors request cancellation)
http.HandleFunc("/users", middleware.ValidateRequestContext[PersonRequest](handler))
```

- `ValidateRequest[T]` - Calls `Validate()`, ignores request context
- `ValidateRequestContext[T]` - Calls `ValidateContext(r.Context())`, respects client cancellation

## Migration

Existing code requires no changes. To add context support:

```go
// Before
err := ValidatePersonRequest(&person)

// After (with context)
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
err := ValidatePersonRequestContext(ctx, &person)
```
