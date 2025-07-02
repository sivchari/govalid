// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
)

type lteValidator struct {
	pass     *codegen.Pass
	field    *ast.Field
	lteValue string
}

var _ validator.Validator = (*lteValidator)(nil)

const lteKey = "%s-lte"

func (m *lteValidator) Validate() string {
	return fmt.Sprintf("!(t.%s <= %s)", m.FieldName(), m.lteValue)
}

func (m *lteValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *lteValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(lteKey, m.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(lteKey, m.FieldName())] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@LTEValidation is the error returned when the value of the field is greater than %s.
	Err@LTEValidation = errors.New("field @ must be less than or equal to %s")`, "@", m.FieldName()), m.lteValue, m.lteValue)
}

func (m *lteValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@LTEValidation", "@", m.FieldName())
}

func (m *lteValidator) Imports() []string {
	return []string{}
}

// ValidateLTE creates a new lteValidator if the field type is numeric and the lte marker is present.
func ValidateLTE(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	lteValue, ok := expressions[markers.GoValidMarkerLTE]
	if !ok {
		return nil
	}

	return &lteValidator{
		pass:     pass,
		field:    field,
		lteValue: lteValue,
	}
}
