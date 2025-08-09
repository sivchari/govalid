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
	ruleName       string
	parentPath     string
}

var _ validator.Validator = (*minLengthValidator)(nil)

const minLengthKey = "%s-minlength"

func (m *minLengthValidator) Validate() string {
	return fmt.Sprintf("utf8.RuneCountInString(t.%s) < %s", m.FieldName(), m.minLengthValue)
}

func (m *minLengthValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *minLengthValidator) FieldPath() validator.FieldPath {
	if m.parentPath == "" {
		return validator.NewFieldPath(m.structName, m.FieldName())
	}
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *minLengthValidator) Err() string {
	key := fmt.Sprintf(minLengthKey, m.structName+m.FieldPath().WithoutDots())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the length of the field is less than the minimum of [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must have a minimum length of [@VALUE]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", m.ErrVariable(),
		"[@FIELD]", m.FieldName(),
		"[@PATH]", m.FieldPath().String(),
		"[@VALUE]", m.minLengthValue,
		"[@TYPE]", m.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (m *minLengthValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]MinLengthValidation", "[@PATH]", m.FieldPath().WithoutDots())
}

func (m *minLengthValidator) Imports() []string {
	return []string{"unicode/utf8"}
}

// ValidateMinLength creates a new minLengthValidator if the field type is string and the minlength marker is present.
func ValidateMinLength(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName, ruleName string, parentPath string) validator.Validator {
	fieldPath := validator.NewFieldPath(structName, field.Names[0].Name)
	if parentPath != "" {
		fieldPath = validator.NewFieldPath(structName, parentPath, field.Names[0].Name)
	}
	validator.GeneratorMemory[fmt.Sprintf(minLengthKey, structName+fieldPath.WithoutDots())] = false
	
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	minLengthValue, ok := expressions[markers.GoValidMarkerMinlength]
	if !ok {
		return nil
	}

	return &minLengthValidator{
		pass:           pass,
		field:          field,
		minLengthValue: minLengthValue,
		structName:     structName,
		ruleName:       ruleName,
		parentPath:     parentPath,
	}
}
