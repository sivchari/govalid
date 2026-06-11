package rules

import (
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type alphanumValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*alphanumValidator)(nil)

func (v *alphanumValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr:    expr.Not(expr.Call("validationhelper", "IsValidAlphanum", expr.Field("t", v.FieldName()))),
		Imports: []string{"github.com/sivchari/govalid/validation/validationhelper"},
	}
}

func (v *alphanumValidator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *alphanumValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(v.structName, v.parentPath, v.FieldName())
}

func (v *alphanumValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + v.FieldPath().CleanedPath() + "AlphanumValidation",
		Comment: "is the error returned when field " + v.FieldName() + " is not alphanumeric.",
		Reason:  "field " + v.FieldName() + " must be alphanumeric",
		Path:    v.FieldPath().String(),
		Type:    v.ruleName,
	}
}

// ValidateAlphanum creates a new alphanumValidator for string types.
func ValidateAlphanum(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &alphanumValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
