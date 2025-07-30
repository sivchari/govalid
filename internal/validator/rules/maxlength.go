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
	structName     string
}

var _ validator.Validator = (*maxLengthValidator)(nil)

const maxLengthKey = "%s-maxlength"

func (m *maxLengthValidator) Validate() string {
	return fmt.Sprintf("utf8.RuneCountInString(t.%s) > %s", m.FieldName(), m.maxLengthValue)
}

func (m *maxLengthValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxLengthValidator) Err() string {
	key := fmt.Sprintf(maxLengthKey, m.structName+m.FieldName())

	var result strings.Builder

	if validator.GeneratorMemory[key] {
		return result.String()
	}

	validator.GeneratorMemory[key] = true

	result.WriteString(
		strings.ReplaceAll(`
			// Err@MaxLengthValidation is the error returned when the length of the field exceeds the maximum of %s.
			Err@MaxLengthValidation = govaliderrors.ValidationError{Reason:"field @ must have a maximum length of %s",Path:"PATH"}
			`,
			"@",
			m.structName+m.FieldName(),
		),
	)

	return strings.ReplaceAll(result.String(), "PATH", fmt.Sprintf("%s.%s", m.structName, m.FieldName()))
}

func (m *maxLengthValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@MaxLengthValidation", "@", m.structName+m.FieldName())
}

func (m *maxLengthValidator) Imports() []string {
	return []string{"unicode/utf8"}
}

// ValidateMaxLength creates a new maxLengthValidator if the field type is string and the maxlength marker is present.
func ValidateMaxLength(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	maxLengthValue, ok := expressions[markers.GoValidMarkerMaxlength]
	if !ok {
		return nil
	}

	return &maxLengthValidator{
		pass:           pass,
		field:          field,
		maxLengthValue: maxLengthValue,
		structName:     structName,
	}
}
