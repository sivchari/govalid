// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
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

func (m *maxItemsValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.GT(
			expr.Call("", "len", expr.Field("t", m.FieldName())),
			expr.IntLit(m.maxItemsValue),
		),
	}
}

func (m *maxItemsValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxItemsValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *maxItemsValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "MaxItemsValidation",
		Comment: fmt.Sprintf("is the error returned when the length of the field exceeds the maximum of %s.", m.maxItemsValue),
		Reason:  fmt.Sprintf("field %s must have a maximum of %s items", m.FieldName(), m.maxItemsValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
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
