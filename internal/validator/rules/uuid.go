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
)

type uuidValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
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

func (u *uuidValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(u.structName, u.parentPath, u.FieldName())
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
	var result strings.Builder

	// Generate isValidUUID function only once
	if !validator.GeneratorMemory["uuid-function-generated"] {
		validator.GeneratorMemory["uuid-function-generated"] = true

		result.WriteString(u.generateValidationFunction())
	}

	key := fmt.Sprintf(uuidKey, u.structName+u.FieldPath().CleanedPath())
	if validator.GeneratorMemory[key] {
		return result.String()
	}

	validator.GeneratorMemory[key] = true

	const deprecationNoticeTemplate = `
		// Deprecated: Use [@ERRVARIABLE]
		//
		// [@LEGACYERRVAR] is deprecated and is kept for compatibility purpose.
		[@LEGACYERRVAR] = [@ERRVARIABLE]
	`

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the field is not a valid UUID.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be a valid UUID",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	legacyErrVarName := fmt.Sprintf("Err%s%sUUIDValidation", u.structName, u.FieldName())
	currentErrVarName := u.ErrVariable()

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", currentErrVarName,
		"[@LEGACYERRVAR]", legacyErrVarName,
		"[@FIELD]", u.FieldName(),
		"[@PATH]", u.FieldPath().String(),
		"[@TYPE]", u.ruleName,
	)

	if currentErrVarName != legacyErrVarName {
		result.WriteString(replacer.Replace(deprecationNoticeTemplate + errTemplate))
	} else {
		result.WriteString(replacer.Replace(errTemplate))
	}

	return result.String()
}

func (u *uuidValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]UUIDValidation", "[@PATH]", u.FieldPath().CleanedPath())
}

func (u *uuidValidator) Imports() []string {
	return []string{}
}

// ValidateUUID creates a new uuidValidator for string types.
func ValidateUUID(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &uuidValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
