// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/validatorhelper"
)

type requiredValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
}

var _ validator.Validator = (*requiredValidator)(nil)

const requiredKey = "%s-required"

func (r *requiredValidator) Validate() string {
	typ := r.pass.TypesInfo.TypeOf(r.field.Type)

	return required(r.FieldName(), typ)
}

func required(name string, typ types.Type) string {
	// Handle slices, maps, and channels specifically for required validation
	switch typ.(type) {
	case *types.Slice, *types.Map, *types.Chan:
		return fmt.Sprintf("t.%s == nil", name)
	case *types.Array:
		return fmt.Sprintf("len(t.%s) == 0", name)
	}

	zero := validatorhelper.Zero(typ)
	if zero == "" {
		return ""
	}

	return fmt.Sprintf("t.%s == %s", name, zero)
}

func (r *requiredValidator) FieldName() string {
	return r.field.Names[0].Name
}

func (r *requiredValidator) Err() string {
	key := fmt.Sprintf(requiredKey, r.structName+r.FieldName())

	var result strings.Builder

	if validator.GeneratorMemory[key] {
		return result.String()
	}

	validator.GeneratorMemory[key] = true

	result.WriteString(
		strings.ReplaceAll(`
			// Err@RequiredValidation is returned when the @ is required but not provided.
			Err@RequiredValidation = govaliderrors.ValidationError{Reason:"field @ is required",Path:"PATH"}
			`,
			"@",
			r.structName+r.FieldName(),
		),
	)

	return strings.ReplaceAll(result.String(), "PATH", fmt.Sprintf("%s.%s", r.structName, r.FieldName()))
}

func (r *requiredValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@RequiredValidation", "@", r.structName+r.FieldName())
}

func (r *requiredValidator) Imports() []string {
	return []string{}
}

// ValidateRequired creates a new required validator for the given field.
func ValidateRequired(pass *codegen.Pass, field *ast.Field, structName string) validator.Validator {
	validator.GeneratorMemory[fmt.Sprintf(requiredKey, structName+field.Names[0].Name)] = false

	return &requiredValidator{
		pass:       pass,
		field:      field,
		structName: structName,
	}
}
