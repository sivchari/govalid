package errors

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Path   string // "Name", "Address.City", "Users[0].Email"
	Type   string // "required", "email", "maxlength"
	Value  any    // The actual value being validated
	Reason string // "field is required"
}

type ValidationErrors []ValidationError

// Implement error interface
func (e ValidationErrors) Error() string {
	buff := strings.Builder{}

	for i := range len(e) {
		buff.WriteString(e[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// Implement error interface
func (e ValidationError) Error() string {
	return fmt.Errorf(
		"Field %s with value %v has failed validation %s because %s",
		e.Path, e.Value, e.Type, e.Reason,
	).Error()
}
