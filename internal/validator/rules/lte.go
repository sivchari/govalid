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

type lteValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	lteValue   string
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*lteValidator)(nil)

func (m *lteValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.Not(expr.Paren(expr.LTE(
			expr.Field("t", m.FieldName()),
			expr.IntLit(m.lteValue),
		))),
	}
}

func (m *lteValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *lteValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *lteValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "LTEValidation",
		Comment: fmt.Sprintf("is the error returned when the value of the field is greater than %s.", m.lteValue),
		Reason:  fmt.Sprintf("field %s must be less than or equal to %s", m.FieldName(), m.lteValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
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
