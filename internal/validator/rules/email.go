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

type emailValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
}

var _ validator.Validator = (*emailValidator)(nil)

const emailKey = "%s-email"

func (e *emailValidator) Validate() string {
	fieldName := e.FieldName()
	// Use external helper function for better maintainability
	return fmt.Sprintf("!validationhelper.IsValidEmail(t.%s)", fieldName)
}

func (e *emailValidator) FieldName() string {
	return e.field.Names[0].Name
}

func (e *emailValidator) Err() string {
	fieldName := e.FieldName()

	// No need to generate inline function - using external helper
	key := fmt.Sprintf(emailKey, e.structName+fieldName)
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the field is not a valid email address.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be a valid email address",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", e.ErrVariable(),
		"[@FIELD]", fieldName,
		"[@PATH]", fmt.Sprintf("%s.%s", e.structName, fieldName),
		"[@TYPE]", e.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (e *emailValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]EmailValidation", `[@PATH]`, e.structName+e.FieldName())
}

func (e *emailValidator) Imports() []string {
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateEmail creates a new emailValidator for string types.
func ValidateEmail(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName string, ruleName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &emailValidator{
		pass:       pass,
		field:      field,
		structName: structName,
		ruleName:   ruleName,
	}
}
