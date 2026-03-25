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

type gtValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	gtValue    string
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*gtValidator)(nil)

func (m *gtValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.Not(expr.Paren(expr.GT(
			expr.Field("t", m.FieldName()),
			expr.IntLit(m.gtValue),
		))),
	}
}

func (m *gtValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *gtValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *gtValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "GTValidation",
		Comment: fmt.Sprintf("is the error returned when the value of the field is less than the %s.", m.gtValue),
		Reason:  fmt.Sprintf("field %s must be greater than %s", m.FieldName(), m.gtValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
}

// ValidateGT creates a new gtValidator if the field type is numeric and the max marker is present.
func ValidateGT(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	gtValue, ok := input.Expressions[markers.GoValidMarkerGt]
	if !ok {
		return nil
	}

	return &gtValidator{
		pass:       input.Pass,
		field:      input.Field,
		gtValue:    gtValue,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
