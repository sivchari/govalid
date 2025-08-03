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
	ruleName   string
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

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the value of the field is less than the [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be greater than [@VALUE]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", m.ErrVariable(),
		"[@FIELD]", m.FieldName(),
		"[@PATH]", fmt.Sprintf("%s.%s", m.structName, m.FieldName()),
		"[@VALUE]", m.gtValue,
		"[@TYPE]", m.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (m *gtValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]GTValidation", "[@PATH]", m.structName+m.FieldName())
}

func (m *gtValidator) Imports() []string {
	return []string{}
}

// ValidateGT creates a new gtValidator if the field type is numeric and the max marker is present.
func ValidateGT(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string, ruleName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	gtValue, ok := expressions[markers.GoValidMarkerGt]
	if !ok {
		return nil
	}

	return &gtValidator{
		pass:       pass,
		field:      field,
		gtValue:    gtValue,
		structName: structName,
		ruleName:   ruleName,
	}
}
