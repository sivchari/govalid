<div align="center">
  <img src="assets/govalid-icon.svg" alt="govalid" width="128" height="128">
  <h1>govalid</h1>
  <p><strong>Blazing fast, zero-allocation, type-safe validation for Go</strong></p>

  ![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)
  ![License](https://img.shields.io/badge/license-MIT-green.svg)
  [![CI](https://github.com/sivchari/govalid/actions/workflows/unit-test.yaml/badge.svg)](https://github.com/sivchari/govalid/actions/workflows/unit-test.yaml)
  [![codecov](https://codecov.io/gh/sivchari/govalid/graph/badge.svg)](https://codecov.io/gh/sivchari/govalid)
  [![Go Report Card](https://goreportcard.com/badge/github.com/sivchari/govalid)](https://goreportcard.com/report/github.com/sivchari/govalid)
</div>

---

govalid generates **type-safe validation code** from struct field markers. No reflection, no runtime overhead — just fast validation.

## Why govalid?

- **Zero allocations** — All validation performs no heap allocations
- **5-44x faster** — Outperforms reflection-based validators
- **Type safe** — Errors caught at generation time, not runtime
- **Full collection support** — Maps, channels, slices, and arrays (not just slices)
- **CEL expressions** — Complex validation with Common Expression Language

## Installation

```bash
go install github.com/sivchari/govalid/cmd/govalid@latest
```

Verify the installation:

```bash
govalid -h
```

## Quick Start

### 1. Define Your Struct

```go
//govalid:required
type Person struct {
    Name  string `json:"name"`
    //govalid:email
    Email string `json:"email"`
}
```

### 2. Generate Validation Code

```bash
govalid ./...
```

### 3. Use the Validator

```go
func main() {
    p := &Person{Name: "John", Email: "invalid-email"}

    if err := p.Validate(); err != nil {
        log.Printf("Validation failed: %v", err)
    }
}
```

## Supported Markers

For complete details, see [MARKERS.md](MARKERS.md).

### Field Validators

| Marker | Description | Example |
|--------|-------------|---------|
| `required` | Field must not be zero value | `//govalid:required` |
| `gt` | Greater than | `//govalid:gt=0` |
| `gte` | Greater than or equal | `//govalid:gte=1` |
| `lt` | Less than | `//govalid:lt=100` |
| `lte` | Less than or equal | `//govalid:lte=99` |
| `maxlength` | Maximum string length | `//govalid:maxlength=255` |
| `minlength` | Minimum string length | `//govalid:minlength=1` |
| `maxitems` | Maximum collection size | `//govalid:maxitems=10` |
| `minitems` | Minimum collection size | `//govalid:minitems=1` |
| `enum` | Must be one of specified values | `//govalid:enum=active,inactive` |

### Format Validators

| Marker | Description |
|--------|-------------|
| `email` | Valid email address |
| `url` | Valid URL |
| `uuid` | Valid UUID |
| `numeric` | Numeric string |

### Advanced Validators

| Marker | Description | Example |
|--------|-------------|---------|
| `cel` | CEL expression | `//govalid:cel=value >= 18` |

## Struct-Level Validation

Apply markers to the entire struct:

```go
//govalid:required
type Person struct {
    Name  string  // required
    Email string  // required
    Age   int     // required
}
```

## CEL Expression Support

Use [Common Expression Language](https://github.com/google/cel-spec) for complex validation:

```go
type User struct {
    //govalid:cel=value >= 18 && value <= 120
    Age int
    //govalid:cel=value >= this.Age
    RetirementAge int
}
```

## Collection Support

Validate maps, channels, slices, and arrays:

```go
type UserList struct {
    //govalid:maxitems=10
    Users []User

    //govalid:minitems=1
    UserMap map[string]User

    //govalid:maxitems=5
    UserChan chan User
}
```

## Error Handling

### Single Error

```go
if err := p.Validate(); err != nil {
    log.Printf("Validation failed: %v", err)
}
```

### Multiple Errors

```go
if err := p.Validate(); err != nil {
    if errors.Is(err, ErrPersonEmailEmailValidation) {
        // Handle email error
    }
    if errors.Is(err, ErrPersonNameRequiredValidation) {
        // Handle required error
    }
}
```

### HTTP Middleware Integration

```go
import "github.com/sivchari/govalid/validation/middleware"

func main() {
    http.HandleFunc("/person", middleware.ValidateRequest[*Person](handler))
    http.ListenAndServe(":8080", nil)
}
```

## Context-Aware Validation

govalid supports context-aware validation for request cancellation and timeout handling:

```go
// Standard validation (no context)
if err := p.Validate(); err != nil {
    log.Printf("Validation failed: %v", err)
}

// Context-aware validation with timeout
ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
defer cancel()

if err := p.ValidateContext(ctx); err != nil {
    switch err {
    case context.DeadlineExceeded:
        log.Println("Validation timed out")
    case context.Canceled:
        log.Println("Validation cancelled")
    default:
        log.Printf("Validation failed: %v", err)
    }
}
```

Context-aware validation is particularly useful for:

- **HTTP handlers** - Respect client request cancellation
- **Timeouts** - Prevent runaway validation processes
- **High-concurrency servers** - Handle many concurrent requests efficiently

See [Context Validation Documentation](docs/content/context-validation.md) for complete details and examples.

## Performance

govalid outperforms reflection-based validators by **5x to 44x** with **zero allocations**.

| Validator | govalid | go-playground | Improvement |
|-----------|---------|---------------|-------------|
| Required | 1.9ns | 85.5ns | 44x |
| GT/LT | 1.9ns | 63.0ns | 32x |
| MaxLength | 15.7ns | 73.5ns | 5x |
| Email | 38.2ns | 649.4ns | 17x |

See [benchmark details](test/benchmark/README.md) for more information.

## Development

### Setup

```bash
git clone https://github.com/sivchari/govalid.git
cd govalid
make install-lefthook
```

### Build

```bash
go install ./cmd/govalid/
```

## License

MIT License - see [LICENSE](LICENSE) for details.
