package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// GTInitializer implements ValidatorInitializer for the gt validator.
type GTInitializer struct{}

// Marker returns the marker identifier for the gt validator.
func (g GTInitializer) Marker() string {
	return markers.GoValidMarkerGT
}

// Init initializes the gt validator factory.
func (g GTInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateGT
}