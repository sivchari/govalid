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

type lengthValidator struct {
	pass        *codegen.Pass
	field       *ast.Field
	lengthValue string
	structName  string
	ruleName    string
	parentPath  string
}

var _ validator.Validator = (*lengthValidator)(nil)

func (l *lengthValidator) Condition() *validator.Condition {
	return &validator.Condition{
		Expr: expr.NEq(
			expr.Call("utf8", "RuneCountInString", expr.Field("t", l.FieldName())),
			expr.IntLit(l.lengthValue),
		),
		Imports: []string{"unicode/utf8"},
	}
}

func (l *lengthValidator) FieldName() string {
	return l.field.Names[0].Name
}

func (l *lengthValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(l.structName, l.parentPath, l.FieldName())
}

func (l *lengthValidator) ErrDecl() validator.ErrDecl {
	return validator.ErrDecl{
		VarName: "Err" + l.FieldPath().CleanedPath() + "LengthValidation",
		Comment: fmt.Sprintf("is the error returned when the length of the field is not exactly %s.", l.lengthValue),
		Reason:  fmt.Sprintf("field %s length must be exactly %s", l.FieldName(), l.lengthValue),
		Path:    l.FieldPath().String(),
		Type:    l.ruleName,
	}
}

// ValidateLength creates a new lengthValidator if the field type is string and the length marker is present.
func ValidateLength(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	lengthValue, ok := input.Expressions[markers.GoValidMarkerLength]
	if !ok {
		return nil
	}

	return &lengthValidator{
		pass:        input.Pass,
		field:       input.Field,
		lengthValue: lengthValue,
		structName:  input.StructName,
		ruleName:    input.RuleName,
		parentPath:  input.ParentPath,
	}
}
