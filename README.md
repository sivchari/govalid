# govalid

## Overview

govalid is a Go package designed to generate type-safe validation code for structs based on markers. It provides a mechanism to apply validation rules directly in the code by marking struct fields with specific markers. The tool processes these markers and generates corresponding validation functions.

## Why govalid?

govalid addresses key limitations of reflection-based validation libraries by generating type-safe, high-performance validation code at compile time.

### Performance Benefits
- **Zero allocations**: All validation functions perform zero heap allocations
- **Up to 45x faster**: Significantly outperforms reflection-based validators
- **Compile-time optimization**: Generated code is optimized by the Go compiler

### Developer Experience
- **Type safety**: Validation functions are generated with proper types, eliminating runtime reflection
- **Compile-time errors**: Invalid validation rules are caught during code generation, not at runtime
- **No runtime dependencies**: Generated code has minimal external dependencies

### Extended Go Support
- **Full collection support**: Maps and channels work with size validators (not supported by most libraries)
- **Go zero-value semantics**: Proper handling of Go's zero values and nil states
- **Unicode-aware**: String validators properly handle Unicode characters

### govalid vs Reflection-based Validators

| Feature | govalid | Reflection-based validators |
|---------|---------|----------------------------|
| Performance | ~1-14ns, 0 allocs | ~50-700ns, 0-5 allocs |
| Type Safety | Compile-time | Runtime reflection |
| Collection Support | slice, array, map, channel | slice, array only |
| Dependencies | Minimal | Heavy runtime deps |
| Error Detection | Compile-time | Runtime |

## Installation

To install govalid, use:

```bash
# Clone the repository
$ git clone https://github.com/sivchari/govalid.git

# Navigate to the project directory
$ cd govalid

# Run installation commands
$ go install ./...
```

## Usage

To use govalid, follow these steps:

1. Define your struct with field markers.
2. Run the govalid generation tool to create type-safe validation code.
3. Use the generated validation functions in your application.

### Example

```go
// Example struct with markers
// +govalid:required
type Person struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Generated code will include validation functions like:
func ValidatePerson(t *Person) error {
    if t == nil {
        return ErrNilPerson
    }

    if t.Name == "" {
        return ErrNameRequiredValidation
    }

    if t.Email == "" {
        return ErrEmailRequiredValidation
    }

    return nil
}
```

### Struct-Level Markers

govalid supports applying markers at the struct level. When a marker is placed on a struct, it applies to all fields within that struct:

```go
// Apply to entire struct - all fields will be validated
// +govalid:required
type Person struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

// Generated code for struct-level marker:
func ValidatePerson(t *Person) error {
    if t == nil {
        return ErrNilPerson
    }

    if t.Name == "" {
        return ErrNameRequiredValidation
    }

    if t.Email == "" {
        return ErrEmailRequiredValidation
    }

    if validatorhelper.IsZero(t.Age) {
        return ErrAgeRequiredValidation
    }

    return nil
}
```

## Supported Markers

For a complete reference of all supported markers, see [MARKERS.md](MARKERS.md).

## License

govalid is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
