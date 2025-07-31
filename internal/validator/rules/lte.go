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
	pass       *codegen.Pass
	field      *ast.Field
	lteValue   string
	structName string
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
	key := fmt.Sprintf(lteKey, m.structName+m.FieldName())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// Err@LTEValidation is the error returned when the value of the field is greater than [VALUE].
		Err@LTEValidation = govaliderrors.ValidationError{Reason:"field @ must be less than or equal to [VALUE]",Path:"PATH"}
	`

	replacer := strings.NewReplacer(
		"@", m.structName+m.FieldName(),
		"PATH", fmt.Sprintf("%s.%s", m.structName, m.FieldName()),
		"[VALUE]", m.lteValue,
	)

	return replacer.Replace(errTemplate)
}

func (m *lteValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@LTEValidation", "@", m.structName+m.FieldName())
}

func (m *lteValidator) Imports() []string {
	return []string{}
}

// ValidateLTE creates a new lteValidator if the field type is numeric and the lte marker is present.
func ValidateLTE(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	lteValue, ok := expressions[markers.GoValidMarkerLte]
	if !ok {
		return nil
	}

	return &lteValidator{
		pass:       pass,
		field:      field,
		lteValue:   lteValue,
		structName: structName,
	}
}
