// Package govalid provides type-safe validation code generation for structs based on markers.
package govalid

import "context"

// Validator is the interface that wraps the basic Validation method.
// This interface is implemented automatically by generated validation code to enable
// middleware and other consumers to validate structs polymorphically.
type Validator interface {
	Validate() error
}

// ContextValidator is the interface that extends Validator with context support.
// Use this interface when you need to pass context (e.g., for cancellation or timeout).
type ContextValidator interface {
	Validator
	ValidateContext(ctx context.Context) error
}
