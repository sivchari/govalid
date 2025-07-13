// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
)

type celValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	expression string
}

var _ validator.Validator = (*celValidator)(nil)

const celKey = "%s-cel"

func (c *celValidator) Validate() string {
	fieldName := c.FieldName()
	// Use external helper function for CEL evaluation
	// The helper function will handle expression compilation and evaluation
	return fmt.Sprintf("!validationhelper.IsValidCEL(%q, t.%s, t)", c.expression, fieldName)
}

func (c *celValidator) FieldName() string {
	return c.field.Names[0].Name
}

func (c *celValidator) Err() string {
	fieldName := c.FieldName()

	var result strings.Builder

	// Generate error variable only once per field
	if validator.GeneratorMemory[fmt.Sprintf(celKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(celKey, fieldName)] = true

	// Generate error variable with the CEL expression included for debugging
	result.WriteString(strings.ReplaceAll(`
	// Err@CELValidation is the error returned when the CEL expression evaluation fails.
	Err@CELValidation = errors.New("field @ failed CEL validation: EXPRESSION")`, "@", fieldName))

	// Replace EXPRESSION placeholder with the actual CEL expression
	errorString := result.String()
	errorString = strings.ReplaceAll(errorString, "EXPRESSION", c.expression)

	return errorString
}

func (c *celValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@CELValidation", "@", c.FieldName())
}

func (c *celValidator) Imports() []string {
	// Import validation helper package for CEL functions
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateCEL creates a new celValidator for fields with CEL marker.
// This validator supports all field types since CEL can handle various data types.
func ValidateCEL(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	celExpression, ok := expressions[markers.GoValidMarkerCEL]
	if !ok {
		return nil
	}

	// CEL expressions must not be empty
	if strings.TrimSpace(celExpression) == "" {
		return nil
	}

	return &celValidator{
		pass:       pass,
		field:      field,
		expression: celExpression,
	}
}