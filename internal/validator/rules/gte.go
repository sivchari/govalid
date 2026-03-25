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

type gteValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	gteValue   string
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*gteValidator)(nil)

func (m *gteValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.Not(expr.Paren(expr.GTE(
			expr.Field("t", m.FieldName()),
			expr.IntLit(m.gteValue),
		))),
	}
}

func (m *gteValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *gteValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *gteValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "GTEValidation",
		Comment: fmt.Sprintf("is the error returned when the value of the field is less than %s.", m.gteValue),
		Reason:  fmt.Sprintf("field %s must be greater than or equal to %s", m.FieldName(), m.gteValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
}

// ValidateGTE creates a new gteValidator if the field type is numeric and the gte marker is present.
func ValidateGTE(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	gteValue, ok := input.Expressions[markers.GoValidMarkerGte]
	if !ok {
		return nil
	}

	return &gteValidator{
		pass:       input.Pass,
		field:      input.Field,
		gteValue:   gteValue,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
