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

type minItemsValidator struct {
	pass          *codegen.Pass
	field         *ast.Field
	minItemsValue string
	structName    string
	ruleName      string
	parentPath    string
}

var _ validator.Validator = (*minItemsValidator)(nil)

func (m *minItemsValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.LT(
			expr.Call("", "len", expr.Field("t", m.FieldName())),
			expr.IntLit(m.minItemsValue),
		),
	}
}

func (m *minItemsValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *minItemsValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *minItemsValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "MinItemsValidation",
		Comment: fmt.Sprintf("is the error returned when the length of the field is less than the minimum of %s.", m.minItemsValue),
		Reason:  fmt.Sprintf("field %s must have a minimum of %s items", m.FieldName(), m.minItemsValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
}

// ValidateMinItems creates a new minItemsValidator if the field type supports len() and the minitems marker is present.
func ValidateMinItems(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a type that supports len() (exclude strings - use minlength instead)
	switch typ.Underlying().(type) {
	case *types.Slice, *types.Array, *types.Map, *types.Chan:
		// Valid types for minitems
	default:
		return nil
	}

	minItemsValue, ok := input.Expressions[markers.GoValidMarkerMinitems]
	if !ok {
		return nil
	}

	return &minItemsValidator{
		pass:          input.Pass,
		field:         input.Field,
		minItemsValue: minItemsValue,
		structName:    input.StructName,
		ruleName:      input.RuleName,
		parentPath:    input.ParentPath,
	}
}
