// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type enumValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	enumValues []string
	isString   bool
	isNumeric  bool
	isCustom   bool
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*enumValidator)(nil)

func (e *enumValidator) Condition() *validator.Condition {
	conditions := make([]ast.Expr, 0, len(e.enumValues))

	for _, value := range e.enumValues {
		var lit ast.Expr
		if e.isString || e.isCustom {
			lit = expr.StringLit(value)
		} else if e.isNumeric {
			lit = expr.IntLit(value)
		}

		conditions = append(conditions, expr.NEq(expr.Field("t", e.FieldName()), lit))
	}

	return &validator.Condition{
		Expr: expr.AndAll(conditions...),
	}
}

func (e *enumValidator) FieldName() string {
	return e.field.Names[0].Name
}

func (e *enumValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(e.structName, e.parentPath, e.FieldName())
}

func (e *enumValidator) ErrDecl() validator.ErrDecl {
	enumList := strings.Join(e.enumValues, ", ")

	return validator.ErrDecl{
		VarName: "Err" + e.FieldPath().CleanedPath() + "EnumValidation",
		Comment: fmt.Sprintf("is the error returned when the value is not in the allowed enum values %s.", enumList),
		Reason:  fmt.Sprintf("field %s must be one of %s", e.FieldName(), enumList),
		Path:    e.FieldPath().String(),
		Type:    e.ruleName,
	}
}

// ValidateEnum creates a new enumValidator for string, numeric, and custom types.
func ValidateEnum(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	enumValue, ok := input.Expressions[markers.GoValidMarkerEnum]
	if !ok {
		return nil
	}

	enumValues := strings.Split(enumValue, ",")

	for i, v := range enumValues {
		enumValues[i] = strings.TrimSpace(v)
	}

	if len(enumValues) == 0 {
		return nil
	}

	validator := &enumValidator{
		pass:       input.Pass,
		field:      input.Field,
		enumValues: enumValues,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}

	// Determine the type and set appropriate flags
	//nolint:exhaustive // This is a simplified version, assuming basic types and custom types.
	switch underlying := typ.Underlying().(type) {
	case *types.Basic:
		switch underlying.Kind() {
		case types.String:
			validator.isString = true
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
			types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
			types.Float32, types.Float64:
			validator.isNumeric = true
		default:
			// Unsupported basic type
			return nil
		}
	default:
		// For custom types (e.g., type Status string), treat as custom
		validator.isCustom = true
	}

	return validator
}
