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
	// Use validationhelper.IsValidUUID for centralized UUID validation
	return fmt.Sprintf("!validationhelper.IsValidUUID(t.%s)", fieldName)
}

func (u *uuidValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *uuidValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(u.structName, u.parentPath, u.FieldName())
}

func (u *uuidValidator) Err() string {
	var result strings.Builder

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
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
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
