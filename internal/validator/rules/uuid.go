package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator"
)

type uuidValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ validator.Validator = (*uuidValidator)(nil)

const uuidKey = "%s-uuid"

func (u *uuidValidator) Validate() string {
	fieldName := u.FieldName()
	// Use external helper function for better maintainability
	return fmt.Sprintf("!validationhelper.IsValidUUID(t.%s)", fieldName)
}

func (u *uuidValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *uuidValidator) Err() string {
	fieldName := u.FieldName()

	var result strings.Builder

	// No need to generate inline function - using external helper
	if validator.GeneratorMemory[fmt.Sprintf(uuidKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(uuidKey, fieldName)] = true

	result.WriteString(strings.ReplaceAll(`
	// Err@UUIDValidation is the error returned when the field is not a valid UUID.
	Err@UUIDValidation = errors.New("field @ must be a valid UUID")`, `@`, fieldName))

	return result.String()
}

func (u *uuidValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@UUIDValidation", `@`, u.FieldName())
}

func (u *uuidValidator) Imports() []string {
	// Import validation helper package
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateUUID creates a new uuidValidator for string types.
func ValidateUUID(pass *codegen.Pass, field *ast.Field, _ map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &uuidValidator{
		pass:  pass,
		field: field,
	}
}
