// Package govalid provides type-safe validation code generation for structs based on markers.
package govalid

// Validator interface for structs that can validate themselves.
// This interface is implemented by generated validation code to enable
// middleware and other consumers to validate structs polymorphically.
type Validator interface {
	Validate() error
}
