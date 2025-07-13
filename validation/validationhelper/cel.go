// Package validationhelper provides validation helper functions for govalid.
package validationhelper

import (
	"fmt"
	"reflect"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

// CompiledCELExpression represents a pre-compiled CEL expression with metadata.
// This struct is used to store compiled CEL programs for efficient validation.
type CompiledCELExpression struct {
	Program cel.Program
	Env     *cel.Env
}

// CELValidationCache stores compiled CEL expressions to avoid recompilation.
// The key is the expression string, and the value is the compiled expression.
var CELValidationCache = make(map[string]*CompiledCELExpression)

// CompileCELExpression compiles a CEL expression with the given type information.
// This function should be called during code generation to pre-compile expressions.
//
// Parameters:
//   - expression: The CEL expression string to compile
//   - fieldType: The Go type of the field being validated
//   - structType: The Go type of the struct containing the field (for cross-field validation)
//
// Returns:
//   - *CompiledCELExpression: The compiled expression ready for evaluation
//   - error: Any compilation error
//
// Example usage during code generation:
//
//	compiled, err := CompileCELExpression("value >= 18 && value <= 120", reflect.TypeOf(int(0)), reflect.TypeOf(User{}))
//	if err != nil {
//	    // Handle compilation error
//	}
//	CELValidationCache["age_validation"] = compiled
func CompileCELExpression(expression string, fieldType, structType reflect.Type) (*CompiledCELExpression, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("%s_%s_%s", expression, fieldType.String(), structType.String())
	if cached, exists := CELValidationCache[cacheKey]; exists {
		return cached, nil
	}

	// Create CEL environment with type declarations
	env, err := createCELEnvironment(fieldType, structType)
	if err != nil {
		return nil, fmt.Errorf("failed to create CEL environment: %w", err)
	}

	// Compile the expression
	ast, issues := env.Compile(expression)
	if issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("CEL compilation error: %w", issues.Err())
	}

	// Create program
	program, err := env.Program(ast)
	if err != nil {
		return nil, fmt.Errorf("failed to create CEL program: %w", err)
	}

	compiled := &CompiledCELExpression{
		Program: program,
		Env:     env,
	}

	// Cache for future use
	CELValidationCache[cacheKey] = compiled

	return compiled, nil
}

// createCELEnvironment creates a CEL environment with 'value' and 'this' variables.
//
// Variables available in CEL expressions:
//   - value: The current field being validated
//   - this: The entire struct instance (for cross-field validation)
//
// The environment includes standard CEL functions and type declarations.
func createCELEnvironment(fieldType, structType reflect.Type) (*cel.Env, error) {
	var envOptions []cel.EnvOption

	// Add standard library functions
	envOptions = append(envOptions, cel.StdLib())

	// Declare 'value' variable with the field's type
	fieldCELType, err := goTypeToCELType(fieldType)
	if err != nil {
		return nil, fmt.Errorf("failed to convert field type to CEL type: %w", err)
	}
	envOptions = append(envOptions, cel.Variable("value", fieldCELType))

	// Declare 'this' variable with the struct's type (for cross-field validation)
	structCELType, err := goTypeToCELType(structType)
	if err != nil {
		return nil, fmt.Errorf("failed to convert struct type to CEL type: %w", err)
	}
	envOptions = append(envOptions, cel.Variable("this", structCELType))

	return cel.NewEnv(envOptions...)
}

// goTypeToCELType converts Go types to CEL types.
// This function maps common Go types to their CEL equivalents.
func goTypeToCELType(goType reflect.Type) (*cel.Type, error) {
	switch goType.Kind() {
	case reflect.Bool:
		return cel.BoolType, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cel.IntType, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return cel.UintType, nil
	case reflect.Float32, reflect.Float64:
		return cel.DoubleType, nil
	case reflect.String:
		return cel.StringType, nil
	case reflect.Slice:
		elemType, err := goTypeToCELType(goType.Elem())
		if err != nil {
			return nil, err
		}
		return cel.ListType(elemType), nil
	case reflect.Map:
		keyType, err := goTypeToCELType(goType.Key())
		if err != nil {
			return nil, err
		}
		valueType, err := goTypeToCELType(goType.Elem())
		if err != nil {
			return nil, err
		}
		return cel.MapType(keyType, valueType), nil
	case reflect.Struct:
		// For structs, we use DynType to allow field access
		return cel.DynType, nil
	case reflect.Ptr:
		// For pointers, use the underlying type
		return goTypeToCELType(goType.Elem())
	default:
		// For unsupported types, use DynType as fallback
		return cel.DynType, nil
	}
}

// EvaluateCEL evaluates a compiled CEL expression with the given values.
// This function is called during runtime validation to check if the expression evaluates to true.
//
// Parameters:
//   - compiled: The pre-compiled CEL expression
//   - value: The current field value being validated
//   - structInstance: The entire struct instance (for cross-field validation)
//
// Returns:
//   - bool: true if validation passes, false if validation fails
//   - error: Any evaluation error
//
// Example usage in generated validation code:
//
//	result, err := EvaluateCEL(compiled, t.Age, t)
//	if err != nil {
//	    return err
//	}
//	if !result {
//	    return ErrAgeValidation
//	}
func EvaluateCEL(compiled *CompiledCELExpression, value interface{}, structInstance interface{}) (bool, error) {
	// Prepare evaluation variables
	vars := map[string]interface{}{
		"value": value,
		"this":  structInstance,
	}

	// Evaluate the expression
	out, _, err := compiled.Program.Eval(vars)
	if err != nil {
		return false, fmt.Errorf("CEL evaluation error: %w", err)
	}

	// Convert result to boolean
	result, ok := out.Value().(bool)
	if !ok {
		return false, fmt.Errorf("CEL expression must evaluate to boolean, got %T", out.Value())
	}

	return result, nil
}

// IsValidCEL is a convenience function for simple CEL validation without pre-compilation.
// This function compiles and evaluates the expression in one call.
// Note: This should only be used for testing or simple cases where performance is not critical.
// For production use, prefer CompileCELExpression + EvaluateCEL pattern.
//
// Parameters:
//   - expression: The CEL expression string
//   - value: The current field value being validated
//   - structInstance: The entire struct instance
//
// Returns:
//   - bool: true if validation passes, false if validation fails or compilation fails
//
// Example usage:
//
//	if !IsValidCEL("value >= 18 && value <= 120", age, user) {
//	    return errors.New("age must be between 18 and 120")
//	}
func IsValidCEL(expression string, value interface{}, structInstance interface{}) bool {
	// Get types for compilation
	fieldType := reflect.TypeOf(value)
	structType := reflect.TypeOf(structInstance)

	// Compile expression
	compiled, err := CompileCELExpression(expression, fieldType, structType)
	if err != nil {
		return false
	}

	// Evaluate expression
	result, err := EvaluateCEL(compiled, value, structInstance)
	if err != nil {
		return false
	}

	return result
}

// ConvertValueToCEL converts Go values to CEL-compatible values.
// This function handles type conversions needed for CEL evaluation.
func ConvertValueToCEL(value interface{}) ref.Val {
	return types.DefaultTypeAdapter.NativeToValue(value)
}