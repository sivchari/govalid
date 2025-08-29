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
	"github.com/sivchari/govalid/internal/validator/registry"
)

type maxLengthValidator struct {
	pass           *codegen.Pass
	field          *ast.Field
	maxLengthValue string
	structName     string
	ruleName       string
	parentPath     string
}

var _ validator.Validator = (*maxLengthValidator)(nil)

const maxLengthKey = "%s-maxlength"

func (m *maxLengthValidator) Validate() string {
	return fmt.Sprintf("utf8.RuneCountInString(t.%s) > %s", m.FieldName(), m.maxLengthValue)
}

func (m *maxLengthValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxLengthValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *maxLengthValidator) Err() string {
	key := fmt.Sprintf(maxLengthKey, m.structName+m.FieldPath().CleanedPath())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const deprecationNoticeTemplate = `
		// Deprecated: Use [@ERRVARIABLE]
		//
		// [@LEGACYERRVAR] is deprecated and is kept for compatibility purpose.
		[@LEGACYERRVAR] = [@ERRVARIABLE]
	`

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the length of the field exceeds the maximum of [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must have a maximum length of [@VALUE]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	legacyErrVarName := fmt.Sprintf("Err%s%sMaxLengthValidation", m.structName, m.FieldName())
	currentErrVarName := m.ErrVariable()

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", currentErrVarName,
		"[@LEGACYERRVAR]", legacyErrVarName,
		"[@FIELD]", m.FieldName(),
		"[@PATH]", m.FieldPath().String(),
		"[@VALUE]", m.maxLengthValue,
		"[@TYPE]", m.ruleName,
	)

	if currentErrVarName != legacyErrVarName {
		return replacer.Replace(deprecationNoticeTemplate + errTemplate)
	}

	return replacer.Replace(errTemplate)
}

func (m *maxLengthValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]MaxLengthValidation", "[@PATH]", m.FieldPath().CleanedPath())
}

func (m *maxLengthValidator) Imports() []string {
	return []string{"unicode/utf8"}
}

// ValidateMaxLength creates a new maxLengthValidator if the field type is string and the maxlength marker is present.
func ValidateMaxLength(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	maxLengthValue, ok := input.Expressions[markers.GoValidMarkerMaxlength]
	if !ok {
		return nil
	}

	return &maxLengthValidator{
		pass:           input.Pass,
		field:          input.Field,
		maxLengthValue: maxLengthValue,
		structName:     input.StructName,
		ruleName:       input.RuleName,
		parentPath:     input.ParentPath,
	}
}
