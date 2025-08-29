// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type numericValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*numericValidator)(nil)

const numericKey = "%s-numeric"

func (m *numericValidator) Validate() string {
	return fmt.Sprintf("!validationhelper.IsNumeric(t.%s)", m.FieldName())
}

func (m *numericValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *numericValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *numericValidator) Err() string {
	key := fmt.Sprintf(numericKey, m.structName+m.FieldPath().CleanedPath())
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
		// [@ERRVARIABLE] is the error returned when the field [@FIELD] is not numeric.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be numeric",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	legacyErrVarName := fmt.Sprintf("Err%s%sNumericValidation", m.structName, m.FieldName())
	currentErrVarName := m.ErrVariable()

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", currentErrVarName,
		"[@LEGACYERRVAR]", legacyErrVarName,
		"[@FIELD]", m.FieldName(),
		"[@PATH]", m.FieldPath().String(),
		"[@TYPE]", m.ruleName,
	)

	if currentErrVarName != legacyErrVarName {
		return replacer.Replace(deprecationNoticeTemplate + errTemplate)
	}

	return replacer.Replace(errTemplate)
}

func (m *numericValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]NumericValidation", "[@PATH]", m.FieldPath().CleanedPath())
}

func (m *numericValidator) Imports() []string {
	return []string{
		"github.com/sivchari/govalid/validation/validationhelper",
	}
}

// ValidateNumeric creates a new numericValidator if the 'numeric' marker is present and field is string.
func ValidateNumeric(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &numericValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
