// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

type numericValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
}

var _ validator.Validator = (*numericValidator)(nil)

const numericKey = "%s-numeric"

func (m *numericValidator) Validate() string {
	return fmt.Sprintf("!validationhelper.IsNumeric(t.%s)", m.FieldName())
}

func (m *numericValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *numericValidator) Err() string {
	key := fmt.Sprintf(numericKey, m.structName+m.FieldName())
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the field [@FIELD] is not numeric.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be numeric",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", m.ErrVariable(),
		"[@FIELD]", m.FieldName(),
		"[@PATH]", fmt.Sprintf("%s.%s", m.structName, m.FieldName()),
		"[@TYPE]", m.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (m *numericValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]NumericValidation", "[@PATH]", m.structName+m.FieldName())
}

func (m *numericValidator) Imports() []string {
	return []string{
		"github.com/sivchari/govalid/validation/validationhelper",
	}
}

// ValidateNumeric creates a new numericValidator if the 'numeric' marker is present and field is string.
func ValidateNumeric(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName, ruleName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &numericValidator{
		pass:       pass,
		field:      field,
		structName: structName,
		ruleName:   ruleName,
	}
}
