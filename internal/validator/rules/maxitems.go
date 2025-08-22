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

type maxItemsValidator struct {
	pass          *codegen.Pass
	field         *ast.Field
	maxItemsValue string
	structName    string
	ruleName      string
	parentPath    string
}

var _ validator.Validator = (*maxItemsValidator)(nil)

const maxItemsKey = "%s-maxitems"

func (m *maxItemsValidator) Validate() string {
	return fmt.Sprintf("len(t.%s) > %s", m.FieldName(), m.maxItemsValue)
}

func (m *maxItemsValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxItemsValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *maxItemsValidator) Err() string {
	key := fmt.Sprintf(maxItemsKey, m.structName+m.FieldPath().CleanedPath())

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
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must have a maximum of [@VALUE] items",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	legacyErrVarName := fmt.Sprintf("Err%s%sMaxItemsValidation", m.structName, m.FieldName())
	currentErrVarName := m.ErrVariable()

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", currentErrVarName,
		"[@LEGACYERRVAR]", legacyErrVarName,
		"[@FIELD]", m.FieldName(),
		"[@PATH]", m.FieldPath().String(),
		"[@VALUE]", m.maxItemsValue,
		"[@TYPE]", m.ruleName,
	)

	if currentErrVarName != legacyErrVarName {
		return replacer.Replace(deprecationNoticeTemplate + errTemplate)
	}

	return replacer.Replace(errTemplate)
}

func (m *maxItemsValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]MaxItemsValidation", "[@PATH]", m.FieldPath().CleanedPath())
}

func (m *maxItemsValidator) Imports() []string {
	return []string{}
}

// ValidateMaxItems creates a new maxItemsValidator if the field type supports len() and the maxitems marker is present.
func ValidateMaxItems(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a type that supports len() (exclude strings - use maxlength instead)
	switch typ.Underlying().(type) {
	case *types.Slice, *types.Array, *types.Map, *types.Chan:
		// Valid types for maxitems
	default:
		return nil
	}

	maxItemsValue, ok := input.Expressions[markers.GoValidMarkerMaxitems]
	if !ok {
		return nil
	}

	return &maxItemsValidator{
		pass:          input.Pass,
		field:         input.Field,
		maxItemsValue: maxItemsValue,
		structName:    input.StructName,
		ruleName:      input.RuleName,
		parentPath:    input.ParentPath,
	}
}
