package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// LTEInitializer implements ValidatorInitializer for the lte validator.
type LTEInitializer struct{}

// Marker returns the marker identifier for the lte validator.
func (l LTEInitializer) Marker() string {
	return markers.GoValidMarkerLTE
}

// Init initializes the lte validator factory.
func (l LTEInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateLTE
}