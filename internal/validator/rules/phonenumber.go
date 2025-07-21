// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

type phonenumberValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	value string
}

var _ validator.Validator = (*phonenumberValidator)(nil)

const phonenumberKey = "%s-phonenumber"

func (v *phonenumberValidator) Validate() string {
	// TODO: Implement validation logic
	// Return a Go expression that evaluates to true when validation fails
	return fmt.Sprintf("!isValidPhonenumber(t.%s, v.value)", v.FieldName())
}

func (v *phonenumberValidator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *phonenumberValidator) Err() string {
	key := fmt.Sprintf(phonenumberKey, v.structName+v.FieldName())
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	return strings.ReplaceAll(` + "`" + `
	// Err@PhonenumberValidation is returned when the @ fails phonenumber validation.
	Err@PhonenumberValidation = errors.New("field @ failed phonenumber validation")` + "`" + `, "@", v.structName+v.FieldName())
}

func (v *phonenumberValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@PhonenumberValidation", "@", v.structName+v.FieldName())
}

func (v *phonenumberValidator) Imports() []string {
	// TODO: Add any required imports
	return []string{}
}

// ValidatePhonenumber creates a new phonenumber validator for the given field.
func ValidatePhonenumber(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Type checking
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	validator.GeneratorMemory[fmt.Sprintf(phonenumberKey, structName+field.Names[0].Name)] = false

	return &phonenumberValidator{
		pass:       pass,
		field:      field,
		structName: structName,
		value: expressions["value"], // TODO: Parse parameter from expressions
	}
}