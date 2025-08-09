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

type alphaValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
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

func (v *alphaValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(v.structName, v.parentPath, v.FieldName())
}

func (v *alphaValidator) Err() string {
	key := fmt.Sprintf(alphaKey, v.FieldPath().CleanedPath())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when field [@FIELD] is not alphabetic.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be alphabetic",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", v.ErrVariable(),
		"[@FIELD]", v.FieldName(),
		"[@PATH]", v.FieldPath().String(),
		"[@TYPE]", v.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (v *alphaValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]AlphaValidation", "[@PATH]", v.FieldPath().CleanedPath())
}

func (v *alphaValidator) Imports() []string {
	// Import validation helper package
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateAlpha creates a new alphaValidator for string types.
func ValidateAlpha(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &alphaValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
