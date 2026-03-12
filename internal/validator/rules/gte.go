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
	"github.com/sivchari/govalid/internal/validator/registry"
)

type gteValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	gteValue   string
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*gteValidator)(nil)

const gteKey = "%s-gte"

func (m *gteValidator) Validate() string {
	return fmt.Sprintf("!(t.%s >= %s)", m.FieldName(), m.gteValue)
}

func (m *gteValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *gteValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *gteValidator) Err() string {
	key := fmt.Sprintf(gteKey, m.structName+m.FieldPath().CleanedPath())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the value of the field is less than [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be greater than or equal to [@VALUE]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", m.ErrVariable(),
		"[@FIELD]", m.FieldName(),
		"[@PATH]", m.FieldPath().String(),
		"[@VALUE]", m.gteValue,
		"[@TYPE]", m.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (m *gteValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]GTEValidation", "[@PATH]", m.FieldPath().CleanedPath())
}

func (m *gteValidator) Imports() []string {
	return []string{}
}

// ValidateGTE creates a new gteValidator if the field type is numeric and the gte marker is present.
func ValidateGTE(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	gteValue, ok := input.Expressions[markers.GoValidMarkerGte]
	if !ok {
		return nil
	}

	return &gteValidator{
		pass:       input.Pass,
		field:      input.Field,
		gteValue:   gteValue,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
