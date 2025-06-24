package validator

import (
	"go/ast"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator"
)

type requiredValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ Validator = (*requiredValidator)(nil)

func (r *requiredValidator) Validate() string {
	typ := r.pass.TypesInfo.TypeOf(r.field.Type)
	return validator.ZeroExpr(r.FieldName(), typ)
}

func (r *requiredValidator) FieldName() string {
	return r.field.Names[0].Name
}

func (r *requiredValidator) Err() string {
	if generatorMemory[r.FieldName()] {
		return ""
	}
	generatorMemory[r.FieldName()] = true
	return strings.ReplaceAll(`
	// Err @ is returned when the @ is required but not provided.
	Err@ = errors.New("field @ is required")`, "@", r.FieldName())
}

func ValidateRequired(pass *codegen.Pass, field *ast.Field) Validator {
	generatorMemory[field.Names[0].Name] = false
	return &requiredValidator{
		pass:  pass,
		field: field,
	}
}
