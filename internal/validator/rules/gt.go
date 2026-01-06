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
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*gtValidator)(nil)

const gtKey = "%s-gt"

func (m *gtValidator) Validate() string {
	isIterable := func(t types.Type) bool {
		switch t.Underlying().(type) {
		case *types.Slice, *types.Array:
			return true
		default:
			return false
		}
	}

	fieldType := m.pass.TypesInfo.TypeOf(m.field.Type)
	var validateExpr string
	if isIterable(fieldType) {
		validateExpr = fmt.Sprintf(
			"func() bool { for _, v := range t.%s { if !(v > %s) { return true } }; return false }()",
			m.FieldName(),
			m.gtValue,
		)
	} else {
		validateExpr = fmt.Sprintf("!(t.%s > %s)", m.FieldName(), m.gtValue)
	}

	return validateExpr
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

// ValidateGT creates a new gtValidator if the field type is numeric or numeric iterable and the max marker is present.
func ValidateGT(input registry.ValidatorInput) validator.Validator {
	isNumeric := func(t types.Type) bool {
		basic, ok := t.Underlying().(*types.Basic)
		return ok && (basic.Info()&types.IsNumeric) != 0
	}

	isNumericIterable := func(t types.Type) bool {
		switch tp := t.Underlying().(type) {
		case *types.Slice:
			return isNumeric(tp.Elem())
		case *types.Array:
			return isNumeric(tp.Elem())
		default:
			return false
		}
	}

	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)
	validType := isNumeric(typ) || isNumericIterable(typ)
	if !validType {
		return nil
	}

	gtValue, ok := input.Expressions[markers.GoValidMarkerGt]
	if !ok {
		return nil
	}

	return &gtValidator{
		pass:       input.Pass,
		field:      input.Field,
		gtValue:    gtValue,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}
}
