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
	pass       *codegen.Pass
	field      *ast.Field
	gtValue    string
	structName string
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
	key := fmt.Sprintf(gtKey, m.structName+m.FieldName())

	var result strings.Builder

	if validator.GeneratorMemory[key] {
		return result.String()
	}

	validator.GeneratorMemory[key] = true

	result.WriteString(
		strings.ReplaceAll(`
			// Err@GTValidation is the error returned when the value of the field is less than the %s.
			Err@GTValidation = govaliderrors.ValidationError{Reason:"field @ must be greater than %s",Path:"PATH"}
			`,
			"@",
			m.structName+m.FieldName(),
		),
	)

	return strings.ReplaceAll(result.String(), "PATH", fmt.Sprintf("%s.%s", m.structName, m.FieldName()))
}

func (m *gtValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@GTValidation", "@", m.structName+m.FieldName())
}

func (m *gtValidator) Imports() []string {
	return []string{}
}

// ValidateGT creates a new gtValidator if the field type is numeric and the max marker is present.
func ValidateGT(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
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
		pass:       pass,
		field:      field,
		gtValue:    gtValue,
		structName: structName,
	}
}
