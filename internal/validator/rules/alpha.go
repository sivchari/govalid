package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

type alphaValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
}

var _ validator.Validator = (*alphaValidator)(nil)

const alphaKey = "%s-alpha"

func (v *alphaValidator) Validate() string {
	// Use external helper function for better maintainability
	return fmt.Sprintf(`!validationhelper.IsValidAlpha(t.%s)`, v.FieldName())
}

func (v *alphaValidator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *alphaValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(alphaKey, v.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(alphaKey, v.FieldName())] = true

	return strings.ReplaceAll(`
    // Err@AlphaValidation is the error returned when field @ is not alphabetic.
    Err@AlphaValidation = errors.New("field @ must be alphabetic")`, "@", v.structName+v.FieldName())
}

func (v *alphaValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@AlphaValidation", "@", v.structName+v.FieldName())
}

func (v *alphaValidator) Imports() []string {
	// Import validation helper package
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateAlpha creates a new alphaValidator for string types.
func ValidateAlpha(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &alphaValidator{
		pass:       pass,
		field:      field,
		structName: structName,
	}
}
