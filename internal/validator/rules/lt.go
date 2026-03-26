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

type ltValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	ltValue    string
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*ltValidator)(nil)

func (m *ltValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.Not(expr.Paren(expr.LT(
			expr.Field("t", m.FieldName()),
			expr.IntLit(m.ltValue),
		))),
	}
}

func (m *ltValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *ltValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *ltValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "LTValidation",
		Comment: fmt.Sprintf("is the error returned when the value of the field is greater than the %s.", m.ltValue),
		Reason:  fmt.Sprintf("field %s must be less than %s", m.FieldName(), m.ltValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
}

// ValidateLT creates a new ltValidator if the field type is numeric and the min marker is present.
func ValidateLT(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	ltValue, ok := input.Expressions[markers.GoValidMarkerLt]
	if !ok {
		return nil
	}

	return &ltValidator{
		pass:       input.Pass,
		field:      input.Field,
		ltValue:    ltValue,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
