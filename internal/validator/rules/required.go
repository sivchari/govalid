// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/validatorhelper"
)

type requiredValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
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

func (r *requiredValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(r.structName, r.parentPath, r.FieldName())
}

func (r *requiredValidator) Err() string {
	key := fmt.Sprintf(requiredKey, r.FieldPath().CleanedPath())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is returned when the [@FIELD] is required but not provided.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] is required",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", r.ErrVariable(),
		"[@FIELD]", r.FieldName(),
		"[@PATH]", r.FieldPath().String(),
		"[@TYPE]", r.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (r *requiredValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]RequiredValidation", "[@PATH]", r.FieldPath().CleanedPath())
}

func (r *requiredValidator) Imports() []string {
	return []string{}
}

// ValidateRequired creates a new required validator for the given field.
func ValidateRequired(input registry.ValidatorInput) validator.Validator {
	fieldName := input.Field.Names[0].Name
	fieldPath := validator.NewFieldPath(input.StructName, input.ParentPath, fieldName)
	validator.GeneratorMemory[fmt.Sprintf(requiredKey, fieldPath.CleanedPath())] = false

	return &requiredValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
