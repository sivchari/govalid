package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// LTInitializer implements ValidatorInitializer for the lt validator.
type LTInitializer struct{}

// Marker returns the marker identifier for the lt validator.
func (l LTInitializer) Marker() string {
	return markers.GoValidMarkerLT
}

// Init initializes the lt validator factory.
func (l LTInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateLT
}