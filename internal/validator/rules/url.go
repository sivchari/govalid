package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator"
)

type urlValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ validator.Validator = (*urlValidator)(nil)

const urlKey = "%s-url"

func (u *urlValidator) Validate() string {
	fieldName := u.FieldName()
	// Use external helper function for better maintainability
	return fmt.Sprintf("!validationhelper.IsValidURL(t.%s)", fieldName)
}

func (u *urlValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *urlValidator) Err() string {
	fieldName := u.FieldName()

	var result strings.Builder

	// No need to generate inline function - using external helper
	if validator.GeneratorMemory[fmt.Sprintf(urlKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(urlKey, fieldName)] = true

	result.WriteString(strings.ReplaceAll(`
	// Err@URLValidation is the error returned when the field is not a valid URL.
	Err@URLValidation = errors.New("field @ must be a valid URL")`, `@`, fieldName))

	return result.String()
}

func (u *urlValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@URLValidation", `@`, u.FieldName())
}

func (u *urlValidator) Imports() []string {
	// Import validation helper package
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateURL creates a new urlValidator for string types.
func ValidateURL(pass *codegen.Pass, field *ast.Field, _ map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &urlValidator{
		pass:  pass,
		field: field,
	}
}
