---
title: "Examples"
description: "Real-world examples of using govalid"
weight: 40
---

# Examples

This page provides practical examples of using govalid in real-world scenarios.

## User Registration System

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

//go:generate govalid .

type CreateUserRequest struct {
    // +govalid:required
    // +govalid:minlength=2
    // +govalid:maxlength=50
    Name string `json:"name"`
    
    // +govalid:required
    // +govalid:email
    Email string `json:"email"`
    
    // +govalid:required
    // +govalid:minlength=8
    // +govalid:maxlength=100
    Password string `json:"password"`
    
    // +govalid:gte=13
    // +govalid:lte=120
    Age int `json:"age"`
    
    // +govalid:enum=admin,user,guest
    Role string `json:"role"`
    
    // +govalid:maxitems=10
    Tags []string `json:"tags,omitempty"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // Validate the request using generated function
    if err := ValidateCreateUserRequest(&req); err != nil {
        http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
        return
    }
    
    // Process the valid request
    user := createUser(req)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func main() {
    http.HandleFunc("/users", createUserHandler)
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## E-commerce Product Catalog

```go
package catalog

//go:generate govalid .

type Product struct {
    // +govalid:required
    // +govalid:uuid
    ID string `json:"id"`
    
    // +govalid:required
    // +govalid:minlength=3
    // +govalid:maxlength=200
    Name string `json:"name"`
    
    // +govalid:required
    // +govalid:maxlength=2000
    Description string `json:"description"`
    
    // +govalid:required
    // +govalid:gt=0
    Price float64 `json:"price"`
    
    // +govalid:required
    // +govalid:enum=electronics,clothing,books,home,sports
    Category string `json:"category"`
    
    // +govalid:required
    // +govalid:minitems=1
    // +govalid:maxitems=10
    Images []string `json:"images"`
    
    // +govalid:maxitems=20
    Tags []string `json:"tags,omitempty"`
    
    // +govalid:gte=0
    Stock int `json:"stock"`
    
    // +govalid:required
    // +govalid:enum=draft,active,inactive,discontinued
    Status string `json:"status"`
    
    // +govalid:maxitems=50
    Attributes map[string]string `json:"attributes,omitempty"`
}

type ProductFilter struct {
    // +govalid:enum=electronics,clothing,books,home,sports
    Category string `json:"category,omitempty"`
    
    // +govalid:gte=0
    MinPrice float64 `json:"min_price,omitempty"`
    
    // +govalid:gte=0
    MaxPrice float64 `json:"max_price,omitempty"`
    
    // +govalid:maxitems=10
    Tags []string `json:"tags,omitempty"`
    
    // +govalid:gte=1
    // +govalid:lte=100
    Limit int `json:"limit,omitempty"`
    
    // +govalid:gte=0
    Offset int `json:"offset,omitempty"`
}

func CreateProduct(product *Product) error {
    if err := ValidateProduct(product); err != nil {
        return fmt.Errorf("invalid product: %w", err)
    }
    
    // Business logic validation
    if product.Price > 10000 && product.Category == "books" {
        return errors.New("books cannot cost more than $10,000")
    }
    
    return saveProduct(product)
}

func SearchProducts(filter *ProductFilter) ([]*Product, error) {
    if err := ValidateProductFilter(filter); err != nil {
        return nil, fmt.Errorf("invalid filter: %w", err)
    }
    
    return queryProducts(filter)
}
```

## Configuration Management

```go
package config

import (
    "fmt"
    "os"
    "strconv"
)

//go:generate govalid .

type DatabaseConfig struct {
    // +govalid:required
    // +govalid:minlength=3
    Host string `json:"host"`
    
    // +govalid:required
    // +govalid:gte=1
    // +govalid:lte=65535
    Port int `json:"port"`
    
    // +govalid:required
    // +govalid:minlength=1
    Database string `json:"database"`
    
    // +govalid:required
    // +govalid:minlength=1
    Username string `json:"username"`
    
    // +govalid:required
    // +govalid:minlength=8
    Password string `json:"password"`
    
    // +govalid:gte=1
    // +govalid:lte=1000
    MaxConnections int `json:"max_connections"`
    
    // +govalid:required
    // +govalid:enum=disable,require,verify-ca,verify-full
    SSLMode string `json:"ssl_mode"`
}

type RedisConfig struct {
    // +govalid:required
    // +govalid:minlength=3
    Host string `json:"host"`
    
    // +govalid:required
    // +govalid:gte=1
    // +govalid:lte=65535
    Port int `json:"port"`
    
    // +govalid:minlength=1
    Password string `json:"password,omitempty"`
    
    // +govalid:gte=0
    // +govalid:lte=15
    Database int `json:"database"`
    
    // +govalid:gte=1
    // +govalid:lte=1000
    PoolSize int `json:"pool_size"`
}

type ServerConfig struct {
    // +govalid:required
    // +govalid:gte=1
    // +govalid:lte=65535
    Port int `json:"port"`
    
    // +govalid:required
    // +govalid:enum=debug,info,warn,error
    LogLevel string `json:"log_level"`
    
    // +govalid:required
    // +govalid:enum=development,staging,production
    Environment string `json:"environment"`
    
    // +govalid:gte=1
    // +govalid:lte=300
    TimeoutSeconds int `json:"timeout_seconds"`
    
    // +govalid:required
    // +govalid:minitems=1
    // +govalid:maxitems=10
    AllowedOrigins []string `json:"allowed_origins"`
}

type AppConfig struct {
    // +govalid:required
    Database DatabaseConfig `json:"database"`
    
    // +govalid:required
    Redis RedisConfig `json:"redis"`
    
    // +govalid:required
    Server ServerConfig `json:"server"`
    
    // +govalid:required
    // +govalid:minlength=32
    // +govalid:maxlength=256
    JWTSecret string `json:"jwt_secret"`
    
    // +govalid:required
    // +govalid:url
    BaseURL string `json:"base_url"`
    
    // +govalid:maxitems=100
    FeatureFlags map[string]bool `json:"feature_flags,omitempty"`
}

func LoadConfig() (*AppConfig, error) {
    config := &AppConfig{
        Database: DatabaseConfig{
            Host:           getEnv("DB_HOST", "localhost"),
            Port:           getEnvInt("DB_PORT", 5432),
            Database:       getEnv("DB_NAME", "myapp"),
            Username:       getEnv("DB_USER", ""),
            Password:       getEnv("DB_PASSWORD", ""),
            MaxConnections: getEnvInt("DB_MAX_CONNECTIONS", 25),
            SSLMode:        getEnv("DB_SSL_MODE", "disable"),
        },
        Redis: RedisConfig{
            Host:     getEnv("REDIS_HOST", "localhost"),
            Port:     getEnvInt("REDIS_PORT", 6379),
            Password: getEnv("REDIS_PASSWORD", ""),
            Database: getEnvInt("REDIS_DB", 0),
            PoolSize: getEnvInt("REDIS_POOL_SIZE", 10),
        },
        Server: ServerConfig{
            Port:           getEnvInt("SERVER_PORT", 8080),
            LogLevel:       getEnv("LOG_LEVEL", "info"),
            Environment:    getEnv("ENVIRONMENT", "development"),
            TimeoutSeconds: getEnvInt("TIMEOUT_SECONDS", 30),
            AllowedOrigins: getEnvSlice("ALLOWED_ORIGINS", []string{"*"}),
        },
        JWTSecret: getEnv("JWT_SECRET", ""),
        BaseURL:   getEnv("BASE_URL", "http://localhost:8080"),
    }
    
    // Validate the entire configuration
    if err := ValidateAppConfig(config); err != nil {
        return nil, fmt.Errorf("invalid configuration: %w", err)
    }
    
    return config, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if parsed, err := strconv.Atoi(value); err == nil {
            return parsed
        }
    }
    return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
    if value := os.Getenv(key); value != "" {
        return strings.Split(value, ",")
    }
    return defaultValue
}
```

## API Response Validation

```go
package api

//go:generate govalid .

type PaginationRequest struct {
    // +govalid:gte=1
    // +govalid:lte=100
    Page int `json:"page" form:"page"`
    
    // +govalid:gte=1
    // +govalid:lte=100
    PerPage int `json:"per_page" form:"per_page"`
    
    // +govalid:maxlength=100
    Query string `json:"query" form:"query"`
    
    // +govalid:enum=asc,desc
    SortOrder string `json:"sort_order" form:"sort_order"`
    
    // +govalid:maxlength=50
    SortBy string `json:"sort_by" form:"sort_by"`
}

type ErrorResponse struct {
    // +govalid:required
    // +govalid:minlength=1
    Message string `json:"message"`
    
    // +govalid:required
    // +govalid:gte=400
    // +govalid:lte=599
    Code int `json:"code"`
    
    // +govalid:maxitems=50
    Details []string `json:"details,omitempty"`
    
    // +govalid:required
    // +govalid:uuid
    RequestID string `json:"request_id"`
}

type SuccessResponse struct {
    // +govalid:required
    // +govalid:minlength=1
    Message string `json:"message"`
    
    // +govalid:required
    // +govalid:gte=200
    // +govalid:lte=299
    Code int `json:"code"`
    
    // Data can be any type
    Data interface{} `json:"data,omitempty"`
    
    // +govalid:required
    // +govalid:uuid
    RequestID string `json:"request_id"`
}

func ValidateAndParsePagination(r *http.Request) (*PaginationRequest, error) {
    pagination := &PaginationRequest{
        Page:      1,
        PerPage:   10,
        SortOrder: "asc",
    }
    
    if page := r.URL.Query().Get("page"); page != "" {
        if p, err := strconv.Atoi(page); err == nil {
            pagination.Page = p
        }
    }
    
    if perPage := r.URL.Query().Get("per_page"); perPage != "" {
        if pp, err := strconv.Atoi(perPage); err == nil {
            pagination.PerPage = pp
        }
    }
    
    pagination.Query = r.URL.Query().Get("query")
    pagination.SortOrder = r.URL.Query().Get("sort_order")
    pagination.SortBy = r.URL.Query().Get("sort_by")
    
    if err := ValidatePaginationRequest(pagination); err != nil {
        return nil, err
    }
    
    return pagination, nil
}
```

## Custom Type Validation

```go
package types

//go:generate govalid .

type UserRole string
type Priority int
type Status string

const (
    RoleAdmin UserRole = "admin"
    RoleUser  UserRole = "user"
    RoleGuest UserRole = "guest"
)

const (
    PriorityLow    Priority = 1
    PriorityMedium Priority = 2
    PriorityHigh   Priority = 3
)

type Task struct {
    // +govalid:required
    // +govalid:uuid
    ID string `json:"id"`
    
    // +govalid:required
    // +govalid:minlength=3
    // +govalid:maxlength=200
    Title string `json:"title"`
    
    // +govalid:maxlength=2000
    Description string `json:"description,omitempty"`
    
    // +govalid:required
    // +govalid:enum=admin,user,guest
    AssignedToRole UserRole `json:"assigned_to_role"`
    
    // +govalid:required
    // +govalid:enum=1,2,3
    Priority Priority `json:"priority"`
    
    // +govalid:required
    // +govalid:enum=pending,in_progress,completed,cancelled
    Status Status `json:"status"`
    
    // +govalid:minitems=1
    // +govalid:maxitems=10
    Tags []string `json:"tags"`
    
    // +govalid:maxitems=20
    Metadata map[string]string `json:"metadata,omitempty"`
    
    // +govalid:gte=1
    // +govalid:lte=100
    EstimatedHours int `json:"estimated_hours,omitempty"`
    
    // +govalid:required
    // +govalid:email
    CreatedBy string `json:"created_by"`
}

func CreateTask(task *Task) error {
    if err := ValidateTask(task); err != nil {
        return fmt.Errorf("invalid task: %w", err)
    }
    
    // Business logic validation
    if task.Priority == PriorityHigh && task.AssignedToRole == RoleGuest {
        return errors.New("high priority tasks cannot be assigned to guests")
    }
    
    return saveTask(task)
}
```

## Integration with Middleware

```go
package middleware

import (
    "encoding/json"
    "net/http"
    "reflect"
)

func ValidateJSON(v interface{}) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Create a new instance of the validation type
            instance := reflect.New(reflect.TypeOf(v).Elem()).Interface()
            
            // Decode JSON
            if err := json.NewDecoder(r.Body).Decode(instance); err != nil {
                http.Error(w, "Invalid JSON", http.StatusBadRequest)
                return
            }
            
            // Use reflection to call the validation function
            validateFunc := getValidationFunction(instance)
            if validateFunc != nil {
                if err := validateFunc(instance); err != nil {
                    http.Error(w, err.Error(), http.StatusBadRequest)
                    return
                }
            }
            
            // Store validated data in context
            ctx := context.WithValue(r.Context(), "validated_data", instance)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func getValidationFunction(v interface{}) func(interface{}) error {
    typeName := reflect.TypeOf(v).Elem().Name()
    
    switch typeName {
    case "CreateUserRequest":
        return func(i interface{}) error {
            return ValidateCreateUserRequest(i.(*CreateUserRequest))
        }
    case "Product":
        return func(i interface{}) error {
            return ValidateProduct(i.(*Product))
        }
    // Add more cases as needed
    default:
        return nil
    }
}

// Usage:
// router.Use(ValidateJSON(&CreateUserRequest{}))
```

These examples demonstrate how govalid can be used in real-world scenarios to provide fast, type-safe validation with zero allocations. The generated validation functions integrate seamlessly with existing Go code and provide clear error messages for invalid data.