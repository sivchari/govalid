# govalid Examples

This directory contains working examples demonstrating how to use govalid.

## Directory Structure

```
example/
├── go.work          # Go workspace configuration
├── basic/           # Struct definitions with validation markers
│   ├── go.mod
│   ├── user.go
│   └── user_user_validator.go (generated)
└── basic-usage/     # Example usage code
    ├── go.mod
    └── main.go
```

## Prerequisites

Install govalid:

```bash
go install github.com/sivchari/govalid/cmd/govalid@latest
```

## Running the Example

```bash
cd example
go run ./basic-usage/
```

## Regenerating Validation Code

If you modify `basic/user.go`, regenerate the validation code:

```bash
cd example/basic
go generate ./...
```

## What the Example Demonstrates

1. **Valid User Validation** - Shows successful validation
2. **Invalid User Validation** - Shows multiple validation errors
3. **Loop Through Errors** - Shows how to iterate over `ValidationErrors`
4. **Check Specific Error Type** - Shows how to filter errors by type
5. **Using Validate() Method** - Shows the `govalid.Validator` interface

## Error Handling

govalid returns `ValidationErrors` which is a slice of `ValidationError`. You can iterate over them:

```go
err := basic.ValidateUser(user)
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
