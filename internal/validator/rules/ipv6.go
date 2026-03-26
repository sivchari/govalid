// Package rules implements validation rules for fields in structs.
package rules

import (
	"go/ast"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type ipv6Validator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*ipv6Validator)(nil)

func (v *ipv6Validator) Condition() *validator.Condition {
	return &validator.Condition{
		IfInitStmt: expr.ShortVarDecl("ip",
			expr.Call("net", "ParseIP", expr.Field("t", v.FieldName())),
		),
		Expr: expr.Or(
			expr.Eq(expr.Ident("ip"), expr.Nil()),
			expr.NEq(expr.MethodCall(expr.Ident("ip"), "To4"), expr.Nil()),
		),
		Imports: []string{"net"},
	}
}

func (v *ipv6Validator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *ipv6Validator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(v.structName, v.parentPath, v.FieldName())
}

func (v *ipv6Validator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + v.FieldPath().CleanedPath() + "Ipv6Validation",
		Comment: "is returned when the " + v.FieldName() + " fails ipv6 validation.",
		Reason:  "field " + v.FieldName() + " failed ipv6 validation",
		Path:    v.FieldPath().String(),
		Type:    v.ruleName,
	}
}

// ValidateIpv6 creates a new ipv6 validator for the given field.
func ValidateIpv6(input registry.ValidatorInput) validator.Validator {
	return &ipv6Validator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
