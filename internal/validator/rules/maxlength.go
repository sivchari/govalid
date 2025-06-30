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

type maxLengthValidator struct {
	pass           *codegen.Pass
	field          *ast.Field
	maxLengthValue string
}

var _ validator.Validator = (*maxLengthValidator)(nil)

const maxLengthKey = "%s-maxlength"

func (m *maxLengthValidator) Validate() string {
	return fmt.Sprintf("len([]rune(t.%s)) > %s", m.FieldName(), m.maxLengthValue)
}

func (m *maxLengthValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxLengthValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(maxLengthKey, m.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(maxLengthKey, m.FieldName())] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@MaxLengthValidation is the error returned when the length of the field exceeds the maximum of %s.
	Err@MaxLengthValidation = errors.New("field @ must have a maximum length of %s")`, "@", m.FieldName()), m.maxLengthValue, m.maxLengthValue)
}

func (m *maxLengthValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@MaxLengthValidation", "@", m.FieldName())
}

// ValidateMaxLength creates a new maxLengthValidator if the field type is string and the maxlength marker is present.
func ValidateMaxLength(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	maxLengthValue, ok := expressions[markers.GoValidMarkerMaxLength]
	if !ok {
		return nil
	}

	return &maxLengthValidator{
		pass:           pass,
		field:          field,
		maxLengthValue: maxLengthValue,
	}
}