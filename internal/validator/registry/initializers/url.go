package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// UrlInitializer implements ValidatorInitializer for the url validator.
type UrlInitializer struct{}

// Marker returns the marker identifier for the url validator.
func (u UrlInitializer) Marker() string {
	return markers.GoValidMarkerUrl
}

// Init initializes the url validator factory.
func (u UrlInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateURL
}