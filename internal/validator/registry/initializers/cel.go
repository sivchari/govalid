package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// CelInitializer implements ValidatorInitializer for the cel validator.
type CelInitializer struct{}

// Marker returns the marker identifier for the cel validator.
func (c CelInitializer) Marker() string {
	return markers.GoValidMarkerCel
}

// Init initializes the cel validator factory.
func (c CelInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateCEL
}