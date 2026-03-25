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

type urlValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*urlValidator)(nil)

func (u *urlValidator) Validate() string {
	fieldName := u.FieldName()
	// Use external helper function for better maintainability
	return fmt.Sprintf("!validationhelper.IsValidURL(t.%s)", fieldName)
}

func (u *urlValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *urlValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(u.structName, u.parentPath, u.FieldName())
}

func (u *urlValidator) Err() string {
	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the field is not a valid URL.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be a valid URL",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", u.ErrVariable(),
		"[@FIELD]", u.FieldName(),
		"[@PATH]", u.FieldPath().String(),
		"[@TYPE]", u.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (u *urlValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]URLValidation", `[@PATH]`, u.FieldPath().CleanedPath())
}

func (u *urlValidator) Imports() []string {
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateURL creates a new urlValidator for string types.
func ValidateURL(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &urlValidator{
		pass:       input.Pass,
		field:      input.Field,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
