// Package validator implements rules for validating fields.
package validator

// Validator is an interface for validating fields in structs.
type Validator interface {
	Validate() string
	FieldName() string
	FieldPath() FieldPath
	Err() string
	ErrVariable() string
	Imports() []string
}

// GeneratorMemory is a map used to track the state of generated validators.
var GeneratorMemory = map[string]bool{}
