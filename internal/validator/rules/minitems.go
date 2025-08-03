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

type minItemsValidator struct {
	pass          *codegen.Pass
	field         *ast.Field
	minItemsValue string
	structName    string
	ruleName      string
}

var _ validator.Validator = (*minItemsValidator)(nil)

const minItemsKey = "%s-minitems"

func (m *minItemsValidator) Validate() string {
	return fmt.Sprintf("len(t.%s) < %s", m.FieldName(), m.minItemsValue)
}

func (m *minItemsValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *minItemsValidator) Err() string {
	key := fmt.Sprintf(minItemsKey, m.structName+m.FieldName())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the length of the field is less than the minimum of [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must have a minimum of [@VALUE] items",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", m.ErrVariable(),
		"[@FIELD]", m.FieldName(),
		"[@PATH]", fmt.Sprintf("%s.%s", m.structName, m.FieldName()),
		"[@VALUE]", m.minItemsValue,
		"[@TYPE]", m.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (m *minItemsValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]MinItemsValidation", "[@PATH]", m.structName+m.FieldName())
}

func (m *minItemsValidator) Imports() []string {
	return []string{}
}

// ValidateMinItems creates a new minItemsValidator if the field type supports len() and the minitems marker is present.
func ValidateMinItems(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName, ruleName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a type that supports len() (exclude strings - use minlength instead)
	switch typ.Underlying().(type) {
	case *types.Slice, *types.Array, *types.Map, *types.Chan:
		// Valid types for minitems
	default:
		return nil
	}

	minItemsValue, ok := expressions[markers.GoValidMarkerMinitems]
	if !ok {
		return nil
	}

	return &minItemsValidator{
		pass:          pass,
		field:         field,
		minItemsValue: minItemsValue,
		structName:    structName,
		ruleName:      ruleName,
	}
}
