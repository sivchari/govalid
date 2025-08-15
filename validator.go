// Package govalid provides type-safe validation code generation for structs based on markers.
package govalid

// Validator is the interface that wraps the basic Validate method.
// This interface is implemented automatically by generated validation code to enable
// middleware and other consumers to validate structs polymorphically.
type Validator interface {
	Validate() error
}
