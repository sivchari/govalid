// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"regexp"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator"
)

type emailValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ validator.Validator = (*emailValidator)(nil)

const emailKey = "%s-email"

// HTML5-compliant email regex pattern
// This pattern covers most practical email formats while being performant
const emailRegexPattern = `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`

var emailRegex = regexp.MustCompile(emailRegexPattern)

func (e *emailValidator) Validate() string {
	fieldName := e.FieldName()
	// Generate inline regex validation for zero-allocation performance
	return fmt.Sprintf("!emailRegex.MatchString(t.%s)", fieldName)
}

func (e *emailValidator) FieldName() string {
	return e.field.Names[0].Name
}

func (e *emailValidator) Err() string {
	fieldName := e.FieldName()
	
	var result strings.Builder
	
	// Generate regex variable only once
	if !validator.GeneratorMemory["email-regex-generated"] {
		validator.GeneratorMemory["email-regex-generated"] = true
		result.WriteString(fmt.Sprintf(`
	// emailRegex is the compiled regex pattern for email validation.
	emailRegex = regexp.MustCompile(%q)`, emailRegexPattern))
	}
	
	if validator.GeneratorMemory[fmt.Sprintf(emailKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(emailKey, fieldName)] = true

	result.WriteString(fmt.Sprintf(strings.ReplaceAll(`
	// Err@EmailValidation is the error returned when the field is not a valid email address.
	Err@EmailValidation = errors.New("field @ must be a valid email address")`, "@", fieldName)))
	
	return result.String()
}

func (e *emailValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@EmailValidation", "@", e.FieldName())
}

func (e *emailValidator) Imports() []string {
	return []string{"regexp"}
}

// ValidateEmail creates a new emailValidator for string types.
func ValidateEmail(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &emailValidator{
		pass:  pass,
		field: field,
	}
}