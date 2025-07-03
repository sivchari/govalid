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

type urlValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ validator.Validator = (*urlValidator)(nil)

const urlKey = "%s-url"

// Optimized URL regex pattern for performance
// This pattern validates common URL formats efficiently
const urlRegexPattern = `^https?://[a-zA-Z0-9.-]+[a-zA-Z0-9]+(:[0-9]+)?(/[^?\s]*)?(\?[^#\s]*)?(#[^\s]*)?$`

func (u *urlValidator) Validate() string {
	fieldName := u.FieldName()
	// Generate inline regex validation for zero-allocation performance
	return fmt.Sprintf("!urlRegex.MatchString(t.%s)", fieldName)
}

func (u *urlValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *urlValidator) Err() string {
	fieldName := u.FieldName()
	
	var result strings.Builder
	
	// Generate regex variable only once
	if !validator.GeneratorMemory["url-regex-generated"] {
		validator.GeneratorMemory["url-regex-generated"] = true
		result.WriteString(fmt.Sprintf(`
	// urlRegex is the compiled regex pattern for URL validation.
	urlRegex = regexp.MustCompile(%q)`, urlRegexPattern))
	}
	
	if validator.GeneratorMemory[fmt.Sprintf(urlKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(urlKey, fieldName)] = true

	result.WriteString(fmt.Sprintf(strings.ReplaceAll(`
	// Err@URLValidation is the error returned when the field is not a valid URL.
	Err@URLValidation = errors.New("field @ must be a valid URL")`, "@", fieldName)))
	
	return result.String()
}

func (u *urlValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@URLValidation", "@", u.FieldName())
}

func (u *urlValidator) Imports() []string {
	return []string{"regexp"}
}

// ValidateURL creates a new urlValidator for string types.
func ValidateURL(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &urlValidator{
		pass:  pass,
		field: field,
	}
}