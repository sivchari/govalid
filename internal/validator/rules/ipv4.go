// Package rules implements validation rules for fields in structs.
package rules

import (
	"go/ast"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type ipv4Validator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*ipv4Validator)(nil)

func (v *ipv4Validator) Condition() *validator.Condition {
	return &validator.Condition{
		IfInitStmt: expr.ShortVarDecl("ip",
			expr.Call("net", "ParseIP", expr.Field("t", v.FieldName())),
		),
		Expr: expr.Or(
			expr.Eq(expr.Ident("ip"), expr.Nil()),
			expr.Eq(expr.MethodCall(expr.Ident("ip"), "To4"), expr.Nil()),
		),
		Imports: []string{"net"},
	}
}

func (v *ipv4Validator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *ipv4Validator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(v.structName, v.parentPath, v.FieldName())
}

func (v *ipv4Validator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + v.FieldPath().CleanedPath() + "Ipv4Validation",
		Comment: "is returned when the " + v.FieldName() + " fails ipv4 validation.",
		Reason:  "field " + v.FieldName() + " failed ipv4 validation",
		Path:    v.FieldPath().String(),
		Type:    v.ruleName,
	}
}

// ValidateIpv4 creates a new ipv4 validator for the given field.
func ValidateIpv4(input registry.ValidatorInput) validator.Validator {
	return &ipv4Validator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
