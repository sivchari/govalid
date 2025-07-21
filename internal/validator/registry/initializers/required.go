package initializers

import (
	"go/ast"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// RequiredInitializer implements ValidatorInitializer for the required validator.
type RequiredInitializer struct{}

// Marker returns the marker identifier for the required validator.
func (r RequiredInitializer) Marker() string {
	return markers.GoValidMarkerRequired
}

// Init initializes the required validator factory.
func (r RequiredInitializer) Init() registry.ValidatorFactory {
	// Special case: required doesn't use expressions parameter
	return func(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
		return rules.ValidateRequired(pass, field, structName)
	}
}