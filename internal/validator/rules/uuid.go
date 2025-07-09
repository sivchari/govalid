package rules

import (
	"fmt"
	"go/ast"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/errors"
	"github.com/sivchari/govalid/internal/validator"
)

// ValidateUUID validates a UUID field using external helper function.
// This function generates validation code that uses validationhelper.IsValidUUID
// for RFC 4122 compliant UUID validation.
func ValidateUUID(pass *codegen.Pass, field *ast.Field, expressions []string) validator.Validator {
	if len(expressions) != 0 {
		return &uuidValidator{
			pass:  pass,
			field: field,
			err:   errors.ErrInvalidUUIDExpression,
		}
	}

	return &uuidValidator{
		pass:  pass,
		field: field,
	}
}

type uuidValidator struct {
	pass  *codegen.Pass
	field *ast.Field
	err   error
}

func (u *uuidValidator) Validate() string {
	if u.err != nil {
		return ""
	}

	fieldName := u.FieldName()
	// Use external helper function for better maintainability
	return fmt.Sprintf("!validationhelper.IsValidUUID(t.%s)", fieldName)
}

func (u *uuidValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *uuidValidator) Err() string {
	if u.err != nil {
		return u.err.Error()
	}

	return fmt.Sprintf("field %s must be a valid UUID", u.FieldName())
}

func (u *uuidValidator) ErrVariable() string {
	return fmt.Sprintf("Err%s%sValidation", u.FieldName(), "UUID")
}

func (u *uuidValidator) Imports() []string {
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}