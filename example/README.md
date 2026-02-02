# govalid Examples

This directory contains working examples demonstrating how to use govalid.

## Prerequisites

Install govalid:

```bash
go install github.com/sivchari/govalid/cmd/govalid@latest
```

## Basic Example

The `basic/` directory shows fundamental usage of govalid:

- Defining structs with validation markers
- Running the code generator
- Using generated validation functions
- Iterating through validation errors

### Running the Example

```bash
cd basic

# Generate validation code
go generate ./...

# Run the example
go run .
```

### What the Example Demonstrates

1. **Valid User Validation** - Shows successful validation
2. **Invalid User Validation** - Shows multiple validation errors
3. **Loop Through Errors** - Shows how to iterate over `ValidationErrors`
4. **Check Specific Error Type** - Shows how to filter errors by type
5. **Using Validate() Method** - Shows the `govalid.Validator` interface

### Error Handling

govalid returns `ValidationErrors` which is a slice of `ValidationError`. You can iterate over them:

```go
err := ValidateUser(user)
if errs, ok := err.(govaliderrors.ValidationErrors); ok {
    for _, e := range errs {
        fmt.Printf("Field: %s, Type: %s, Reason: %s\n",
            e.Path, e.Type, e.Reason)
    }
}
```

Each `ValidationError` contains:
- `Path` - The field path (e.g., "User.Name", "User.Address.City")
- `Type` - The validation type (e.g., "required", "email", "maxlength")
- `Value` - The actual value that was validated
- `Reason` - A human-readable error message
