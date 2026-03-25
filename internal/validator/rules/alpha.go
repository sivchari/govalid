package rules

import (
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type alphaValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*alphaValidator)(nil)

func (v *alphaValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr:    expr.Not(expr.Call("validationhelper", "IsValidAlpha", expr.Field("t", v.FieldName()))),
		Imports: []string{"github.com/sivchari/govalid/validation/validationhelper"},
	}
}

func (v *alphaValidator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *alphaValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(v.structName, v.parentPath, v.FieldName())
}

func (v *alphaValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + v.FieldPath().CleanedPath() + "AlphaValidation",
		Comment: "is the error returned when field " + v.FieldName() + " is not alphabetic.",
		Reason:  "field " + v.FieldName() + " must be alphabetic",
		Path:    v.FieldPath().String(),
		Type:    v.ruleName,
	}
}

// ValidateAlpha creates a new alphaValidator for string types.
func ValidateAlpha(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &alphaValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
