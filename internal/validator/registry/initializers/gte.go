package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// GteInitializer implements ValidatorInitializer for the gte validator.
type GteInitializer struct{}

// Marker returns the marker identifier for the gte validator.
func (g GteInitializer) Marker() string {
	return markers.GoValidMarkerGte
}

// Init initializes the gte validator factory.
func (g GteInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateGTE
}