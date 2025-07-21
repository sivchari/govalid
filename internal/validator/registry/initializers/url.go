package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// URLInitializer implements ValidatorInitializer for the url validator.
type URLInitializer struct{}

// Marker returns the marker identifier for the url validator.
func (u URLInitializer) Marker() string {
	return markers.GoValidMarkerURL
}

// Init initializes the url validator factory.
func (u URLInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateURL
}