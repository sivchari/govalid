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

type minLengthValidator struct {
	pass           *codegen.Pass
	field          *ast.Field
	minLengthValue string
	structName     string
}

var _ validator.Validator = (*minLengthValidator)(nil)

const minLengthKey = "%s-minlength"

func (m *minLengthValidator) Validate() string {
	return fmt.Sprintf("utf8.RuneCountInString(t.%s) < %s", m.FieldName(), m.minLengthValue)
}

func (m *minLengthValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *minLengthValidator) Err() string {
	key := fmt.Sprintf(minLengthKey, m.structName+m.FieldName())
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@MinLengthValidation is the error returned when the length of the field is less than the minimum of %s.
	Err@MinLengthValidation = govaliderrors.ValidationError{Reason:"field @ must have a minimum length of %s"}`, "@", m.structName+m.FieldName()), m.minLengthValue, m.minLengthValue)
}

func (m *minLengthValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@MinLengthValidation", "@", m.structName+m.FieldName())
}

func (m *minLengthValidator) Imports() []string {
	return []string{"unicode/utf8"}
}

// ValidateMinLength creates a new minLengthValidator if the field type is string and the minlength marker is present.
func ValidateMinLength(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	minLengthValue, ok := expressions[markers.GoValidMarkerMinLength]
	if !ok {
		return nil
	}

	return &minLengthValidator{
		pass:           pass,
		field:          field,
		minLengthValue: minLengthValue,
		structName:     structName,
	}
}
