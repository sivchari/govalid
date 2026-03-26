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

type maxLengthValidator struct {
	pass           *codegen.Pass
	field          *ast.Field
	maxLengthValue string
	structName     string
	ruleName       string
	parentPath     string
}

var _ validator.Validator = (*maxLengthValidator)(nil)

func (m *maxLengthValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.GT(
			expr.Call("utf8", "RuneCountInString", expr.Field("t", m.FieldName())),
			expr.IntLit(m.maxLengthValue),
		),
		Imports: []string{"unicode/utf8"},
	}
}

func (m *maxLengthValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxLengthValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *maxLengthValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + m.FieldPath().CleanedPath() + "MaxLengthValidation",
		Comment: fmt.Sprintf("is the error returned when the length of the field exceeds the maximum of %s.", m.maxLengthValue),
		Reason:  fmt.Sprintf("field %s must have a maximum length of %s", m.FieldName(), m.maxLengthValue),
		Path:    m.FieldPath().String(),
		Type:    m.ruleName,
	}
}

// ValidateMaxLength creates a new maxLengthValidator if the field type is string and the maxlength marker is present.
func ValidateMaxLength(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	maxLengthValue, ok := input.Expressions[markers.GoValidMarkerMaxlength]
	if !ok {
		return nil
	}

	return &maxLengthValidator{
		pass:           input.Pass,
		field:          input.Field,
		maxLengthValue: maxLengthValue,
		structName:     input.StructName,
		ruleName:       input.RuleName,
		parentPath:     input.ParentPath,
	}
}
