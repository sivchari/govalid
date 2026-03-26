// Package rules implements validation rules for fields in structs.
package rules

import (
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type emailValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*emailValidator)(nil)

func (e *emailValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr:    expr.Not(expr.Call("validationhelper", "IsValidEmail", expr.Field("t", e.FieldName()))),
		Imports: []string{"github.com/sivchari/govalid/validation/validationhelper"},
	}
}

func (e *emailValidator) FieldName() string {
	return e.field.Names[0].Name
}

func (e *emailValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(e.structName, e.parentPath, e.FieldName())
}

func (e *emailValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + e.FieldPath().CleanedPath() + "EmailValidation",
		Comment: "is the error returned when the field is not a valid email address.",
		Reason:  "field " + e.FieldName() + " must be a valid email address",
		Path:    e.FieldPath().String(),
		Type:    e.ruleName,
	}
}

// ValidateEmail creates a new emailValidator for string types.
func ValidateEmail(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &emailValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
