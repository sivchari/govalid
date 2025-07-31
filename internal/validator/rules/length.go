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

type lengthValidator struct {
	pass        *codegen.Pass
	field       *ast.Field
	lengthValue string
	structName  string
}

var _ validator.Validator = (*lengthValidator)(nil)

const lengthKey = "%s-length"

func (l *lengthValidator) Validate() string {
	return fmt.Sprintf("utf8.RuneCountInString(t.%s) != %s", l.FieldName(), l.lengthValue)
}

func (l *lengthValidator) FieldName() string {
	return l.field.Names[0].Name
}

func (l *lengthValidator) Err() string {
	key := fmt.Sprintf(lengthKey, l.structName+l.FieldName())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// Err@LengthValidation is the error returned when the length of the field is not exactly [VALUE].
		Err@LengthValidation = govaliderrors.ValidationError{Reason:"field @ length must be exactly [VALUE]",Path:"PATH"}
	`

	replacer := strings.NewReplacer(
		"@", l.structName+l.FieldName(),
		"PATH", fmt.Sprintf("%s.%s", l.structName, l.FieldName()),
		"[VALUE]", l.lengthValue,
	)

	return replacer.Replace(errTemplate)
}

func (l *lengthValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@LengthValidation", "@", l.structName+l.FieldName())
}

func (l *lengthValidator) Imports() []string {
	return []string{"unicode/utf8"}
}

// ValidateLength creates a new lengthValidator if the field type is string and the length marker is present.
func ValidateLength(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	lengthValue, ok := expressions[markers.GoValidMarkerLength]
	if !ok {
		return nil
	}

	return &lengthValidator{
		pass:        pass,
		field:       field,
		lengthValue: lengthValue,
		structName:  structName,
	}
}
