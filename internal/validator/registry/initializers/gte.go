package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// GTEInitializer implements ValidatorInitializer for the gte validator.
type GTEInitializer struct{}

// Marker returns the marker identifier for the gte validator.
func (g GTEInitializer) Marker() string {
	return markers.GoValidMarkerGTE
}

// Init initializes the gte validator factory.
func (g GTEInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateGTE
}