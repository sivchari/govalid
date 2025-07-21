package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// LteInitializer implements ValidatorInitializer for the lte validator.
type LteInitializer struct{}

// Marker returns the marker identifier for the lte validator.
func (l LteInitializer) Marker() string {
	return markers.GoValidMarkerLte
}

// Init initializes the lte validator factory.
func (l LteInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateLTE
}