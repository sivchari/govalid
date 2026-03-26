// Package validator implements rules for validating fields.
package validator

import "go/ast"

// Condition represents a validation condition with an optional if-init statement.
type Condition struct {
	// IfInitStmt is the initialization statement in Go's if-init syntax: if init; cond {}.
	// Example: ip := net.ParseIP(t.IP) (used by ipv4/ipv6 validators).
	// When nil, the condition becomes a simple if cond {}.
	IfInitStmt ast.Stmt

	// Expr is the condition expression that evaluates to true when validation fails.
	Expr ast.Expr

	// Imports is the list of import paths required by this condition.
	Imports []string
}

// ErrDecl represents a structured error variable declaration.
type ErrDecl struct {
	// VarName is the Go identifier for the error variable.
	// Example: "ErrUserNameGTValidation"
	VarName string

	// Comment is the documentation comment for the error variable.
	// Example: "is the error returned when the value of the field is less than the 5."
	Comment string

	// Reason is the human-readable error message included in the ValidationError.
	// Example: "field Name must be greater than 5"
	Reason string

	// Path is the dot-separated field path in the struct hierarchy.
	// Example: "User.Name"
	Path string

	// Type is the validation rule name.
	// Example: "gt"
	Type string
}

// Validator is an interface for validating fields in structs.
type Validator interface {
	// Condition returns the validation condition as an AST-based representation.
	// Returns nil if no validation is needed.
	Condition() *Condition

	// FieldName returns the name of the field being validated.
	FieldName() string

	// FieldPath returns the dot-separated path of the field in the struct hierarchy.
	FieldPath() FieldPath

	// ErrDecl returns the structured error variable declaration.
	ErrDecl() ErrDecl
}
