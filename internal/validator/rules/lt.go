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

type ltValidator struct {
	pass    *codegen.Pass
	field   *ast.Field
	ltValue string
}

var _ validator.Validator = (*ltValidator)(nil)

const ltKey = "%s-lt"

func (m *ltValidator) Validate() string {
	return fmt.Sprintf("!(t.%s < %s)", m.FieldName(), m.ltValue)
}

func (m *ltValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *ltValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(ltKey, m.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(ltKey, m.FieldName())] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@LTValidation is the error returned when the value of the field is greater than the %s.
	Err@LTValidation = errors.New("field @ must be less than %s")`, "@", m.FieldName()), m.ltValue, m.ltValue)
}

func (m *ltValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@LTValidation", "@", m.FieldName())
}

// ValidateLT creates a new ltValidator if the field type is numeric and the min marker is present.
func ValidateLT(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	ltValue, ok := expressions[markers.GoValidMarkerLT]
	if !ok {
		return nil
	}

	return &ltValidator{
		pass:    pass,
		field:   field,
		ltValue: ltValue,
	}
}
