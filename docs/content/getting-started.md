---
title: "Getting Started"
description: "Quick start guide for govalid"
weight: 10
---

# Getting Started

This guide will help you get up and running with govalid in just a few minutes.

## Installation

Install the govalid command-line tool:

```bash
go install github.com/sivchari/govalid/cmd/govalid@latest
```

Verify the installation:

```bash
govalid -h
```

## Basic Workflow

### 1. Add Validation Markers

Define your struct with validation markers in comments:

```go
package main

type User struct {
    // +govalid:required
    Name string `json:"name"`
    
    // +govalid:email
    Email string `json:"email"`
    
    // +govalid:gte=0
    // +govalid:lte=120
    Age int `json:"age"`
    
    // +govalid:maxlength=500
    Bio string `json:"bio,omitempty"`
}
```

### 2. Generate Validation Code

Run the govalid generator:

```bash
# Generate for current package
govalid .

# Generate for all packages recursively
govalid ./...

# Generate for specific package
govalid ./internal/models
```

This creates validation functions and error definitions:

```go
// Generated validation function
func ValidateUser(t *User) error {
    if t == nil {
        return ErrNilUser
    }
    
    if t.Name == "" {
        return ErrNameRequiredValidation
    }
    
    if !emailRegex.MatchString(t.Email) {
        return ErrEmailEmailValidation
    }
    
    if !(t.Age >= 0) {
        return ErrAgeGTEValidation
    }
    
    if !(t.Age <= 120) {
        return ErrAgeLTEValidation
    }
    
    if utf8.RuneCountInString(t.Bio) > 500 {
        return ErrBioMaxLengthValidation
    }
    
    return nil
}
```

### 3. Use Generated Validation

```go
func main() {
    user := &User{
        Name:  "Alice Johnson",
        Email: "alice@example.com",
        Age:   28,
        Bio:   "Software engineer passionate about Go",
    }
    
    // Validate the user
    if err := ValidateUser(user); err != nil {
        log.Fatalf("Validation failed: %v", err)
    }
    
    fmt.Println("User is valid!")
}
```

## Error Handling

govalid generates descriptive error messages and variables:

```go
// Generated error variables
var (
    ErrNilUser                    = errors.New("User is nil")

	// ErrUserNameRequiredValidation is returned when the Name is required but not provided.
	ErrUserNameRequiredValidation = govaliderrors.ValidationError{Reason: "field Name is required", Path: "User.Name", Type: "required"}

	// ErrUserEmailEmailValidation is the error returned when the field is not a valid email address.
	ErrUserEmailEmailValidation = govaliderrors.ValidationError{Reason: "field Email must be a valid email address", Path: "User.Email", Type: "email"}

	// ErrUserAgeGTEValidation is the error returned when the value of the field is less than 0.
	ErrUserAgeGTEValidation = govaliderrors.ValidationError{Reason: "field Age must be greater than or equal to 0", Path: "User.Age", Type: "gte"}

	// ErrUserAgeLTEValidation is the error returned when the value of the field is greater than 120.
	ErrUserAgeLTEValidation = govaliderrors.ValidationError{Reason: "field Age must be less than or equal to 120", Path: "User.Age", Type: "lte"}

	// ErrUserBioMaxLengthValidation is the error returned when the length of the field exceeds the maximum of 500.
	ErrUserBioMaxLengthValidation = govaliderrors.ValidationError{Reason: "field Bio must have a maximum length of 500", Path: "User.Bio", Type: "maxlength"}
)

```

You can check for specific errors:

```go
if err := ValidateUser(user); err != nil {
    switch err {
    case ErrEmailEmailValidation:
        fmt.Println("Please provide a valid email address")
    case ErrAgeGTEValidation:
        fmt.Println("Age cannot be negative")
    default:
        fmt.Printf("Validation error: %v\n", err)
    }
}

```

If multiple errors in are found during validation errors will be returned as a slice of structs that implement error interface.

```go

if err := ValidateUser(user); err != nil {
  var validationErrors govaliderrors.ValidationErrors
  if errors.As(err, &validationErrors) {
  	for _, e := range validationErrors {
  		log.Printf("Field %s: %s", e.Path, e.Reason)
  	}
  }
}
```

## Integration with Go Generate

Add a `go:generate` directive to automatically run govalid:

```go
//go:generate govalid .
package main

type User struct {
    // +govalid:required
    Name string `json:"name"`
}
```

Then run:

```bash
go generate ./...
```

## Best Practices

### 1. Organize Validation Rules

Group related validation rules together:

```go
type CreateUserRequest struct {
    // Basic required fields
    // +govalid:required
    // +govalid:minlength=2
    // +govalid:maxlength=50
    Name string `json:"name"`
    
    // Email validation
    // +govalid:required
    // +govalid:email
    Email string `json:"email"`
    
    // Age constraints
    // +govalid:gte=13
    // +govalid:lte=120
    Age int `json:"age"`
}
```

### 2. Use Descriptive Names

Choose clear, descriptive names for your structs and fields:

```go
type ProductCreateRequest struct {
    // +govalid:required
    // +govalid:minlength=3
    // +govalid:maxlength=100
    ProductName string `json:"product_name"`
    
    // +govalid:required
    // +govalid:gt=0
    Price float64 `json:"price"`
}
```

### 3. Combine with Standard Library

govalid works well with other Go validation patterns:

```go
func CreateUser(req *CreateUserRequest) error {
    // First, validate the struct
    if err := ValidateCreateUserRequest(req); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    // Then, perform business logic validation
    if userExists(req.Email) {
        return errors.New("user already exists")
    }
    
    // Create the user
    return createUser(req)
}
```

## Next Steps

- Check out all [available validators](/validators/)
- See [performance benchmarks](/benchmarks/)
- Browse [example implementations](/examples/)
- View the [source code](https://github.com/sivchari/govalid)
