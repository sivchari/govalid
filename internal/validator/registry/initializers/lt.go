package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// LtInitializer implements ValidatorInitializer for the lt validator.
type LtInitializer struct{}

// Marker returns the marker identifier for the lt validator.
func (l LtInitializer) Marker() string {
	return markers.GoValidMarkerLt
}

// Init initializes the lt validator factory.
func (l LtInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateLT
}