package rules

import (
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type urlValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*urlValidator)(nil)

func (u *urlValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr:    expr.Not(expr.Call("validationhelper", "IsValidURL", expr.Field("t", u.FieldName()))),
		Imports: []string{"github.com/sivchari/govalid/validation/validationhelper"},
	}
}

func (u *urlValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *urlValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(u.structName, u.parentPath, u.FieldName())
}

func (u *urlValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + u.FieldPath().CleanedPath() + "URLValidation",
		Comment: "is the error returned when the field is not a valid URL.",
		Reason:  "field " + u.FieldName() + " must be a valid URL",
		Path:    u.FieldPath().String(),
		Type:    u.ruleName,
	}
}

// ValidateURL creates a new urlValidator for string types.
func ValidateURL(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &urlValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
