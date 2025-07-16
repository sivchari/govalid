---
title: "govalid"
description: "High-performance, type-safe validation library for Go with zero allocations"
---

# govalid

**High-performance, type-safe validation library for Go with zero allocations**

[![GitHub](https://img.shields.io/badge/GitHub-sivchari/govalid-blue?logo=github)](https://github.com/sivchari/govalid)
[![Go Report Card](https://goreportcard.com/badge/github.com/sivchari/govalid)](https://goreportcard.com/report/github.com/sivchari/govalid)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

govalid is a Go package designed to generate type-safe validation code for structs based on markers. It provides a mechanism to apply validation rules directly in the code by marking struct fields with specific markers. The tool processes these markers and generates corresponding validation functions with **zero heap allocations**.

## Key Features

- **🚀 Zero Allocations**: All validators perform zero heap allocations
- **🔒 Type Safety**: Compile-time validation function generation
- **⚡ High Performance**: 5x to 44x faster than reflection-based validators
- **📝 Marker-Based**: Simple comment-based validation rules
- **🔧 Code Generation**: Generates optimized validation functions
- **🎯 Comprehensive**: Support for all common validation patterns

## Quick Start

### Installation

```bash
go install github.com/sivchari/govalid/cmd/govalid@latest
```

### Basic Usage

1. **Define your struct with markers:**

```go
package main

type User struct {
    // +govalid:required
    Name string `json:"name"`
    
    // +govalid:email
    Email string `json:"email"`
    
    // +govalid:gte=18
    Age int `json:"age"`
    
    // +govalid:maxlength=100
    Bio string `json:"bio"`
}
```

2. **Generate validation code:**

```bash
govalid ./...
```

3. **Use the generated validation:**

```go
func main() {
    user := &User{
        Name:  "Alice",
        Email: "alice@example.com",
        Age:   25,
        Bio:   "Software developer",
    }
    
    if err := ValidateUser(user); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("User is valid!")
}
```

## Performance Comparison

| Validator | govalid | go-playground | go-validator | ozzo-validation | Best Improvement |
|-----------|---------|---------------|--------------|-----------------|------------------|
| Required  | 1.904 | 84.59 | **44.4x faster** |
| Email     | 37.65 | 633.5 | **16.8x faster** |
| GT/LT     | ~1.888 | ~61.09 | **32.3x faster** |
| MaxLength | 15.37 | 72.60 | **4.7x faster** |
| Enum      | 2.227 | N/A (unique to govalid)| **govalid exclusive** |

*All benchmarks show 0 allocations for govalid vs 0-80 allocations for competitors*

**Methodology**: 10 runs per benchmark analyzed with `benchstat` for statistical significance

## Supported Validators

### String Validators
- `govalid:required` - Required field validation
- `govalid:minlength=N` - Minimum string length (Unicode-aware)
- `govalid:maxlength=N` - Maximum string length (Unicode-aware)
- `govalid:email` - HTML5-compliant email validation
- `govalid:url` - HTTP/HTTPS URL validation
- `govalid:uuid` - RFC 4122 UUID validation

### Numeric Validators
- `govalid:gt=N` - Greater than validation
- `govalid:gte=N` - Greater than or equal validation
- `govalid:lt=N` - Less than validation
- `govalid:lte=N` - Less than or equal validation

### Collection Validators
- `govalid:minitems=N` - Minimum collection size
- `govalid:maxitems=N` - Maximum collection size

### General Validators
- `govalid:enum=val1,val2,val3` - Enum validation

## Why govalid?

### 🔥 Performance First
- **Zero allocations**: No heap allocations during validation
- **Compile-time optimization**: Generated code is optimized by the Go compiler
- **Minimal overhead**: Direct field access with no reflection

### 🎯 Developer Experience
- **Familiar syntax**: Similar to popular validation libraries
- **Type safety**: Catch validation errors at compile time
- **Clear error messages**: Descriptive validation error messages

### 🚀 Production Ready
- **Extensive test coverage**: Comprehensive unit and benchmark tests
- **Battle tested**: Used in production environments
- **Maintained**: Regular updates and improvements

## Get Started

Ready to make your Go validation blazingly fast? Check out our [Getting Started](/getting-started/) guide or browse the [available validators](/validators/).

[Get Started →](/getting-started/) [View on GitHub →](https://github.com/sivchari/govalid)
