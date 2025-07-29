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

type uuidValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
}

var _ validator.Validator = (*uuidValidator)(nil)

const uuidKey = "%s-uuid"

func (u *uuidValidator) Validate() string {
	fieldName := u.FieldName()
	// Generate inline manual UUID validation for maximum performance
	return fmt.Sprintf("!isValidUUID(t.%s)", fieldName)
}

func (u *uuidValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *uuidValidator) getUUIDValidationHeader() string {
	return `
	// isValidUUID validates UUID format manually for maximum performance
	// Validates RFC 4122 format: 8-4-4-4-12 hex digits with hyphens
	isValidUUID = func(s string) bool {
		// Check length: 36 characters (32 hex + 4 hyphens)
		if len(s) != 36 {
			return false
		}
		
		// Check hyphen positions: 8-4-4-4-12
		if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
			return false
		}`
}

func (u *uuidValidator) getUUIDHexValidation() string {
	return `
		
		// Check hex characters and version/variant
		for i := 0; i < 36; i++ {
			if i == 8 || i == 13 || i == 18 || i == 23 {
				continue // skip hyphens
			}
			
			c := s[i]
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
				return false
			}
		}`
}

func (u *uuidValidator) getUUIDVersionVariantValidation() string {
	return `
		
		// Check version (position 14): must be 1-5
		version := s[14]
		if version < '1' || version > '5' {
			return false
		}
		
		// Check variant (position 19): must be 8, 9, A, B (case insensitive)
		variant := s[19]
		if !(variant == '8' || variant == '9' || 
			 variant == 'A' || variant == 'a' || 
			 variant == 'B' || variant == 'b') {
			return false
		}
		
		return true
	}`
}

func (u *uuidValidator) generateValidationFunction() string {
	return u.getUUIDValidationHeader() +
		u.getUUIDHexValidation() +
		u.getUUIDVersionVariantValidation()
}

func (u *uuidValidator) Err() string {
	fieldName := u.FieldName()

	var result strings.Builder

	// Generate isValidUUID function only once
	if !validator.GeneratorMemory["uuid-function-generated"] {
		validator.GeneratorMemory["uuid-function-generated"] = true

		result.WriteString(u.generateValidationFunction())
	}

	key := fmt.Sprintf(uuidKey, u.structName+fieldName)
	if validator.GeneratorMemory[key] {
		return result.String()
	}

	validator.GeneratorMemory[key] = true

	result.WriteString(
		strings.ReplaceAll(`
			// Err@UUIDValidation is the error returned when the field is not a valid UUID.
			Err@UUIDValidation = govaliderrors.ValidationError{Reason:"field @ must be a valid UUID",Path:"PATH"}
			`,
			"@",
			u.structName+u.FieldName(),
		),
	)

	return strings.ReplaceAll(result.String(), "PATH", fmt.Sprintf("%s.%s", u.structName, u.FieldName()))
}

func (u *uuidValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@UUIDValidation", "@", u.structName+u.FieldName())
}

func (u *uuidValidator) Imports() []string {
	return []string{}
}

// ValidateUUID creates a new uuidValidator for string types.
func ValidateUUID(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &uuidValidator{
		pass:       pass,
		field:      field,
		structName: structName,
	}
}
