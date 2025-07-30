// Package errors provides structures for handling validation errors.
package errors

import (
	"fmt"
	"strings"
)

// ValidationError represents a single validation error.
type ValidationError struct {
	// Path is the path to the field that failed validation, e.g., "Name", "Address.City", "Users[0].Email".
	Path string
	// Type is the type of validation that failed, e.g., "required", "email", "maxlength".
	Type string
	// Value is the actual value that was validated.
	Value any
	// Reason is a human-readable message explaining why the validation failed.
	Reason string
}

// ValidationErrors is a slice of ValidationError, representing a collection of validation errors.
type ValidationErrors []ValidationError

// Error implements the error interface for ValidationErrors.
// It returns a string representation of all validation errors, separated by newlines.
func (e ValidationErrors) Error() string {
	buff := strings.Builder{}

	for i := range len(e) {
		buff.WriteString(e[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// Error implements the error interface for ValidationError.
// It returns a string representation of a single validation error.
func (e ValidationError) Error() string {
	return fmt.Errorf(
		"field %s with value %v has failed validation %s because %s",
		e.Path, e.Value, e.Type, e.Reason,
	).Error()
}
