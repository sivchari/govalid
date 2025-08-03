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

	for i := range e {
		buff.WriteString(e[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// Is implements error matching for ValidationErrors.
// It checks if any of the contained errors match the target.
func (e ValidationErrors) Is(target error) bool {
	for _, err := range e {
		if err.Is(target) {
			return true
		}
	}

	return false
}

// Error implements the error interface for ValidationError.
// It returns a string representation of a single validation error.
func (e ValidationError) Error() string {
	return fmt.Errorf(
		"field %s with value %v has failed validation %s because %s",
		e.Path, e.Value, e.Type, e.Reason,
	).Error()
}

// Is implements error matching for ValidationError.
// It allows errors.Is to work with ValidationError instances.
func (e ValidationError) Is(target error) bool {
	if ve, ok := target.(ValidationError); ok {
		return e.Path == ve.Path && e.Type == ve.Type && e.Reason == ve.Reason
	}

	return false
}
