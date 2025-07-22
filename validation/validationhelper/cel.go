// Package validationhelper provides validation helper functions for govalid.
package validationhelper

import (
	"fmt"
	"sync"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
)

// celCache caches compiled CEL programs to avoid recompilation.
// sync.Map is used for better performance in read-heavy scenarios.
var celCache sync.Map

// IsValidCEL evaluates a CEL expression for validation.
// This function uses a simple approach without reflection, following govalid's design principles.
// Compiled CEL programs are cached for performance.
//
// Parameters:
//   - expression: The CEL expression string to evaluate
//   - value: The current field value being validated
//   - structInstance: The entire struct instance (reserved for future use)
//
// Returns:
//   - bool: true if validation passes, false if validation fails
//
// Example usage:
//
//	if !IsValidCEL("value > 0.0", score, instance) {
//	    return errors.New("score must be positive")
//	}
//
// Note: This implementation prioritizes simplicity and performance,
// following govalid's zero-reflection philosophy. Cross-field validation
// is not supported without reflection.
func IsValidCEL(expression string, value, structInstance any) bool {
	// Try to get cached program first
	if cached, ok := celCache.Load(expression); ok {
		program, ok := cached.(cel.Program)
		if !ok {
			return false
		}

		return evaluateCELProgram(program, value, structInstance)
	}

	// Compile and cache the program
	compiledProgram, err := compileCELExpression(expression)
	if err != nil {
		return false
	}

	// Store in cache for future use
	celCache.Store(expression, compiledProgram)

	return evaluateCELProgram(compiledProgram, value, structInstance)
}

// evaluateCELProgram evaluates a compiled CEL program with given values.
func evaluateCELProgram(program cel.Program, value, structInstance any) bool {
	// Prepare evaluation variables
	// Note: Without reflection, we cannot support cross-field validation
	// that accesses struct fields. Only 'value' based validation is supported.
	vars := map[string]any{
		"value": value,
		"this":  structInstance,
	}

	// Evaluate the expression
	out, _, err := program.Eval(vars)
	if err != nil {
		return false
	}

	// Convert result to boolean
	switch result := out.Value().(type) {
	case bool:
		return result
	case types.Bool:
		return bool(result)
	default:
		// If not a boolean, consider it invalid
		return false
	}
}

// compileCELExpression compiles a CEL expression into a program.
// This function is used internally for compilation and caching.
func compileCELExpression(expression string) (cel.Program, error) {
	// Create a CEL environment with struct field access support
	env, err := cel.NewEnv(
		cel.StdLib(),
		// Enable dynamic type access for struct fields
		cel.Variable("value", cel.DynType),
		cel.Variable("this", cel.DynType),
		// Enable field access and type conversions
		cel.OptionalTypes(),
		cel.EnableMacroCallTracking(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create CEL environment: %w", err)
	}

	// Compile the expression
	ast, issues := env.Compile(expression)
	if issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("failed to compile CEL expression: %w", issues.Err())
	}

	// Create and return program with optimizations
	program, err := env.Program(ast, cel.EvalOptions(cel.OptOptimize))
	if err != nil {
		return nil, fmt.Errorf("failed to create CEL program: %w", err)
	}

	return program, nil
}
