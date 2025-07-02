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

type maxItemsValidator struct {
	pass          *codegen.Pass
	field         *ast.Field
	maxItemsValue string
}

var _ validator.Validator = (*maxItemsValidator)(nil)

const maxItemsKey = "%s-maxitems"

func (m *maxItemsValidator) Validate() string {
	return fmt.Sprintf("len(t.%s) > %s", m.FieldName(), m.maxItemsValue)
}

func (m *maxItemsValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxItemsValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(maxItemsKey, m.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(maxItemsKey, m.FieldName())] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@MaxItemsValidation is the error returned when the length of the field exceeds the maximum of %s.
	Err@MaxItemsValidation = errors.New("field @ must have a maximum of %s items")`, "@", m.FieldName()), m.maxItemsValue, m.maxItemsValue)
}

func (m *maxItemsValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@MaxItemsValidation", "@", m.FieldName())
}

func (m *maxItemsValidator) Imports() []string {
	return []string{}
}

// ValidateMaxItems creates a new maxItemsValidator if the field type supports len() and the maxitems marker is present.
func ValidateMaxItems(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a type that supports len() (exclude strings - use maxlength instead)
	switch typ.Underlying().(type) {
	case *types.Slice, *types.Array, *types.Map, *types.Chan:
		// Valid types for maxitems
	default:
		return nil
	}

	maxItemsValue, ok := expressions[markers.GoValidMarkerMaxItems]
	if !ok {
		return nil
	}

	return &maxItemsValidator{
		pass:          pass,
		field:         field,
		maxItemsValue: maxItemsValue,
	}
}
