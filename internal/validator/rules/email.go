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
	pass  *codegen.Pass
	field *ast.Field
}

var _ validator.Validator = (*emailValidator)(nil)

const emailKey = "%s-email"

func (e *emailValidator) Validate() string {
	fieldName := e.FieldName()
	// Generate inline manual validation for maximum performance
	return fmt.Sprintf("!isValidEmail(t.%s)", fieldName)
}

func (e *emailValidator) FieldName() string {
	return e.field.Names[0].Name
}

func (e *emailValidator) getEmailValidationHeader() string {
	return `
	// isValidEmail validates email format manually for maximum performance
	// Validates basic email structure: local@domain format
	isValidEmail = func(email string) bool {
		// Basic length check
		if len(email) < 5 || len(email) > 254 { // a@b.c = 5 chars minimum, RFC 5321 limit
			return false
		}
		
		// Must contain exactly one @
		atIndex := -1
		atCount := 0
		for i, c := range email {
			if c == '@' {
				atIndex = i
				atCount++
			}
		}
		if atCount != 1 || atIndex <= 0 || atIndex >= len(email)-1 {
			return false
		}`
}

func (e *emailValidator) getEmailLocalPartValidation() string {
	return `
		
		// Validate local part (before @)
		local := email[:atIndex]
		if len(local) == 0 || len(local) > 64 { // RFC 5321 limit
			return false
		}
		
		// Check for consecutive dots or leading/trailing dots
		if local[0] == '.' || local[len(local)-1] == '.' {
			return false
		}
		for i := 0; i < len(local)-1; i++ {
			if local[i] == '.' && local[i+1] == '.' {
				return false
			}
		}
		
		// Check allowed characters in local part
		for _, c := range local {
			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || 
				 (c >= '0' && c <= '9') || c == '.' || c == '_' || c == '-' || 
				 c == '+' || c == '=' || c == '!' || c == '#' || c == '$' || 
				 c == '%' || c == '&' || c == '\'' || c == '*' || c == '/' || 
				 c == '?' || c == '^' || c == 96 || c == '{' || c == '|' || 
				 c == '}' || c == '~') { // c == 96 is backtick
				return false
			}
		}`
}

func (e *emailValidator) getEmailDomainPartValidation() string {
	return `
		
		// Validate domain part (after @)
		domain := email[atIndex+1:]
		if len(domain) == 0 || len(domain) > 253 { // RFC 1035 limit
			return false
		}
		
		// Must contain at least one dot
		if !strings.Contains(domain, ".") {
			return false
		}
		
		// Cannot start or end with dot or hyphen
		if domain[0] == '.' || domain[len(domain)-1] == '.' ||
		   domain[0] == '-' || domain[len(domain)-1] == '-' {
			return false
		}
		
		// Split domain into labels and validate each
		labels := strings.Split(domain, ".")
		if len(labels) < 2 {
			return false
		}
		
		for _, label := range labels {
			if len(label) == 0 || len(label) > 63 { // RFC 1035 limit
				return false
			}
			
			// Cannot start or end with hyphen
			if label[0] == '-' || label[len(label)-1] == '-' {
				return false
			}
			
			// Check allowed characters (alphanumeric + hyphen)
			for _, c := range label {
				if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || 
					 (c >= '0' && c <= '9') || c == '-') {
					return false
				}
			}
		}
		
		return true
	}`
}

func (e *emailValidator) generateValidationFunction() string {
	return e.getEmailValidationHeader() +
		e.getEmailLocalPartValidation() +
		e.getEmailDomainPartValidation()
}

func (e *emailValidator) Err() string {
	fieldName := e.FieldName()

	var result strings.Builder

	// Generate isValidEmail function only once
	if !validator.GeneratorMemory["email-function-generated"] {
		validator.GeneratorMemory["email-function-generated"] = true

		result.WriteString(e.generateValidationFunction())
	}

	if validator.GeneratorMemory[fmt.Sprintf(emailKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(emailKey, fieldName)] = true

	result.WriteString(strings.ReplaceAll(`
	// Err@EmailValidation is the error returned when the field is not a valid email address.
	Err@EmailValidation = errors.New("field @ must be a valid email address")`, `@`, fieldName))

	return result.String()
}

func (e *emailValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@EmailValidation", `@`, e.FieldName())
}

func (e *emailValidator) Imports() []string {
	return []string{"strings"} // Need strings.Contains and strings.Split
}

// ValidateEmail creates a new emailValidator for string types.
func ValidateEmail(pass *codegen.Pass, field *ast.Field, _ map[string]string) validator.Validator {
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
