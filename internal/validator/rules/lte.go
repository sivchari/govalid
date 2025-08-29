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

type lteValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	lteValue   string
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*lteValidator)(nil)

const lteKey = "%s-lte"

func (m *lteValidator) Validate() string {
	return fmt.Sprintf("!(t.%s <= %s)", m.FieldName(), m.lteValue)
}

func (m *lteValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *lteValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *lteValidator) Err() string {
	key := fmt.Sprintf(lteKey, m.structName+m.FieldPath().CleanedPath())

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
		// [@ERRVARIABLE] is the error returned when the value of the field is greater than [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be less than or equal to [@VALUE]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	legacyErrVarName := fmt.Sprintf("Err%s%sLTEValidation", m.structName, m.FieldName())
	currentErrVarName := m.ErrVariable()

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", currentErrVarName,
		"[@LEGACYERRVAR]", legacyErrVarName,
		"[@FIELD]", m.FieldName(),
		"[@PATH]", m.FieldPath().String(),
		"[@VALUE]", m.lteValue,
		"[@TYPE]", m.ruleName,
	)

	if currentErrVarName != legacyErrVarName {
		return replacer.Replace(deprecationNoticeTemplate + errTemplate)
	}

	return replacer.Replace(errTemplate)
}

func (m *lteValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]LTEValidation", "[@PATH]", m.FieldPath().CleanedPath())
}

func (m *lteValidator) Imports() []string {
	return []string{}
}

// ValidateLTE creates a new lteValidator if the field type is numeric and the lte marker is present.
func ValidateLTE(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	lteValue, ok := input.Expressions[markers.GoValidMarkerLte]
	if !ok {
		return nil
	}

	return &lteValidator{
		pass:       input.Pass,
		field:      input.Field,
		lteValue:   lteValue,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
