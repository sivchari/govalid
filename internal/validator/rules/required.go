// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/validatorhelper"
)

type requiredValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*requiredValidator)(nil)

func (r *requiredValidator) Condition() *validator.Condition {
	typ := r.pass.TypesInfo.TypeOf(r.field.Type)
	field := expr.Field("t", r.FieldName())

	switch typ.(type) {
	case *types.Slice, *types.Map, *types.Chan:
		return &validator.Condition{Expr: expr.Eq(field, expr.Nil())}
	case *types.Array:
		return &validator.Condition{Expr: expr.Eq(
			expr.Call("", "len", field),
			expr.IntLit("0"),
		)}
	}

	zero := validatorhelper.Zero(typ)
	if zero == "" {
		return nil
	}

	zeroExpr, _ := expr.Parse(zero)

	return &validator.Condition{Expr: expr.Eq(field, zeroExpr)}
}

func (r *requiredValidator) FieldName() string {
	return r.field.Names[0].Name
}

func (r *requiredValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(r.structName, r.parentPath, r.FieldName())
}

func (r *requiredValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + r.FieldPath().CleanedPath() + "RequiredValidation",
		Comment: fmt.Sprintf("is returned when the %s is required but not provided.", r.FieldName()),
		Reason:  fmt.Sprintf("field %s is required", r.FieldName()),
		Path:    r.FieldPath().String(),
		Type:    r.ruleName,
	}
}

// ValidateRequired creates a new required validator for the given field.
func ValidateRequired(input registry.ValidatorInput) validator.Validator {
	return &requiredValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
