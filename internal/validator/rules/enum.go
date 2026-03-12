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

type enumValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	enumValues []string
	isString   bool
	isNumeric  bool
	isCustom   bool
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*enumValidator)(nil)

const enumKey = "%s-enum"

func (e *enumValidator) Validate() string {
	fieldName := e.FieldName()

	var conditions []string

	for _, value := range e.enumValues {
		if e.isString || e.isCustom {
			// For string and custom types, use quoted comparison
			conditions = append(conditions, fmt.Sprintf("t.%s != %q", fieldName, value))
		} else if e.isNumeric {
			// For numeric types, use unquoted comparison
			conditions = append(conditions, fmt.Sprintf("t.%s != %s", fieldName, value))
		}
	}

	return strings.Join(conditions, " && ")
}

func (e *enumValidator) FieldName() string {
	return e.field.Names[0].Name
}

func (e *enumValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(e.structName, e.parentPath, e.FieldName())
}

func (e *enumValidator) Err() string {
	key := fmt.Sprintf(enumKey, e.structName+e.FieldPath().CleanedPath())

	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	enumList := strings.Join(e.enumValues, ", ")

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the value is not in the allowed enum values [@ENUM_LIST].
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be one of [@ENUM_LIST]",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", e.ErrVariable(),
		"[@FIELD]", e.FieldName(),
		"[@PATH]", e.FieldPath().String(),
		"[@ENUM_LIST]", enumList,
		"[@TYPE]", e.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (e *enumValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]EnumValidation", "[@PATH]", e.FieldPath().CleanedPath())
}

func (e *enumValidator) Imports() []string {
	return []string{}
}

// ValidateEnum creates a new enumValidator for string, numeric, and custom types.
func ValidateEnum(input registry.ValidatorInput) validator.Validator {
	typ := input.Pass.TypesInfo.TypeOf(input.Field.Type)

	enumValue, ok := input.Expressions[markers.GoValidMarkerEnum]
	if !ok {
		return nil
	}

	enumValues := strings.Split(enumValue, ",")

	for i, v := range enumValues {
		enumValues[i] = strings.TrimSpace(v)
	}

	if len(enumValues) == 0 {
		return nil
	}

	validator := &enumValidator{
		pass:       input.Pass,
		field:      input.Field,
		enumValues: enumValues,
		structName: input.StructName,
		ruleName:   input.RuleName,
		parentPath: input.ParentPath,
	}

	// Determine the type and set appropriate flags
	//nolint:exhaustive // This is a simplified version, assuming basic types and custom types.
	switch underlying := typ.Underlying().(type) {
	case *types.Basic:
		switch underlying.Kind() {
		case types.String:
			validator.isString = true
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
			types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
			types.Float32, types.Float64:
			validator.isNumeric = true
		default:
			// Unsupported basic type
			return nil
		}
	default:
		// For custom types (e.g., type Status string), treat as custom
		validator.isCustom = true
	}

	return validator
}
