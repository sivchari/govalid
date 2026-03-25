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

type uuidValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*uuidValidator)(nil)

func (u *uuidValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr:    expr.Not(expr.Call("validationhelper", "IsValidUUID", expr.Field("t", u.FieldName()))),
		Imports: []string{"github.com/sivchari/govalid/validation/validationhelper"},
	}
}

func (u *uuidValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *uuidValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(u.structName, u.parentPath, u.FieldName())
}

func (u *uuidValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + u.FieldPath().CleanedPath() + "UUIDValidation",
		Comment: "is the error returned when the field is not a valid UUID.",
		Reason:  "field " + u.FieldName() + " must be a valid UUID",
		Path:    u.FieldPath().String(),
		Type:    u.ruleName,
	}
}

// ValidateUUID creates a new uuidValidator for string types.
func ValidateUUID(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &uuidValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
