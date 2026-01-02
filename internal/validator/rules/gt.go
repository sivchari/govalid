// Package rules implements validation rules for fields in structs.
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

type gtValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	gtValue    string
	validate   string
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*gtValidator)(nil)

const gtKey = "%s-gt"

func (m *gtValidator) Validate() string {
	if m.validate != "" {
		return m.validate
	}
	return fmt.Sprintf("!(t.%s > %s)", m.FieldName(), m.gtValue)
}

func (m *gtValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *gtValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(m.structName, m.parentPath, m.FieldName())
}

func (m *gtValidator) Err() string {
	key := fmt.Sprintf(gtKey, m.structName+m.FieldPath().CleanedPath())

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
		// [@ERRVARIABLE] is the error returned when the value of the field is less than the [@VALUE].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be greater than [@VALUE]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	legacyErrVarName := fmt.Sprintf("Err%s%sGTValidation", m.structName, m.FieldName())
	currentErrVarName := m.ErrVariable()

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", currentErrVarName,
		"[@LEGACYERRVAR]", legacyErrVarName,
		"[@FIELD]", m.FieldName(),
		"[@PATH]", m.FieldPath().String(),
		"[@VALUE]", m.gtValue,
		"[@TYPE]", m.ruleName,
	)

	if currentErrVarName != legacyErrVarName {
		return replacer.Replace(deprecationNoticeTemplate + errTemplate)
	}

	return replacer.Replace(errTemplate)
}

func (m *gtValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]GTValidation", "[@PATH]", m.FieldPath().CleanedPath())
}

func (m *gtValidator) Imports() []string {
	return []string{}
}

// ValidateGT creates a new gtValidator if the field type is numeric and the max marker is present.
func ValidateGT(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	underlying := typ.Underlying()

	fieldName := input.Field.Names[0].Name

	isNumeric := func(t types.Type) bool {
		basic, ok := t.Underlying().(*types.Basic)
		return ok && (basic.Info()&types.IsNumeric) != 0
	}

	// Support numeric scalar fields and slices/arrays of numeric elements.
	validType := false
	isIterable := false
	if isNumeric(underlying) {
		validType = true
	} else {
		switch t := underlying.(type) {
		case *types.Slice:
			if isNumeric(t.Elem()) {
				validType = true
				isIterable = true
			}
		case *types.Array:
			if isNumeric(t.Elem()) {
				validType = true
				isIterable = true
			}
		}
	}

	if !validType {
		return nil
	}

	gtValue, ok := input.Expressions[markers.GoValidMarkerGt]
	if !ok {
		return nil
	}

	validateExpr := fmt.Sprintf("!(t.%s > %s)", fieldName, gtValue)
	if isIterable {
		// Evaluate to true when any element violates the constraint.
		// Using an IIFE avoids introducing statements into the template.
		validateExpr = fmt.Sprintf(
			"func() bool { for _, v := range t.%s { if !(v > %s) { return true } }; return false }()",
			fieldName,
			gtValue,
		)
	}

	return &gtValidator{
		pass:       input.Pass,
		field:      input.Field,
		gtValue:    gtValue,
		validate:   validateExpr,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
