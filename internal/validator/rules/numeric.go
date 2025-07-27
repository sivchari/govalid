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

type numericValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
}

var _ validator.Validator = (*numericValidator)(nil)

const numericKey = "%s-numeric"

func (m *numericValidator) Validate() string {
	return fmt.Sprintf("!validationhelper.IsNumeric(t.%s)", m.FieldName())
}

func (m *numericValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *numericValidator) Err() string {
	key := fmt.Sprintf(numericKey, m.structName+m.FieldName())
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	return strings.ReplaceAll(`
	// Err@NumericValidation is the error returned when the field is not numeric.
	Err@NumericValidation = errors.New("field @ must be numeric")`, "@", m.structName+m.FieldName())
}

func (m *numericValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@NumericValidation", "@", m.structName+m.FieldName())
}

func (m *numericValidator) Imports() []string {
	return []string{
		"github.com/sivchari/govalid/validation/validationhelper",
	}
}

// ValidateNumeric creates a new numericValidator if the 'numeric' marker is present and field is string.
func ValidateNumeric(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &numericValidator{
		pass:       pass,
		field:      field,
		structName: structName,
	}
}
