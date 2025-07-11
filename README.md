# govalid

## Overview

govalid is a Go package designed to generate type-safe validation code for structs based on markers. It provides a mechanism to apply validation rules directly in the code by marking struct fields with specific markers. The tool processes these markers and generates corresponding validation functions.

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

## govalid-Specific Features

govalid supports all common validation markers with performance-optimized implementations:

**Numeric Validators:**
- **`govalid:gt`**: Greater than validation
- **`govalid:gte`**: Greater than or equal validation  
- **`govalid:lt`**: Less than validation
- **`govalid:lte`**: Less than or equal validation

**String Validators:**
- **`govalid:minlength`**: Minimum string length (Unicode-aware)
- **`govalid:maxlength`**: Maximum string length (Unicode-aware)
- **`govalid:email`**: HTML5-compliant email validation with zero allocations
- **`govalid:url`**: HTTP/HTTPS URL validation with manual parsing for optimal performance
- **`govalid:uuid`**: RFC 4122 UUID validation with inline functions

**Collection Validators:**
- **`govalid:minitems`**: Minimum collection size (slice, array, map, channel)
- **`govalid:maxitems`**: Maximum collection size (slice, array, map, channel)

**General Validators:**
- **`govalid:required`**: Required field validation with Go zero-value semantics
- **`govalid:enum`**: Comprehensive enum validation for string, numeric, and custom types

**govalid-Exclusive Features:**
- **Extended collection support**: Maps and channels work with `maxitems`/`minitems` (not supported by go-playground/validator)
- **Zero-allocation validation**: All validators perform zero heap allocations
- **Compile-time safety**: Type-safe validation functions generated at compile time

## License

govalid is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
