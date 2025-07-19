package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
)

type alphaValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	alphaValue string
	structName string
}

var _ validator.Validator = (*alphaValidator)(nil)

const alphaKey = "%s-alpha"

func (v *alphaValidator) Validate() string {
	// Return the validation logic as Go code
	return fmt.Sprintf(`if !validationhelper.IsValidAlpha(t.%s) {return %s}`, v.FieldName(), v.alphaValue)
}

func (v *alphaValidator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *alphaValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(alphaKey, v.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(alphaKey, v.FieldName())] = true

	return fmt.Sprint(strings.ReplaceAll(`
    // Err@alphaValidation is the error returned when field @ is not alphabetic.
    Err@alphaValidation = errors.New("field @ must be alphabetic")`, "@", v.FieldName()))
}

func (v *alphaValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@alphaValidation", "@", v.FieldName())
}

func (v *alphaValidator) Imports() []string {
	// Return required imports for your validation logic
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// Validatealpha creates a new alphaValidator if conditions are met.
func ValidateAlpha(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	// Add type checking and marker validation logic
	// Return validator instance or nil
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsString) == 0 {
		return nil
	}

	alphaValue, ok := expressions[markers.GoValidMarkerAlpha]
	if !ok {
		return nil
	}

	return &alphaValidator{
		pass:       pass,
		field:      field,
		alphaValue: alphaValue,
		structName: structName,
	}
}
