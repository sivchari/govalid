package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// CELInitializer implements ValidatorInitializer for the cel validator.
type CELInitializer struct{}

// Marker returns the marker identifier for the cel validator.
func (c CELInitializer) Marker() string {
	return markers.GoValidMarkerCEL
}

// Init initializes the cel validator factory.
func (c CELInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateCEL
}