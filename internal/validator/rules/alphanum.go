package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

type alphanumValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
}

var _ validator.Validator = (*alphanumValidator)(nil)

const alphanumKey = "%s-alphanum"

func (v *alphanumValidator) Validate() string {
	fieldName := v.FieldName()

	return fmt.Sprintf("!validationhelper.IsValidAlphanum(t.%s)", fieldName)
}

func (v *alphanumValidator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *alphanumValidator) Err() string {
	fieldName := v.FieldName()

	var result strings.Builder

	key := fmt.Sprintf(alphanumKey, v.structName+fieldName)
	if validator.GeneratorMemory[key] {
		return result.String()
	}

	validator.GeneratorMemory[key] = true

	result.WriteString(strings.ReplaceAll(`
	// Err@AlphanumValidation is returned when the @ fails alphanum validation.
	Err@AlphanumValidation = errors.New("field @ failed alphanum validation")`, "@", v.structName+fieldName))

	return result.String()
}

func (v *alphanumValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@AlphanumValidation", "@", v.structName+v.FieldName())
}

func (v *alphanumValidator) Imports() []string {
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateAlphanum creates a new alphanum validator for the given field.
func ValidateAlphanum(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName string) validator.Validator {
	validator.GeneratorMemory[fmt.Sprintf(alphanumKey, structName+field.Names[0].Name)] = false

	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &alphanumValidator{
		pass:       pass,
		field:      field,
		structName: structName,
	}
}
