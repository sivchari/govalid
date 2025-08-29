package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/registry"
)

type lengthValidator struct {
	pass        *codegen.Pass
	field       *ast.Field
	lengthValue string
	structName  string
	ruleName    string
	parentPath  string
}

var _ validator.Validator = (*lengthValidator)(nil)

const lengthKey = "%s-length"

func (l *lengthValidator) Validate() string {
	return fmt.Sprintf("utf8.RuneCountInString(t.%s) != %s", l.FieldName(), l.lengthValue)
}

func (l *lengthValidator) FieldName() string {
	return l.field.Names[0].Name
}

func (l *lengthValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(l.structName, l.parentPath, l.FieldName())
}

func (l *lengthValidator) Err() string {
	key := fmt.Sprintf(lengthKey, l.structName+l.FieldPath().CleanedPath())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const deprecationNoticeTemplate = `
		// Deprecated: Use [@ERRVARIABLE]
		//
		// [@LEGACYERRVAR] is deprecated and is kept for compatibility purpose.
		[@LEGACYERRVAR] = [@ERRVARIABLE]
	`

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the length of the field is not exactly [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] length must be exactly [@VALUE]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	legacyErrVarName := fmt.Sprintf("Err%s%sLengthValidation", l.structName, l.FieldName())
	currentErrVarName := l.ErrVariable()

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", currentErrVarName,
		"[@LEGACYERRVAR]", legacyErrVarName,
		"[@FIELD]", l.FieldName(),
		"[@PATH]", l.FieldPath().String(),
		"[@VALUE]", l.lengthValue,
		"[@TYPE]", l.ruleName,
	)

	if currentErrVarName != legacyErrVarName {
		return replacer.Replace(deprecationNoticeTemplate + errTemplate)
	}

	return replacer.Replace(errTemplate)
}

func (l *lengthValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]LengthValidation", "[@PATH]", l.FieldPath().CleanedPath())
}

func (l *lengthValidator) Imports() []string {
	return []string{"unicode/utf8"}
}

// ValidateLength creates a new lengthValidator if the field type is string and the length marker is present.
func ValidateLength(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	lengthValue, ok := input.Expressions[markers.GoValidMarkerLength]
	if !ok {
		return nil
	}

	return &lengthValidator{
		pass:        input.Pass,
		field:       input.Field,
		lengthValue: lengthValue,
		structName:  input.StructName,
		ruleName:    input.RuleName,
		parentPath:  input.ParentPath,
	}
}
