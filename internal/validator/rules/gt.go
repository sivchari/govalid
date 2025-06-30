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

type gtValidator struct {
	pass    *codegen.Pass
	field   *ast.Field
	gtValue string
}

var _ validator.Validator = (*gtValidator)(nil)

const gtKey = "%s-gt"

func (m *gtValidator) Validate() string {
	return fmt.Sprintf("!(t.%s > %s)", m.FieldName(), m.gtValue)
}

func (m *gtValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *gtValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(gtKey, m.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(gtKey, m.FieldName())] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@GTValidation is the error returned when the value of the field is less than the %s.
	Err@GTValidation = errors.New("field @ must be greater than %s")`, "@", m.FieldName()), m.gtValue, m.gtValue)
}

func (m *gtValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@GTValidation", "@", m.FieldName())
}

// ValidateGT creates a new gtValidator if the field type is numeric and the max marker is present.
func ValidateGT(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	gtValue, ok := expressions[markers.GoValidMarkerGT]
	if !ok {
		return nil
	}

	return &gtValidator{
		pass:    pass,
		field:   field,
		gtValue: gtValue,
	}
}
