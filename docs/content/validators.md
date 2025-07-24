---
title: "Validators"
description: "Complete reference for all govalid validators"
weight: 20
---

# Validators Reference

This page provides a comprehensive reference for all validators available in govalid.

## String Validators

### `govalid:required`

Ensures that the field is not empty or nil.

**Supported Types:** `string`, `slice`, `map`, `chan`, `pointer`

**Example:**
```go
type User struct {
    // +govalid:required
    Name string `json:"name"`
    
    // +govalid:required
    Tags []string `json:"tags"`
}
```

**Generated Code:**
```go
if t.Name == "" {
    return ErrNameRequiredValidation
}

if len(t.Tags) == 0 {
    return ErrTagsRequiredValidation
}
```

### `govalid:minlength=N`

Ensures that a string field's length is at least the specified minimum value (Unicode-aware).

**Supported Types:** `string`

**Example:**
```go
type User struct {
    // +govalid:minlength=3
    Username string `json:"username"`
}
```

**Generated Code:**
```go
if utf8.RuneCountInString(t.Username) < 3 {
    return ErrUsernameMinLengthValidation
}
```

### `govalid:maxlength=N`

Ensures that a string field's length does not exceed the specified maximum value (Unicode-aware).

**Supported Types:** `string`

**Example:**
```go
type User struct {
    // +govalid:maxlength=50
    Username string `json:"username"`
}
```

**Generated Code:**
```go
if utf8.RuneCountInString(t.Username) > 50 {
    return ErrUsernameMaxLengthValidation
}
```

### `govalid:length=N`

Ensures that a string field has exactly the specified length (Unicode-aware).

**Supported Types:** `string`

**Example:**
```go
type User struct {
    // +govalid:length=7
    Name string `json:"name"`
}
```

**Generated Code:**
```go
if utf8.RuneCountInString(t.Name) != 7 {
    return ErrNameLengthValidation
}
```

**Note:** Unlike `minlength` and `maxlength`, this validator requires the exact character count. It's perfect for fixed-length fields like postal codes, phone numbers, or product codes.

### `govalid:email`

Ensures that a string field is a valid email address using HTML5-compliant validation.

**Supported Types:** `string`

**Example:**
```go
type User struct {
    // +govalid:email
    Email string `json:"email"`
}
```

**Generated Code:**
```go
if !emailRegex.MatchString(t.Email) {
    return ErrEmailEmailValidation
}
```

### `govalid:url`

Ensures that a string field is a valid URL using HTTP/HTTPS protocol validation.

**Supported Types:** `string`

**Example:**
```go
type Resource struct {
    // +govalid:url
    URL string `json:"url"`
}
```

**Generated Code:**
```go
if !validationhelper.IsValidURL(t.URL) {
    return ErrURLURLValidation
}
```

### `govalid:uuid`

Ensures that a string field is a valid UUID following RFC 4122 format.

**Supported Types:** `string`

**Example:**
```go
type Resource struct {
    // +govalid:uuid
    ID string `json:"id"`
}
```

**Generated Code:**
```go
if !isValidUUID(t.ID) {
    return ErrIDUUIDValidation
}
```

## Numeric Validators

### `govalid:gt=N`

Ensures that a numeric field is greater than a specified value.

**Supported Types:** `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`

**Example:**
```go
type Profile struct {
    // +govalid:gt=0
    Age int `json:"age"`
    
    // +govalid:gt=0.0
    Salary float64 `json:"salary"`
}
```

**Generated Code:**
```go
if t.Age <= 0 {
    return ErrAgeGtValidation
}

if t.Salary <= 0.0 {
    return ErrSalaryGtValidation
}
```

### `govalid:gte=N`

Ensures that a numeric field is greater than or equal to a specified value.

**Supported Types:** `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`

**Example:**
```go
type Profile struct {
    // +govalid:gte=18
    Age int `json:"age"`
}
```

**Generated Code:**
```go
if !(t.Age >= 18) {
    return ErrAgeGTEValidation
}
```

### `govalid:lt=N`

Ensures that a numeric field is less than a specified value.

**Supported Types:** `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`

**Example:**
```go
type Profile struct {
    // +govalid:lt=65
    Age int `json:"age"`
}
```

**Generated Code:**
```go
if t.Age >= 65 {
    return ErrAgeLtValidation
}
```

### `govalid:lte=N`

Ensures that a numeric field is less than or equal to a specified value.

**Supported Types:** `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`

**Example:**
```go
type Profile struct {
    // +govalid:lte=120
    Age int `json:"age"`
}
```

**Generated Code:**
```go
if !(t.Age <= 120) {
    return ErrAgeLTEValidation
}
```

## Collection Validators

### `govalid:minitems=N`

Ensures that a collection field's length is at least the specified minimum number of items.

**Supported Types:** `slice`, `array`, `map`, `chan`

**Example:**
```go
type Collection struct {
    // +govalid:minitems=1
    Items []string `json:"items"`
    
    // +govalid:minitems=2
    Tags []string `json:"tags"`
    
    // +govalid:minitems=1
    Metadata map[string]string `json:"metadata"`
}
```

**Generated Code:**
```go
if len(t.Items) < 1 {
    return ErrItemsMinItemsValidation
}

if len(t.Tags) < 2 {
    return ErrTagsMinItemsValidation
}

if len(t.Metadata) < 1 {
    return ErrMetadataMinItemsValidation
}
```

### `govalid:maxitems=N`

Ensures that a collection field's length does not exceed the specified maximum number of items.

**Supported Types:** `slice`, `array`, `map`, `chan`

**Example:**
```go
type Collection struct {
    // +govalid:maxitems=10
    Items []string `json:"items"`
    
    // +govalid:maxitems=5
    Metadata map[string]string `json:"metadata"`
}
```

**Generated Code:**
```go
if len(t.Items) > 10 {
    return ErrItemsMaxItemsValidation
}

if len(t.Metadata) > 5 {
    return ErrMetadataMaxItemsValidation
}
```

## General Validators

### `govalid:enum=val1,val2,val3`

Ensures that a field value is within a specified set of allowed values.

**Supported Types:** `string`, numeric types, custom types with comparable values

**Example:**
```go
type User struct {
    // +govalid:enum=admin,user,guest
    Role string `json:"role"`
    
    // +govalid:enum=1,2,3
    Level int `json:"level"`
    
    // +govalid:enum=active,inactive,pending
    Status string `json:"status"`
}

type UserRole string
type Priority int

type ExtendedUser struct {
    // +govalid:enum=manager,developer,tester
    UserRole UserRole `json:"user_role"`
    
    // +govalid:enum=10,20,30
    Priority Priority `json:"priority"`
}
```

**Generated Code:**
```go
if t.Role != "admin" && t.Role != "user" && t.Role != "guest" {
    return ErrRoleEnumValidation
}

if t.Level != 1 && t.Level != 2 && t.Level != 3 {
    return ErrLevelEnumValidation
}

if t.Status != "active" && t.Status != "inactive" && t.Status != "pending" {
    return ErrStatusEnumValidation
}
```

## Combining Validators

You can combine multiple validators on a single field:

```go
type User struct {
    // +govalid:required
    // +govalid:minlength=3
    // +govalid:maxlength=50
    Username string `json:"username"`
    
    // +govalid:required
    // +govalid:email
    Email string `json:"email"`
    
    // +govalid:gte=18
    // +govalid:lte=120
    Age int `json:"age"`
    
    // +govalid:required
    // +govalid:enum=admin,user,guest
    Role string `json:"role"`
}
```

This generates validation code that checks all specified rules:

```go
func ValidateUser(t *User) error {
    if t == nil {
        return ErrNilUser
    }
    
    if t.Username == "" {
        return ErrUsernameRequiredValidation
    }
    
    if utf8.RuneCountInString(t.Username) < 3 {
        return ErrUsernameMinLengthValidation
    }
    
    if utf8.RuneCountInString(t.Username) > 50 {
        return ErrUsernameMaxLengthValidation
    }
    
    if t.Email == "" {
        return ErrEmailRequiredValidation
    }
    
    if !emailRegex.MatchString(t.Email) {
        return ErrEmailEmailValidation
    }
    
    if !(t.Age >= 18) {
        return ErrAgeGTEValidation
    }
    
    if !(t.Age <= 120) {
        return ErrAgeLTEValidation
    }
    
    if t.Role == "" {
        return ErrRoleRequiredValidation
    }
    
    if t.Role != "admin" && t.Role != "user" && t.Role != "guest" {
        return ErrRoleEnumValidation
    }
    
    return nil
}
```

## Performance Characteristics

All govalid validators are designed for optimal performance:

- **Zero allocations**: No heap allocations during validation
- **Inlined code**: Simple validators are inlined by the compiler
- **Minimal overhead**: Direct field access with no reflection
- **Optimized patterns**: Hand-tuned validation logic for common cases

See the [benchmarks page](/benchmarks/) for detailed performance comparisons.