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

type gteValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	gteValue   string
	structName string
}

var _ validator.Validator = (*gteValidator)(nil)

const gteKey = "%s-gte"

func (m *gteValidator) Validate() string {
	return fmt.Sprintf("!(t.%s >= %s)", m.FieldName(), m.gteValue)
}

func (m *gteValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *gteValidator) Err() string {
	key := fmt.Sprintf(gteKey, m.structName+m.FieldName())
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@GTEValidation is the error returned when the value of the field is less than %s.
	Err@GTEValidation = errors.New("field @ must be greater than or equal to %s")`, "@", m.structName+m.FieldName()), m.gteValue, m.gteValue)
}

func (m *gteValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@GTEValidation", "@", m.structName+m.FieldName())
}

func (m *gteValidator) Imports() []string {
	return []string{}
}

// ValidateGTE creates a new gteValidator if the field type is numeric and the gte marker is present.
func ValidateGTE(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	gteValue, ok := expressions[markers.GoValidMarkerGTE]
	if !ok {
		return nil
	}

	return &gteValidator{
		pass:       pass,
		field:      field,
		gteValue:   gteValue,
		structName: structName,
	}
}
