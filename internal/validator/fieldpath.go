// Package validator implements rules for validating fields.
package validator

import "strings"

// FieldPath represents a field path with dot-separated components.
type FieldPath string

// NewFieldPath creates a new FieldPath from the given components.
func NewFieldPath(components ...string) FieldPath {
	nonEmpty := make([]string, 0, len(components))

	for _, component := range components {
		if strings.TrimSpace(component) != "" {
			nonEmpty = append(nonEmpty, component)
		}
	}

	return FieldPath(strings.Join(nonEmpty, "."))
}

// CleanedPath returns the field path with all dots removed for use in variable names and keys.
func (fp FieldPath) CleanedPath() string {
	return strings.ReplaceAll(string(fp), ".", "")
}

// String returns the string representation of the field path.
func (fp FieldPath) String() string {
	return string(fp)
}
