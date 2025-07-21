package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
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
	return rules.ValidateRequired
}