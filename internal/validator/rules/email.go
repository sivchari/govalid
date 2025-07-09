// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator"
)

type emailValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ validator.Validator = (*emailValidator)(nil)

const emailKey = "%s-email"

func (e *emailValidator) Validate() string {
	fieldName := e.FieldName()
	// Use external helper function for better maintainability
	return fmt.Sprintf("!validationhelper.IsValidEmail(t.%s)", fieldName)
}

func (e *emailValidator) FieldName() string {
	return e.field.Names[0].Name
}

func (e *emailValidator) Err() string {
	fieldName := e.FieldName()

	var result strings.Builder

	// No need to generate inline function - using external helper
	if validator.GeneratorMemory[fmt.Sprintf(emailKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(emailKey, fieldName)] = true

	result.WriteString(strings.ReplaceAll(`
	// Err@EmailValidation is the error returned when the field is not a valid email address.
	Err@EmailValidation = errors.New("field @ must be a valid email address")`, `@`, fieldName))

	return result.String()
}

func (e *emailValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@EmailValidation", `@`, e.FieldName())
}

func (e *emailValidator) Imports() []string {
	// Import validation helper package
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateEmail creates a new emailValidator for string types.
func ValidateEmail(pass *codegen.Pass, field *ast.Field, _ map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &emailValidator{
		pass:  pass,
		field: field,
	}
}
