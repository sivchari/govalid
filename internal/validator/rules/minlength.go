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

type minLengthValidator struct {
	pass           *codegen.Pass
	field          *ast.Field
	minLengthValue string
	structName     string
	ruleName       string
	parentPath     string
}

var _ validator.Validator = (*minLengthValidator)(nil)

func (m *minLengthValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.LT(
			expr.Call("utf8", "RuneCountInString", expr.Field("t", m.FieldName())),
			expr.IntLit(m.minLengthValue),
		),
		Imports: []string{"unicode/utf8"},
	}
}

func (m *minLengthValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *minLengthValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *minLengthValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "MinLengthValidation",
		Comment: fmt.Sprintf("is the error returned when the length of the field is less than the minimum of %s.", m.minLengthValue),
		Reason:  fmt.Sprintf("field %s must have a minimum length of %s", m.FieldName(), m.minLengthValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
}

// ValidateMinLength creates a new minLengthValidator if the field type is string and the minlength marker is present.
func ValidateMinLength(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	minLengthValue, ok := input.Expressions[markers.GoValidMarkerMinlength]
	if !ok {
		return nil
	}

	return &minLengthValidator{
		pass:           input.Pass,
		field:          input.Field,
		minLengthValue: minLengthValue,
		structName:     input.StructName,
		ruleName:       input.RuleName,
		parentPath:     input.ParentPath,
	}
}
